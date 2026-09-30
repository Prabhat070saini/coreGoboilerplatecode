// src/common/jwt/token.service.ts
import { Injectable, Logger } from '@nestjs/common';
import { ETokenType } from '../constants/app.enum';
import { JwtConfig, TokenPayload } from '../constants/app.interface';
import { config } from 'src/config/config';
import { TokenProvider } from './token.provider';

@Injectable()
export class TokenService {
  private readonly logger = new Logger(TokenService.name);
  private readonly tokenConfigs: Record<ETokenType, JwtConfig> = {
    [ETokenType.AccessToken]: {
      secret: config.token.access_token_secret,
      expiresIn: config.token.access_token_exp_in_min,
      type: ETokenType.AccessToken,
    },
    [ETokenType.RefreshToken]: {
      secret: config.token.refresh_token_secret,
      expiresIn: config.token.refresh_token_exp_in_min,
      type: ETokenType.RefreshToken,
    },
  };

  constructor(private readonly provider: TokenProvider) {}

  generate(payload: TokenPayload): string {
    try {
      const cfg = this.tokenConfigs[payload.type];
      return this.provider.generateToken(cfg, payload);
    } catch (error) {
      this.logger.error(`[generate] ${error}`, payload.type);
      return null;
    }
  }

  validate(type: ETokenType, token: string): TokenPayload {
    try {
      const cfg = this.tokenConfigs[type];
      return this.provider.validateToken(cfg, token);
    } catch (error) {
      this.logger.error(`[validate] error`, error);
      return null;
    }
  }
}
