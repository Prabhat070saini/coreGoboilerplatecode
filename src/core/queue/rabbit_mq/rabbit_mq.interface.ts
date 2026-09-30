export interface RabbitMQConsumeOptions {
  retries?: number;
  retryDelayMs?: number;
  dlq?: string;
  queueOptions?: any;
}

export interface RabbitMQPublishOptions {
  persistent?: boolean;
  headers?: Record<string, any>;
  [key: string]: any;
}

export interface RabbitMQConfig {
  url: string;
  defaultRetries?: number;
  defaultRetryDelayMs?: number;
  globalQueueOptions?: any;
  isDlq?: boolean;
}
