import { Module, Global } from '@nestjs/common';
import { RedisModule } from './redis/redis.module';
import { RedisCacheService } from './redis/redis.service';
import { CACHE_BASE, CacheBase } from './cache.interface';

@Global() // Makes CacheModule globally available
@Module({
  imports: [RedisModule], // Import other modules as necessary (e.g., RedisModule)
  providers: [
    {
      provide: CACHE_BASE, // This tells NestJS to inject this as the CacheBase implementation
      useFactory: (redis: RedisCacheService): CacheBase => {
        return redis;
      },
      inject: [RedisCacheService],
    },
  ],
  exports: [CACHE_BASE], // Export CacheBase to be used in other modules
})
export class CacheModule {}
