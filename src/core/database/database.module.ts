import {
  Global,
  Module,
  OnApplicationShutdown,
  Injectable,
  Inject,
  Logger,
} from '@nestjs/common';
import { DatabaseConfig } from './database.config';
import { CAN_DATABASE_PROVIDER } from './database.provider';
import { DataSource } from 'typeorm';
import { TYPEORM_DATABASE_PROVIDER } from '../../common/constants/app.constant';

@Injectable()
class DatabaseShutdownService implements OnApplicationShutdown {
  private readonly logger = new Logger(DatabaseShutdownService.name);
  constructor(
    @Inject(TYPEORM_DATABASE_PROVIDER)
    private readonly dataSource: DataSource,
  ) {}

  async onApplicationShutdown(signal?: string): Promise<void> {
    this.logger.log(
      `🧹 [DatabaseModule] Shutting down due to signal: ${signal ?? 'manual stop'}`,
    );
    try {
      if (this.dataSource?.isInitialized) {
        await this.dataSource.destroy();
        this.logger.log('✅ Database connection closed gracefully');
      } else {
        this.logger.log(
          'ℹ️ Database connection already closed or not initialized',
        );
      }
    } catch (error) {
      this.logger.error('❌ Error while closing database connection:', error);
    }
  }
}
@Global()
@Module({
  providers: [
    DatabaseConfig,
    ...CAN_DATABASE_PROVIDER,
    DatabaseShutdownService,
  ],
  exports: [...CAN_DATABASE_PROVIDER],
})
export class DatabaseModule {}
