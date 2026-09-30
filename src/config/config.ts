import * as dotenv from 'dotenv';
import * as Joi from 'joi';

dotenv.config();

// Define validation schema
const schema = Joi.object({
  // App Config
  X_API_KEY: Joi.string().required(),
  ENV_NAME: Joi.string().default('development'),
  LOG_LEVEL: Joi.string().default('debug'),
  SERVER_HOST: Joi.string().default('localhost'),
  SERVER_PORT: Joi.number().default(3000),

  // Database Config
  DB_DIALECT: Joi.string().required(),
  DB_HOST: Joi.string().required(),
  DB_PORT: Joi.number().required(),
  DB_USER: Joi.string().required(),
  DB_PASSWORD: Joi.string().required(),
  DB_NAME: Joi.string().required(),
  DB_POOL_MIN: Joi.number().default(5),
  DB_POOL_MAX: Joi.number().default(20),
  DB_CONNECTION_TIMEOUT: Joi.number().default(3000),
  DB_LOGGING: Joi.boolean().default(true).required(),
  DB_CONNECTION_IDLE_TIMEOUT: Joi.number().default(50),

  // Redis Config
  REDIS_HOST: Joi.string().required(),
  REDIS_PORT: Joi.number().required(),
  // REDIS_PASSWORD: Joi.string(),
  REDIS_CONNECTION_LIMIT: Joi.number().default(50),
  USE_REDIS: Joi.boolean().default(true),

  ACCESS_TOKEN_SECRET: Joi.string().required(),
  REFRESH_TOKEN_SECRET: Joi.string().required(),
  ACCESS_TOKEN_EXP_IN_MIN: Joi.string().required(),
  REFRESH_TOKEN_EXP_IN_MIN: Joi.string().required(),

  AWS_ACCESS_KEY_ID: Joi.string().required(),
  AWS_SECRET_ACCESS_KEY: Joi.string().required(),
  AWS_REGION: Joi.string().required(),
  AWS_S3_BUCKET_NAME: Joi.string().required(),
  SIGNED_URL_EXPIRE: Joi.number().required(),
  // STORAGE_CDN_URL: Joi.string(),

  // Email Config
  EMAIL_HOST: Joi.string().required(),
  EMAIL_PORT: Joi.number().required(),
  EMAIL_USERNAME: Joi.string().required(),
  EMAIL_PASSWORD: Joi.string().required(),
  EMAIL_FROM: Joi.string().required(),
  EMAIL_PROVIDER: Joi.string().required(),

  // Queue Config
  RABBITMQ_URL: Joi.string().required(),
  DEFAULT_RETRIES: Joi.number().required(),
  DEFAULT_RETRY_DELAY_MS: Joi.number().required(),
  IS_DLQ: Joi.boolean().required(),

  // Otp Config
  OTP_CRYPTO_SECRET: Joi.string().required(),
  OTP_TTL: Joi.number().required(),
  COOL_DOWN_TTL: Joi.number().required(),
  MAX_PER_HOUR: Joi.number().required(),
  MAX_ATTEMPTS: Joi.number().required(),
}).unknown(true);
// Validate process.env
const { error, value: env } = schema.validate(process.env, {
  allowUnknown: true,
});
if (error) {
  throw new Error(`Config validation error: ${error.message}`);
}

//  Export typed config object
export const config = {
  env: env.ENV_NAME,
  api_key: {
    x_api_key: env.X_API_KEY,
  },
  log_level: env.LOG_LEVEL,
  app: {
    host: env.SERVER_HOST,
    port: env.SERVER_PORT,
  },
  db: {
    dialect: env.DB_DIALECT,
    host: env.DB_HOST,
    port: env.DB_PORT,
    username: env.DB_USER,
    password: env.DB_PASSWORD,
    database: env.DB_NAME,
    db_connection_timeout: env.DB_CONNECTION_TIMEOUT,
    db_pool_min: env.DB_POOL_MIN,
    db_pool_max: env.DB_POOL_MAX,
    logging: env.DB_LOGGING,
    db_connection_idle_timeout: env.DB_CONNECTION_IDLE_TIMEOUT,
  },
  redis: {
    host: env.REDIS_HOST,
    port: env.REDIS_PORT,
    password: env.REDIS_PASSWORD,
    db_connection_limit: env.REDIS_CONNECTION_LIMIT,
    use_redis: env.USE_REDIS,
  },
  token: {
    access_token_secret: env.ACCESS_TOKEN_SECRET,
    access_token_exp_in_min:
      Number.parseInt(env.ACCESS_TOKEN_EXP_IN_MIN, 10) * 60, // Convert minutes to seconds
    refresh_token_exp_in_min:
      Number.parseInt(env.REFRESH_TOKEN_EXP_IN_MIN, 10) * 60, // Convert minutes to seconds
    refresh_token_secret: env.REFRESH_TOKEN_SECRET,
  },
  aws: {
    accessKeyId: env.AWS_ACCESS_KEY_ID,
    secretAccessKey: env.AWS_SECRET_ACCESS_KEY,
    region: env.AWS_REGION,
    bucketName: env.AWS_S3_BUCKET_NAME,
    signedUrlExpiry: env.SIGNED_URL_EXPIRE,
  },
  email: {
    host: env.EMAIL_HOST,
    port: env.EMAIL_PORT,
    username: env.EMAIL_USERNAME,
    password: env.EMAIL_PASSWORD,
    from: env.EMAIL_FROM,
    provider: env.EMAIL_PROVIDER,
  },
  otp: {
    otp_crypto_secret: env.OTP_CRYPTO_SECRET,
    otp_ttl: env.OTP_TTL,
    cool_down_ttl: env.COOL_DOWN_TTL,
    max_per_hour: env.MAX_PER_HOUR,
    max_attempts: env.MAX_ATTEMPTS,
  },
  queue: {
    url: env.RABBITMQ_URL,
    defaultRetries: env.DEFAULT_RETRIES,
    defaultRetryDelayMs: env.DEFAULT_RETRY_DELAY_MS,
    isDlq: env.IS_DLQ,
  },
};
