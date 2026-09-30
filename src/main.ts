import { NestFactory } from '@nestjs/core';
import { SwaggerModule, DocumentBuilder } from '@nestjs/swagger';
import { AppModule } from './app.module';
import { ValidationPipe, VersioningType } from '@nestjs/common';
import { config } from './config/config';
import { CustomLoggerService } from './common/lib/logger/custom.logger';
import * as bodyParser from 'body-parser';
import { xApiKeyHeader } from './common/constants/app.constant';
async function bootstrap(): Promise<void> {
  const app = await NestFactory.create(AppModule);

  /* This tells NestJS to use the global validation pipe for all requests*/
  app.useGlobalPipes(
    new ValidationPipe({
      whitelist: true,
      transform: true,
      forbidNonWhitelisted: true,
      enableDebugMessages: true,
    }),
  );
  app.use(bodyParser.json({ limit: '100mb' }));
  app.use(bodyParser.urlencoded({ limit: '100mb', extended: true }));
  /*This tells NestJS to listen for process termination signals*/
  app.enableShutdownHooks();

  /*This tells NestJS to allow cross-origin requests*/
  app.enableCors({ origin: '*' });

  /*This tells NestJS to set the global prefix to /api/*/
  app.setGlobalPrefix('api/');

  /*This tells NestJS to set the versioning to URI*/
  app.enableVersioning({ type: VersioningType.URI });

  /* Swagger Configuration */
  const options = new DocumentBuilder()
    .setTitle('Boilerplate API')
    .setDescription('The Boilerplate API description')
    .setVersion('1.0')
    .addApiKey(
      { type: 'apiKey', name: xApiKeyHeader, in: 'header' },
      xApiKeyHeader,
    )
    .addSecurityRequirements(xApiKeyHeader)
    .addBearerAuth(
      {
        type: 'http',
        scheme: 'bearer',
        bearerFormat: 'JWT',
      },
      'access-token',
    )
    .addServer(`http://localhost:${config.app.port}/`, 'Local Environment')
    .addServer('https://api.dev.example.com/', 'Development Environment') // Example host
    .addServer('https://api.prod.example.com/', 'Production Environment') // Example host
    .build();

  const document = SwaggerModule.createDocument(app, options);
  SwaggerModule.setup('docs', app, document, {
    swaggerOptions: {
      persistAuthorization: true,
      authAction: {
        [xApiKeyHeader]: {
          name: xApiKeyHeader,
          schema: {
            type: 'apiKey',
            in: 'header',
            name: xApiKeyHeader,
            description: '',
          },
          value: config.api_key.x_api_key,
        },
      },
    },
  });

  console.log(
    `Swagger UI is running on: http://localhost:${config.app.port}/docs`,
  );

  /*This tells NestJS to use the custom logger*/
  app.useLogger(app.get(CustomLoggerService));
  await app.listen(config.app.port);
}
bootstrap();
