import { applyDecorators } from '@nestjs/common';
import { ApiBearerAuth, ApiSecurity } from '@nestjs/swagger';

export const ApiJwtAndApiKey = (): MethodDecorator & ClassDecorator =>
  applyDecorators(ApiBearerAuth('access-token'), ApiSecurity('x-api-key'));
