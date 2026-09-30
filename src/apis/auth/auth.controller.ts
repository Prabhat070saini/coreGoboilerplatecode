/* eslint-disable */

import {
  Controller,
  Post,
  Body,
  Response,
  Get,
  UseGuards,
  Req,
  Headers,
} from '@nestjs/common';
import { AuthService } from './auth.service';
import { SignUpDto, SignInDto, SendOtpDto, VerifyOtpDto } from './dto/auth.dto';
import { UtilsService } from 'src/common/utils/utils.service';
import { AuthGuard } from '@nestjs/passport';
import { IsPublic } from 'src/common/decorators/public.decorator';
import { ApiHeader, ApiTags } from '@nestjs/swagger';

@ApiTags('Auth')
@Controller({ version: '1', path: 'auth' })
export class AuthController {
  constructor(
    private readonly authService: AuthService,
    private readonly utilsService: UtilsService,
  ) {}

  @Post('sign-up')
  async signUp(@Body() signUpDto: SignUpDto, @Response() res): Promise<void> {
    const output = await this.authService.signUp(signUpDto);
    return this.utilsService.sendRestResponse(res, output);
  }
  @Post('sign-in')
  async signIn(@Body() signInDto: SignInDto, @Response() res): Promise<void> {
    const output = await this.authService.signIn(signInDto);
    this.utilsService.sendRestResponse(res, output);
  }
  // Step 1: Redirect user to Google
  @ApiTags('Auth - Social')
  @IsPublic()
  @Get('google')
  @UseGuards(AuthGuard('google'))
  async googleAuth() {
    /* TODOdocument why this async method 'googleAuth' is empty */
  }

  // Step 2: Google redirects back here
  @ApiTags('Auth - Social')
  @IsPublic()
  @Get('google/callback')
  @UseGuards(AuthGuard('google'))
  async googleAuthRedirect(@Req() req) {
    return {
      message: 'Login successful',
      user: req.user,
    };
  }

  @Post('send-otp')
  async sendOtp(
    @Body() sendOtpDto: SendOtpDto,
    @Response() res,
  ): Promise<void> {
    const output = await this.authService.sendOtp(
      sendOtpDto.email,
      sendOtpDto.purpose,
    );
    this.utilsService.sendRestResponse(res, output);
  }

  @Post('verify-otp')
  async verifyOtp(
    @Body() verifyOtpDto: VerifyOtpDto,
    @Response() res,
  ): Promise<void> {
    const output = await this.authService.verifyOtp(
      verifyOtpDto.identifier,
      verifyOtpDto.purpose,
      verifyOtpDto.otp,
    );
    this.utilsService.sendRestResponse(res, output);
  }

  @ApiHeader({
    name: 'refresh-token',
    description: 'Refresh token sent via email',
    required: true,
  })
  @Post('refresh-token')
  async refreshToken(
    @Headers('refresh-token') refreshToken: string,
    @Response() res,
  ): Promise<void> {
    console.log(refreshToken);
    const output = await this.authService.refreshToken(refreshToken);
    this.utilsService.sendRestResponse(res, output);
  }
}
