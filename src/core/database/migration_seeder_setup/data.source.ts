import { DataSource, DataSourceOptions } from 'typeorm';

import { SeederOptions } from 'typeorm-extension';
import InitSeeder from './seeds/init.seeder';
import { config } from '../../../config/config';

export const dataSourceOptions: DataSourceOptions & SeederOptions = {
  type: config.db.dialect,
  host: config.db.host,
  port: config.db.port,
  username: config.db.username,
  password: config.db.password,
  database: config.db.database,
  logging: true,
  entities: ['dist/**/*.entity.js'],
  migrations: ['dist/shared/database/migration_seeder_setup/migrations/*.js'],
  seeds: [InitSeeder],
  extra: {
    connectionLimit: Number(config.db.db_pool_max),
    waitForConnections: true,
    connectTimeout: Number(config.db.db_connection_timeout),
  },
};

const dataSource = new DataSource(dataSourceOptions);

export default dataSource;
