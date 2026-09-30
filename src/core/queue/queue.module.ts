import { DynamicModule, Global, Module, Provider } from '@nestjs/common';
import { QUEUE_PROVIDER, QueueProviderType } from './queue.constants';
import { IQueueProvider, QueueProviderConfig } from './queue.interface';

@Global()
@Module({})
export class QueueModule {
  static forRoot<T extends QueueProviderConfig>(
    provider: QueueProviderType,
    config: T,
  ): DynamicModule {
    const queueProvider: Provider<IQueueProvider> = {
      provide: QUEUE_PROVIDER,
      useFactory: async (): Promise<IQueueProvider> => {
        if (provider === QueueProviderType.RABBITMQ) {
          const { RabbitMQService } = await import(
            './rabbit_mq/rabbit_mq.service.js'
          );
          const instance = new RabbitMQService(config as any);
          await instance.connect();
          return instance;
        }

        throw new Error(`Unsupported queue provider: ${provider}`);
      },
    };

    return {
      module: QueueModule,
      global: true,
      providers: [queueProvider],
      exports: [queueProvider],
    };
  }
}
