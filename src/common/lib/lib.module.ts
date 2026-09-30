/* eslint-disable */

import { Global, Module } from '@nestjs/common';
import { TracingService } from './tracing.middleware.ts/tracing.service';
import { CustomLoggerService } from './logger/custom.logger';
import { config } from 'src/config/config';

type LogLevel = 'debug' | 'info' | 'warn' | 'error';

@Global()
@Module({
  providers: [
    TracingService,
    {
      provide: CustomLoggerService,
      useFactory: (tracingService: TracingService) => {
        return new CustomLoggerService(
          tracingService,
          config.log_level as LogLevel,
        );
      },
      inject: [TracingService],
    },
  ],
  exports: [TracingService, CustomLoggerService],
})
export class LibModule {}
