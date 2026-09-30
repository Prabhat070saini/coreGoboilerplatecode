/* eslint-disable */

import { PassportStrategy } from '@nestjs/passport';
import { Strategy } from 'passport-google-oauth20';
import { Injectable } from '@nestjs/common';
import { AuthService } from './auth.service';
import { IGoogleOauthResponse } from './auth.interface';

@Injectable()
export class GoogleStrategy extends PassportStrategy(Strategy, 'google') {
  constructor(private readonly authService: AuthService) {
    super({
      clientID: process.env.GOOGLE_CLIENT_ID,
      clientSecret: process.env.GOOGLE_CLIENT_SECRET,
      callbackURL: 'http://localhost:3000/api/v1/auth/google/callback',
      scope: ['email', 'profile'],
    });
  }

  async validate(_: string, __: string, profile: any) {
    const user: IGoogleOauthResponse = {
      email: profile.emails[0].value,
      isValid: profile.emails[0].verified,
      name: profile.displayName,
      picture: profile.photos[0].value,
    };

    return this.authService.validateGoogleUser(user);
  }
}
