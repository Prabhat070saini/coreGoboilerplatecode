import { Injectable, Logger } from '@nestjs/common';
import { RedisClient } from './redis-client';
import { CacheBase } from '../cache.interface';
import { config } from '../../../config/config';

@Injectable()
export class RedisCacheService implements CacheBase {
  private readonly logger = new Logger(RedisCacheService.name);
  private readonly client;
  private readonly enabled = config.redis.use_redis;

  constructor(private readonly redisClient: RedisClient) {
    if (this.enabled) {
      this.client = this.redisClient.getClient();
      this.logger.log('🧠 Redis cache is ENABLED');
    } else {
      this.logger.warn('🚫 Redis cache is DISABLED via environment config');
    }
  }

  // Convert "10s", "5m", "2h", "1d" → seconds
  private parseExpiryToSeconds(exp: number | string): number {
    if (typeof exp === 'number') return exp;
    const match = /^(\d+)([smhd])$/.exec(exp.trim());
    if (!match) throw new Error(`Invalid expiry format: "${exp}"`);
    const value = Number.parseInt(match[1], 10);
    const unit = match[2];
    switch (unit) {
      case 's':
        return value;
      case 'm':
        return value * 60;
      case 'h':
        return value * 3600;
      case 'd':
        return value * 86400;
      default:
        throw new Error(`Unknown expiry unit: "${unit}"`);
    }
  }

  // === Get key with type safety ===
  async getKey<T>(key: string): Promise<T | undefined> {
    if (!this.enabled) return undefined;
    try {
      const data = await this.client.get(key);
      if (!data) return undefined;

      try {
        return JSON.parse(data) as T;
      } catch {
        return data as unknown as T; // fallback if it's not JSON
      }
    } catch (error) {
      throw new Error(
        `[RedisCacheService:getKey] Error retrieving key "${key}": ${error}`,
      );
    }
  }

  // === Set key ===
  async setKey<T>(key: string, value: T): Promise<void> {
    if (!this.enabled) return;
    try {
      const data = typeof value === 'string' ? value : JSON.stringify(value);
      await this.client.set(key, data);
      this.logger.log(`✅ Cache saved for key: ${key}`);
    } catch (error) {
      throw new Error(
        `[RedisCacheService:setKey] Error saving key "${key}": ${error}`,
      );
    }
  }

  // === Set key with expiry ===
  async setKeyWithExpiry<T>(
    key: string,
    value: T,
    exp: number | string,
  ): Promise<void> {
    if (!this.enabled) return;
    try {
      const expInSec = this.parseExpiryToSeconds(exp);
      const data = typeof value === 'string' ? value : JSON.stringify(value);
      await this.client.set(key, data, 'EX', expInSec);
      this.logger.log(
        `✅ Cache saved with expiry (${expInSec}s) for key: ${key}`,
      );
    } catch (error) {
      throw new Error(
        `[RedisCacheService:setKeyWithExpiry] Error saving key "${key}": ${error}`,
      );
    }
  }

  // === Delete key ===
  async deleteKey(key: string): Promise<void> {
    if (!this.enabled) return;
    try {
      const result = await this.client.del(key);
      if (result === 1)
        this.logger.log(`[deleteKey] Key "${key}" deleted from cache`);
      else
        this.logger.warn(
          `[deleteKey] Key "${key}" not found or already deleted`,
        );
    } catch (error) {
      throw new Error(
        `[RedisCacheService:deleteKey] Error deleting key "${key}": ${error}`,
      );
    }
  }
}
