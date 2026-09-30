// src/common/jwt/jwt.module.ts
import { Global, Module } from '@nestjs/common';
import { JwtModule } from '@nestjs/jwt';
import { TokenProvider } from './token.provider';
import { TokenService } from './token.service';

@Global()
@Module({
  imports: [JwtModule.register({})], // Optional, can stay empty since we pass secret dynamically
  providers: [TokenProvider, TokenService],
  exports: [TokenProvider, TokenService],
})
export class TokenModule {}
