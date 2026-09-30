/**
 * Centralized Redis key generator for all modules.
 * Ensures consistent naming and easy refactoring.
 *
 * Always use these factories instead of string concatenation.
 */
import { Crypto } from 'src/common/lib/crypto/crypto';

export const RedisKeys = {
  // ===== AUTH KEYS =====
  auth: {
    userId: (id: string): string => `auth:userId:${id}`,
    accessToken: (userId: string): string => `auth:accessToken:${userId}`,
    refreshToken: (userId: string): string => `auth:refreshToken:${userId}`,
    session: (userId: string): string => `auth:session:${userId}`,
  },

  // ===== OTP KEYS =====
  otp: {
    active: (purpose: string, identifier: string): string =>
      Crypto.generateKey('otp:active', purpose, identifier),

    cooldown: (purpose: string, identifier: string): string =>
      Crypto.generateKey('otp:cooldown', purpose, identifier),

    rate: (purpose: string, identifier: string): string =>
      Crypto.generateKey('otp:rate', purpose, identifier),
  },

  // ===== USER CACHE =====
  user: {
    profile: (userId: string): string => `user:profile:${userId}`,
    settings: (userId: string): string => `user:settings:${userId}`,
  },

  // ===== GENERAL PURPOSE CACHE KEYS =====
  cache: {
    temp: (name: string): string => `cache:temp:${name}`,
    lock: (name: string): string => `cache:lock:${name}`,
  },
};
