/* eslint-disable */

import { BeforeApplicationShutdown, Logger } from '@nestjs/common';
import { Connection, Channel, connect } from 'amqplib';
import { IQueueProvider } from '../queue.interface';
import {
  RabbitMQConfig,
  RabbitMQConsumeOptions,
  RabbitMQPublishOptions,
} from './rabbit_mq.interface';
// remove console add this.logger.debug or info which is relawant
export class RabbitMQService
  implements IQueueProvider, BeforeApplicationShutdown
{
  private conn: Connection;
  private channel: Channel;
  private readonly logger = new Logger(RabbitMQService.name);

  constructor(private readonly config: RabbitMQConfig) {}

  async connect(): Promise<void> {
    this.conn = await connect(this.config.url);
    this.channel = await this.conn.createChannel();
    this.logger.log(`[RabbitMQ] ✅ Connected to ${this.config.url}`);
  }

  async publish(
    queue: string,
    message: any,
    options?: RabbitMQPublishOptions,
  ): Promise<void> {
    await this.channel.assertQueue(queue, {
      durable: true,
      ...this.config.globalQueueOptions,
    });
    this.channel.sendToQueue(queue, Buffer.from(JSON.stringify(message)), {
      persistent: true,
      headers: { 'x-attempt': 0 },
      ...options,
    });
  }

  async consume(
    queue: string,
    handler: (msg: any) => Promise<void>,
    options?: RabbitMQConsumeOptions,
    useDlq?: boolean,
  ): Promise<void> {
    const retries = options?.retries ?? this.config.defaultRetries ?? 3;
    const retryDelayMs =
      options?.retryDelayMs ?? this.config.defaultRetryDelayMs ?? 5000;
    const queueOpts = options?.queueOptions ?? this.config.globalQueueOptions;

    // Ensure queue exists
    try {
      await this.channel.assertQueue(queue, { durable: true, ...queueOpts });
    } catch (err: any) {
      if (err.code === 406) {
        this.logger.warn(
          `[RabbitMQ] Queue args mismatch, recreating queue: ${queue}`,
        );
        await this.channel.deleteQueue(queue);
        await this.channel.assertQueue(queue, { durable: true, ...queueOpts });
      } else throw err;
    }

    this.channel.consume(queue, async (msg) => {
      if (!msg) return;
      const headers = msg.properties.headers || {};
      const attempt = headers['x-attempt']
        ? Number.parseInt(headers['x-attempt'])
        : 0;

      try {
        const parsed = JSON.parse(msg.content.toString());
        await handler(parsed);
        this.channel.ack(msg);
      } catch (err) {
        this.logger.error(`[RabbitMQ] Message failed:`, err);

        if (attempt < retries) {
          setTimeout(() => {
            this.channel.sendToQueue(queue, msg.content, {
              headers: { ...headers, 'x-attempt': attempt + 1 },
              persistent: true,
            });
            this.channel.ack(msg);
          }, retryDelayMs);
        } else {
          if (this.config.isDlq && useDlq && !queue.startsWith('dlq-')) {
            const dlq = `dlq-${queue}`;
            await this.channel.assertQueue(dlq, { durable: true });
            this.logger.warn(
              `[RabbitMQ] Max retries reached, moving to DLQ: ${dlq}`,
            );
            await this.channel.sendToQueue(dlq, msg.content, {
              persistent: true,
            });
          } else {
            this.logger.warn(
              `[RabbitMQ] Max retries reached, discarding message`,
            );
          }
          this.channel.ack(msg);
        }
      }
    });
  }

  async disconnect(): Promise<void> {
    await this.channel?.close();
    await this.conn?.close();
  }

  async beforeApplicationShutdown(signal?: string) {
    this.logger.log(`[RabbitMQ] ⚠️ Shutting down due to: ${signal}`);
    await this.disconnect();
  }
}
