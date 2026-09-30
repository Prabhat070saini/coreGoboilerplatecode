import { Module } from '@nestjs/common';
import { DatabaseModule } from './database/database.module';
import { CacheModule } from './cache/cache.module';
import { QueueModule } from './queue/queue.module';
import { QueueProviderType } from './queue/queue.constants';
import { config } from '../config/config';

const queueModule = QueueModule.forRoot(QueueProviderType.RABBITMQ, {
  url: config.queue.url,
  defaultRetries: config.queue.defaultRetries,
  defaultRetryDelayMs: config.queue.defaultRetryDelayMs,
  isDlq: config.queue.isDlq,
});
@Module({
  imports: [DatabaseModule, CacheModule, queueModule],
  controllers: [],
  providers: [],
  exports: [],
})
export class CoreModule {}
