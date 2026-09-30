export interface CacheBase {
  getKey<T>(key: string): Promise<T | undefined>;
  setKey<T>(key: string, value: T): Promise<void>;
  setKeyWithExpiry<T>(
    key: string,
    value: T,
    exp: number | string,
  ): Promise<void>;
  deleteKey(key: string): Promise<void>;
}

export const CACHE_BASE = 'CACHE_BASE';
