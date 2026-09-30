import { Global, Module } from '@nestjs/common';
import { RedisClient } from './redis-client';
import { RedisCacheService } from './redis.service';
@Global()
@Module({
  providers: [RedisClient, RedisCacheService],
  exports: [RedisCacheService],
})
export class RedisModule {}
