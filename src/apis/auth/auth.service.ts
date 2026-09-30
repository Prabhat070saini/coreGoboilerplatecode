import { Injectable, Logger, Inject, HttpStatus } from '@nestjs/common';

import { UserRepository } from '../user/repository/user.repository';
import { SignInDto, SignUpDto } from './dto/auth.dto';
import { ETokenType } from 'src/common/constants/app.enum';
import { exception } from 'src/common/constants/exception';
import { HashingService } from 'src/common/lib/hashing/hashing.service';
import { IUserCreate } from '../user/user.interface';
import {
  IServiceOutput,
  TokenPayload,
} from 'src/common/constants/app.interface';
import { TokenService } from 'src/common/token/token.service';
import {
  IActiveOtp,
  IGoogleOauthResponse,
  IRefreshTokenResponse,
  ISendOtpResponse,
  ISignInResponse,
} from './auth.interface';
import { UtilsService } from 'src/common/utils/utils.service';
import { CACHE_BASE, CacheBase } from 'src/core/cache/cache.interface';
import { config } from '../../config/config';
import { Crypto } from 'src/common/lib/crypto/crypto';
import { RedisKeys } from 'src/core/cache/keys';
import { SUCCESS_MESSAGE } from 'src/common/constants/app.constant';
@Injectable()
export class AuthService {
  private readonly logger = new Logger(AuthService.name);

  constructor(
    private readonly userRepository: UserRepository,
    private readonly hashingService: HashingService,
    private readonly tokenService: TokenService,
    private readonly utilsService: UtilsService,
    @Inject(CACHE_BASE) private readonly cacheService: CacheBase,
  ) {}

  async signUp(signUpDto: SignUpDto): Promise<IServiceOutput<null>> {
    try {
      const { email, password, name } = signUpDto;
      const existingUser = await this.userRepository.findOne({
        where: { email },
        select: { id: true },
      });

      if (existingUser) {
        this.logger.debug(`[signUp] ${JSON.stringify(existingUser)}`);
        return { exception: exception.EMAIL_ALREADY_EXIST };
      }

      const hashedPassword = await this.hashingService.hashPassword(password);
      const newUser: IUserCreate = {
        name,
        email,
        password: hashedPassword,
      };

      const { exception: createUserExp } =
        await this.userRepository.createUser(newUser);
      if (createUserExp) {
        this.logger.debug(`[signUp] ${JSON.stringify(createUserExp)}`);
        return { exception: createUserExp };
      }
      return {
        success: {
          code: HttpStatus.OK,
          message: SUCCESS_MESSAGE.USER_SIGNUP_SUCCESSFULLY,
          data: null,
          httpStatusCode: HttpStatus.OK,
        },
      };
    } catch (error) {
      this.logger.error(`[signUp] ${JSON.stringify(error)}`);
      return { exception: exception.SOMETHING_WENT_WRONG };
    }
  }

  async signIn(signInDto: SignInDto): Promise<IServiceOutput<ISignInResponse>> {
    try {
      const { email, password } = signInDto;
      const existingUser = await this.userRepository.findOne({
        where: { email },
        select: { user_id: true, password: true },
      });

      if (!existingUser) {
        this.logger.debug(`[signIn] User not found with email ${email}`);
        return { exception: exception.USER_NOT_FOUND };
      }
      const isPasswordMatch = await this.hashingService.matchPassword(
        password,
        existingUser.password,
      );
      if (!isPasswordMatch) {
        this.logger.debug(`[signIn] Invalid password for user ${email}`);
        return { exception: exception.INVALID_CREDENTIALS };
      }
      const accessTokenPayload: TokenPayload = {
        ref: existingUser.user_id,
        type: ETokenType.AccessToken,
      };
      const refreshTokenPayload: TokenPayload = {
        ref: existingUser.user_id,
        type: ETokenType.RefreshToken,
      };
      const accessToken = this.tokenService.generate(accessTokenPayload);
      const refreshToken = this.tokenService.generate(refreshTokenPayload);

      if (config.redis.use_redis) {
        await this.cacheService.setKeyWithExpiry(
          RedisKeys.auth.accessToken(existingUser.user_id),
          accessToken,
          config.token.access_token_exp_in_min,
        );
        await this.cacheService.setKeyWithExpiry(
          RedisKeys.auth.refreshToken(existingUser.user_id),
          refreshToken,
          config.token.refresh_token_exp_in_min,
        );
      }

      const resp: ISignInResponse = {
        accessToken,
        refreshToken,
        user_id: existingUser.user_id,
      };
      return {
        success: {
          code: HttpStatus.OK,
          message: SUCCESS_MESSAGE.USER_LOGIN_SUCCESSFULLY,
          data: resp,
          httpStatusCode: HttpStatus.OK,
        },
      };
    } catch (error) {
      this.logger.error(`[signIn] ${error}`);
      return { exception: exception.SOMETHING_WENT_WRONG };
    }
  }
  async validateGoogleUser(
    payload: IGoogleOauthResponse,
  ): Promise<IServiceOutput<ISignInResponse>> {
    try {
      const { isValid, email, name } = payload;
      if (!isValid) {
        this.logger.debug(
          `[validateGoogleUser] User not verified with email ${email}`,
        );
        return { exception: exception.USER_NOT_VERIFIED };
      }

      let existingUser = await this.userRepository.findOne({
        where: { email },
        select: { user_id: true, password: true },
      });

      if (!existingUser) {
        const randomPassword = this.utilsService.generatePassword();
        const hashedPassword =
          await this.hashingService.hashPassword(randomPassword);
        const newUser: IUserCreate = {
          name,
          email,
          password: hashedPassword,
        };

        const { data: createdUser, exception: createUserExp } =
          await this.userRepository.createUser(newUser);
        if (createUserExp) {
          this.logger.debug(`[signUp] ${JSON.stringify(createUserExp)}`);
          return { exception: createUserExp };
        }
        existingUser = createdUser;
      }

      const accessTokenPayload: TokenPayload = {
        ref: existingUser.user_id,
        type: ETokenType.AccessToken,
      };
      const refreshTokenPayload: TokenPayload = {
        ref: existingUser.user_id,
        type: ETokenType.RefreshToken,
      };
      const accessToken = this.tokenService.generate(accessTokenPayload);
      const refreshToken = this.tokenService.generate(refreshTokenPayload);

      if (config.redis.use_redis) {
        await this.cacheService.setKeyWithExpiry(
          RedisKeys.auth.accessToken(existingUser.user_id),
          accessToken,
          config.token.access_token_exp_in_min,
        );
        await this.cacheService.setKeyWithExpiry(
          RedisKeys.auth.refreshToken(existingUser.user_id),
          refreshToken,
          config.token.refresh_token_exp_in_min,
        );
      }

      const resp: ISignInResponse = {
        accessToken,
        refreshToken,
        user_id: existingUser.user_id,
      };
      return {
        success: {
          code: HttpStatus.OK,
          message: SUCCESS_MESSAGE.USER_LOGIN_SUCCESSFULLY,
          data: resp,
          httpStatusCode: HttpStatus.OK,
        },
      };
    } catch (error) {
      this.logger.error(`[validateGoogleUser] ${error}`);
      return { exception: exception.SOMETHING_WENT_WRONG };
    }
  }

  async refreshToken(
    refreshToken: string,
  ): Promise<IServiceOutput<IRefreshTokenResponse>> {
    try {
      const verifyTokenPayload = this.tokenService.validate(
        ETokenType.RefreshToken,
        refreshToken,
      );

      if (config.redis.use_redis) {
        const storedRefreshToken = await this.cacheService.getKey(
          RedisKeys.auth.refreshToken(verifyTokenPayload.ref),
        );

        if (!storedRefreshToken || storedRefreshToken !== refreshToken) {
          this.logger.debug(`[refreshToken] Invalid refresh token`);
          return { exception: exception.INVALID_REFRESH_TOKEN };
        }
      }

      const accessTokenPayload: TokenPayload = {
        ref: verifyTokenPayload.ref,
        type: ETokenType.AccessToken,
      };

      const accessToken = this.tokenService.generate(accessTokenPayload);

      if (config.redis.use_redis) {
        await this.cacheService.setKeyWithExpiry(
          RedisKeys.auth.accessToken(verifyTokenPayload.ref),
          accessToken,
          config.token.access_token_exp_in_min,
        );
      }

      const resp: IRefreshTokenResponse = {
        accessToken,
        refreshToken,
      };

      return {
        success: {
          code: HttpStatus.OK,
          message: SUCCESS_MESSAGE.USER_REFRESH_TOKEN_SUCCESSFULLY,
          data: resp,
          httpStatusCode: HttpStatus.OK,
        },
      };
    } catch (error) {
      this.logger.error(`[refreshToken] ${error}`);
      return { exception: exception.SOMETHING_WENT_WRONG };
    }
  }

  async sendOtp(
    email: string,
    purpose: string,
  ): Promise<IServiceOutput<ISendOtpResponse>> {
    try {
      const existingUser = await this.userRepository.findOne({
        where: { email: email },
        select: {
          user_id: true,
        },
      });

      if (!existingUser) {
        this.logger.debug(`[signIn] ${JSON.stringify(existingUser)}`);
        return { exception: exception.USER_NOT_FOUND };
      }
      const otpIdentifier = this.utilsService.generateUlId();
      const identifier = existingUser.user_id;
      let otpCount = 0;

      if (config.redis.use_redis) {
        const cooldown = await this.cacheService.getKey(
          RedisKeys.otp.cooldown(purpose, identifier),
        );
        if (cooldown) {
          return { exception: exception.OTP_COOLDOWN_ACTIVE };
        }

        const rate = await this.cacheService.getKey(
          RedisKeys.otp.rate(purpose, identifier),
        );
        otpCount = Number(rate || 0);

        if (otpCount >= config.otp.max_per_hour) {
          this.logger.debug(
            `[sendOtp] User has reached the maximum limit of OTPs`,
          );
          return { exception: exception.OTP_RATE_LIMIT_REACHED };
        }
      }

      const otp = this.utilsService.randomNumeric(6);
      const otpHash = Crypto.sha256(
        `${otpIdentifier}:${otp}:${purpose}:${config.otp.otp_crypto_secret}`,
      );

      if (config.redis.use_redis) {
        await this.cacheService.setKeyWithExpiry<IActiveOtp>(
          RedisKeys.otp.active(purpose, otpIdentifier),
          { otpHash, attempts: 0 },
          config.otp.otp_ttl,
        );
        await this.cacheService.setKeyWithExpiry(
          RedisKeys.otp.cooldown(purpose, identifier),
          '1',
          config.otp.cool_down_ttl,
        );
        await this.cacheService.setKeyWithExpiry(
          RedisKeys.otp.rate(purpose, identifier),
          String(otpCount + 1),
          3600,
        );
      }

      this.logger.warn(`[${purpose}] OTP for ${identifier}: ${otp}`);
      const attemptsLeft = config.otp.max_per_hour - (otpCount + 1);
      const resp = {
        cooldownRemaining: config.otp.cool_down_ttl,
        attemptsLeft,
        identifier: otpIdentifier,
      };
      return {
        success: {
          code: HttpStatus.OK,
          message: SUCCESS_MESSAGE.SEND_OTP_SUCCESSFULLY,
          data: resp,
          httpStatusCode: HttpStatus.OK,
        },
      };
    } catch (error) {
      this.logger.error(`[sendOtp] ${error}`);
      return { exception: exception.SOMETHING_WENT_WRONG };
    }
  }

  async verifyOtp(
    identifier: string,
    purpose: string,
    otp: string,
  ): Promise<IServiceOutput<null>> {
    try {
      if (!config.redis.use_redis) {
        // If Redis is disabled, we can't verify OTPs
        return { exception: exception.SOMETHING_WENT_WRONG };
      }

      const record = await this.cacheService.getKey<IActiveOtp>(
        RedisKeys.otp.active(purpose, identifier),
      );
      if (!record) {
        this.logger.debug(`[verifyOtp] OTP not found for ${identifier}`);
        return { exception: exception.INVALID_OTP };
      }

      if (record.attempts >= config.otp.max_attempts) {
        await this.cacheService.deleteKey(
          RedisKeys.otp.active(purpose, identifier),
        );
        return { exception: exception.MAX_ATTEMPTS_REACHED };
      }

      const inputHash = Crypto.sha256(
        `${identifier}:${otp}:${purpose}:${config.otp.otp_crypto_secret}`,
      );
      if (inputHash !== record.otpHash) {
        this.logger.debug(`[verifyOtp] Invalid OTP for ${identifier}`);
        record.attempts++;
        await this.cacheService.setKeyWithExpiry(
          RedisKeys.otp.active(purpose, identifier),
          record,
          config.otp.otp_ttl,
        );
        return { exception: exception.INVALID_OTP };
      }

      // Only delete the OTP if Redis is enabled
      if (config.redis.use_redis) {
        await this.cacheService.deleteKey(
          RedisKeys.otp.active(purpose, identifier),
        );
      }
      return {
        success: {
          code: HttpStatus.OK,
          message: SUCCESS_MESSAGE.OTP_VERIFIED_SUCCESSFULLY,
          data: null,
          httpStatusCode: HttpStatus.OK,
        },
      };
    } catch (error) {
      this.logger.error(`[verifyOtp] ${error}`);
      return { exception: exception.SOMETHING_WENT_WRONG };
    }
  }
}
