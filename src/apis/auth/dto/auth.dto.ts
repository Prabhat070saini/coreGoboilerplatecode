import {
  IsEmail,
  IsEnum,
  IsNotEmpty,
  IsString,
  IsStrongPassword,
  MaxLength,
  MinLength,
} from 'class-validator';
import { ApiProperty } from '@nestjs/swagger';
import { ESendOtpPurpose } from 'src/common/constants/app.enum';
import { IsUlid } from 'src/common/decorators/validate-ulid.decorator';

export class SignUpDto {
  @ApiProperty({ example: 'John Doe', description: 'User full name' })
  @IsNotEmpty()
  @IsString()
  @MinLength(3)
  name: string;

  @ApiProperty({
    example: 'user@example.com',
    description: 'User email address',
  })
  @IsNotEmpty()
  @IsEmail()
  @MaxLength(100)
  email: string;

  @ApiProperty({
    example: 'Password@123',
    description:
      'Strong password with min 8 chars, 1 uppercase, 1 lowercase, 1 number, 1 symbol',
  })
  @IsNotEmpty()
  @IsStrongPassword({
    minLength: 8,
    minLowercase: 1,
    minUppercase: 1,
    minNumbers: 1,
    minSymbols: 1,
  })
  @MaxLength(20)
  password: string;
}

export class SignInDto {
  @ApiProperty({
    example: 'user@example.com',
    description: 'User email address',
  })
  @IsNotEmpty()
  @IsEmail()
  @MaxLength(100)
  email: string;

  @ApiProperty({ example: 'Password@123', description: 'User password' })
  @IsNotEmpty()
  @IsString()
  password: string;
}

export class SendOtpDto {
  @ApiProperty({
    example: 'user@example.com',
    description: 'User email address',
  })
  @IsNotEmpty()
  @IsString()
  @MaxLength(100)
  email: string;

  @ApiProperty({
    enum: ESendOtpPurpose,
    example: ESendOtpPurpose.SIGN_UP,
    description: 'Purpose of the OTP',
  })
  @IsNotEmpty()
  @IsEnum(ESendOtpPurpose)
  purpose: ESendOtpPurpose;
}

export class VerifyOtpDto {
  @ApiProperty({
    example: '01ARZ3NDEKTSV4RRFFQ69G5FAV',
    description: 'ULID identifier',
  })
  @IsNotEmpty()
  @IsUlid()
  identifier: string;

  @ApiProperty({
    enum: ESendOtpPurpose,
    example: ESendOtpPurpose.SIGN_UP,
    description: 'Purpose of the OTP',
  })
  @IsNotEmpty()
  @IsEnum(ESendOtpPurpose)
  purpose: ESendOtpPurpose;

  @ApiProperty({ example: '123456', description: 'OTP code' })
  @IsNotEmpty()
  @IsString()
  otp: string;
}
