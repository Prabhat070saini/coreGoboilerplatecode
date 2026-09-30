import {
  CanActivate,
  ExecutionContext,
  Injectable,
  Logger,
  UnauthorizedException,
} from '@nestjs/common';
import { config } from 'src/config/config';
import { exception } from 'src/common/constants/exception';
import { Reflector } from '@nestjs/core';

@Injectable()
export class ValidateApiKeyGuard implements CanActivate {
  private readonly logger = new Logger(ValidateApiKeyGuard.name);

  constructor(private readonly reflector: Reflector) {}

  async canActivate(context: ExecutionContext): Promise<boolean> {
    this.logger.log(`[ValidateApiKeyGuard] Incoming request`);
    const request = context.switchToHttp().getRequest();
    const xApiKey = request.headers['x-api-key'];
    const isPublic =
      this.reflector.get<boolean>('isPublic', context.getHandler()) || false;
    if (isPublic) {
      return true;
    }

    if (!xApiKey) {
      this.logger.error(`[ValidateApiKeyGuard] Error: API key not found`);
      throw new UnauthorizedException(exception.API_KEY_MISSING);
    }

    if (xApiKey !== config.api_key.x_api_key) {
      this.logger.error(`[ValidateApiKeyGuard] Error: Invalid API key`);
      throw new UnauthorizedException(exception.API_KEY_INVALID);
    }

    return true;
  }
}
