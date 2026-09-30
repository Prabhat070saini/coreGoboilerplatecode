// src/common/jwt/token.provider.ts
import { Injectable, Logger } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { JwtConfig, TokenPayload } from 'src/common/constants/app.interface';

@Injectable()
export class TokenProvider {
  private readonly logger = new Logger(TokenProvider.name);
  constructor(private readonly jwtService: JwtService) {}

  generateToken(config: JwtConfig, payload: TokenPayload): string {
    return this.jwtService.sign(payload, {
      secret: config.secret,
      expiresIn: config.expiresIn,
    });
  }

  validateToken(config: JwtConfig, token: string): TokenPayload {
    try {
      return this.jwtService.verify(token, { secret: config.secret });
    } catch (error) {
      this.logger.error(`Invalid or expired token: ${error.message}`);
      throw new Error('Invalid or expired token');
    }
  }
}
