import * as crypto from 'crypto';

export class Crypto {
  private static readonly SECRET =
    process.env.CRYPTO_SECRET || 'default-secret-key';
  private static readonly IV_LENGTH = 16; // AES IV length

  // ====== HASHING ======

  static sha256(data: string): string {
    if (typeof data !== 'string' || data.length === 0)
      throw new Error('sha256() requires a non-empty string.');
    try {
      return crypto.createHash('sha256').update(data).digest('hex');
    } catch (err) {
      throw new Error(`sha256() failed: ${(err as Error).message}`);
    }
  }

  // ====== ENCRYPTION ======

  static encrypt(text: string): string {
    if (typeof text !== 'string' || text.length === 0)
      throw new Error('encrypt() requires a non-empty string.');
    try {
      const iv = crypto.randomBytes(this.IV_LENGTH);
      const key = crypto.createHash('sha256').update(this.SECRET).digest();
      const cipher = crypto.createCipheriv('aes-256-cbc', key, iv);
      let encrypted = cipher.update(text, 'utf8', 'hex');
      encrypted += cipher.final('hex');
      return iv.toString('hex') + ':' + encrypted;
    } catch (err) {
      throw new Error(`Encryption failed: ${(err as Error).message}`);
    }
  }

  static decrypt(encrypted: string): string {
    if (typeof encrypted !== 'string' || encrypted.length === 0)
      throw new Error('decrypt() requires a non-empty string.');
    try {
      const [ivHex, data] = encrypted.split(':');
      if (!ivHex || !data)
        throw new Error('Invalid encrypted payload format (expected iv:data).');

      const iv = Buffer.from(ivHex, 'hex');
      if (iv.length !== this.IV_LENGTH) throw new Error('Invalid IV length.');

      const key = crypto.createHash('sha256').update(this.SECRET).digest();
      const decipher = crypto.createDecipheriv('aes-256-cbc', key, iv);
      let decrypted = decipher.update(data, 'hex', 'utf8');
      decrypted += decipher.final('utf8');
      return decrypted;
    } catch (err) {
      throw new Error(`Decryption failed: ${(err as Error).message}`);
    }
  }

  // ====== RANDOM GENERATORS ======

  static randomHex(length: number): string {
    if (!Number.isInteger(length) || length <= 0)
      throw new Error('randomHex() requires a positive integer length.');
    try {
      return crypto.randomBytes(length).toString('hex');
    } catch (err) {
      throw new Error(`randomHex() failed: ${(err as Error).message}`);
    }
  }

  // ====== KEY HELPERS ======

  static generateKey(prefix: string, ...parts: (string | number)[]): string {
    if (typeof prefix !== 'string' || prefix.length === 0)
      throw new Error('generateKey() requires a non-empty prefix.');
    try {
      const safeParts = parts.map((p) =>
        String(p).replace(/[^a-zA-Z0-9@._:-]/g, ''),
      );
      return [prefix, ...safeParts].join(':');
    } catch (err) {
      throw new Error(`generateKey() failed: ${(err as Error).message}`);
    }
  }
}
