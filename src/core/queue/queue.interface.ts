export interface IQueueProvider {
  connect(): Promise<void>;
  publish(
    queueOrTopic: string,
    message: any,
    options?: Record<string, any>,
  ): Promise<void>;

  consume(
    queueOrTopic: string,
    handler: (message: any) => Promise<void>,
    options?: Record<string, any>,
    useDlq?: boolean,
  ): Promise<void>;

  disconnect(): Promise<void>;
}

export interface QueueProviderConfig {
  [key: string]: any;
}

export interface RabbitMQConfig extends QueueProviderConfig {
  url: string;
  defaultRetries?: number;
  defaultRetryDelayMs?: number;
  defaultDLQSuffix?: string;
  globalQueueOptions?: any;
  isDlq?: boolean;
}

export interface KafkaConfig extends QueueProviderConfig {
  brokers: string[];
  clientId: string;
  groupId: string;
}
