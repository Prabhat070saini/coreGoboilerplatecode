import { DataSourceOptions } from 'typeorm';
import { config } from '../../config/config';
import { Injectable } from '@nestjs/common';

@Injectable()
export class DatabaseConfig {
  constructor() {}

  get databaseConfig(): DataSourceOptions {
    try {
      return {
        type: config.db.dialect,
        host: config.db.host,
        port: config.db.port,
        username: config.db.username,
        password: config.db.password,
        database: config.db.database,
        synchronize: false,
        logging: config.db.logging,
        entities: [__dirname + '/../../**/*.entity{.ts,.js}'],
        extra: {
          connectionLimit: Number(config.db.db_pool_max), // max connections
          min: Number(config.db.db_pool_min), // min connections
          waitForConnections: true, // wait for connections
          connectTimeout: Number(config.db.db_connection_timeout), // connection timeout
          maxIdleTime: Number(config.db.db_connection_idle_timeout), // max idle time
        },
      } as DataSourceOptions;
    } catch (error) {
      console.error(
        '[MySQL Database connection] Trouble connecting to the database:',
        error,
      );
      throw error;
    }
  }
}
