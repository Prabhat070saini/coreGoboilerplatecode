/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable @typescript-eslint/explicit-function-return-type */

import { Injectable, Logger, OnApplicationShutdown } from '@nestjs/common';
import { Redis } from 'ioredis';
import { config } from 'src/config/config';

@Injectable()
export class RedisClient implements OnApplicationShutdown {
  private readonly logger = new Logger(RedisClient.name);
  private readonly client: Redis;
  private connected = false;

  constructor() {
    if (!config.redis.use_redis) {
      this.logger.warn(
        '🚫 RedisClient not initialized because caching is disabled.',
      );
      return;
    }

    this.client = new Redis({
      host: config.redis.host,
      port: Number(config.redis.port),
      password: config.redis.password || undefined,
      retryStrategy: (times) => {
        const delay = Math.min(times * 200, 2000);
        this.logger.warn(
          `🔁 Redis reconnecting in ${delay}ms (attempt ${times})`,
        );
        return delay;
      },
    });

    // ✅ Listen for connection events
    // ✅ Listen for connection events
    this.client.on('connect', (): void => {
      this.connected = true;
      this.logger.log('✅ Redis connected successfully');
    });

    this.client.on('ready', (): void => {
      this.logger.log('🚀 Redis client ready to use');
    });

    this.client.on('error', (err: Error): void => {
      this.connected = false;
      this.logger.error(`❌ Redis Client Error: ${err.message}`, err.stack);
    });

    this.client.on('close', (): void => {
      this.connected = false;
      this.logger.warn('⚠️ Redis connection closed');
    });
  }

  getClient(): Redis {
    return this.client;
  }

  isConnected(): boolean {
    return this.connected;
  }

  async onApplicationShutdown(signal?: string): Promise<void> {
    this.logger.log(
      `🧹 [RedisClient] Shutting down due to signal: ${signal ?? 'manual stop'}`,
    );
  }
}
