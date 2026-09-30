import { MiddlewareConsumer, Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { ApisModule } from './apis/apis.module';
import { TracingMiddleware } from './common/lib/tracing.middleware.ts/tracing.middleware';
import { CommonModule } from './common/common.module';
import { APP_GUARD } from '@nestjs/core';
import { ValidateApiKeyGuard } from './common/guards/api-key.guard';
import { CoreModule } from './core/core.module';
@Module({
  imports: [CoreModule, ApisModule, CommonModule],
  controllers: [AppController],
  providers: [
    AppService,
    { provide: APP_GUARD, useClass: ValidateApiKeyGuard },
  ],
})
export class AppModule {
  configure(consumer: MiddlewareConsumer): void {
    consumer.apply(TracingMiddleware).forRoutes('*');
  }
}
