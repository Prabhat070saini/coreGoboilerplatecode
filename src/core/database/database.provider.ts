import { DataSource } from 'typeorm';
import { DatabaseConfig } from './database.config';
import { TYPEORM_DATABASE_PROVIDER } from '../../common/constants/app.constant';
import { Logger } from '@nestjs/common';

let dataSource: DataSource | null = null;
const logger = new Logger(DatabaseConfig.name);
export const CAN_DATABASE_PROVIDER = [
  {
    provide: TYPEORM_DATABASE_PROVIDER,
    useFactory: async (databaseConfig: DatabaseConfig): Promise<DataSource> => {
      if (!dataSource) {
        try {
          dataSource = new DataSource(databaseConfig.databaseConfig);
          await dataSource.initialize();
          logger.log(
            `✅,${databaseConfig.databaseConfig.type},Database connected successfully`,
          );
        } catch (error) {
          logger.error(
            '[Database connection] ❌ Failed to connect to the database:',
            error,
          );
          throw error;
        }
      }

      return dataSource;
    },
    inject: [DatabaseConfig],
  },
];
