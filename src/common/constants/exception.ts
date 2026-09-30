import { HttpStatus } from '@nestjs/common';
import { EXCEPTION_MESSAGE } from './app.constant';

export const exception = {
  USER_UNAUTHORIZED_REQUEST: {
    code: HttpStatus.UNAUTHORIZED,
    message: EXCEPTION_MESSAGE.USER_UNAUTHORIZED_REQUEST,
    httpStatusCode: HttpStatus.UNAUTHORIZED,
  },
  PHONE_NOT_REGISTERED: {
    code: HttpStatus.UNAUTHORIZED,
    message: EXCEPTION_MESSAGE.PHONE_NOT_REGISTERED,
    httpStatusCode: HttpStatus.UNAUTHORIZED,
  },
  NO_DATA_FOUND: {
    code: HttpStatus.NOT_FOUND,
    message: EXCEPTION_MESSAGE.NO_DATA_FOUND,
    httpStatusCode: HttpStatus.NOT_FOUND,
  },
  USER_NOT_FOUND: {
    code: HttpStatus.NOT_FOUND,
    message: EXCEPTION_MESSAGE.USER_NOT_FOUND,
    httpStatusCode: HttpStatus.NOT_FOUND,
  },
  BAD_REQUEST: {
    code: HttpStatus.BAD_REQUEST,
    message: EXCEPTION_MESSAGE.BAD_REQUEST,
    httpStatusCode: HttpStatus.BAD_REQUEST,
  },
  SOMETHING_WENT_WRONG: {
    code: HttpStatus.INTERNAL_SERVER_ERROR,
    message: EXCEPTION_MESSAGE.SOMETHING_WENT_WRONG,
    httpStatusCode: HttpStatus.INTERNAL_SERVER_ERROR,
  },
  OTP_EXPIRED: {
    code: HttpStatus.UNAUTHORIZED,
    message: EXCEPTION_MESSAGE.OTP_EXPIRED,
    httpStatusCode: HttpStatus.UNAUTHORIZED,
  },
  INVALID_OTP: {
    code: HttpStatus.UNAUTHORIZED,
    message: EXCEPTION_MESSAGE.INVALID_OTP,
    httpStatusCode: HttpStatus.UNAUTHORIZED,
  },
  EMAIL_ALREADY_EXIST: {
    code: HttpStatus.CONFLICT,
    message: EXCEPTION_MESSAGE.EMAIL_ALREADY_EXIST,
    httpStatusCode: HttpStatus.CONFLICT,
  },
  DISABLE_USER: {
    code: HttpStatus.UNAUTHORIZED,
    message: EXCEPTION_MESSAGE.DISABLE_USER,
    httpStatusCode: HttpStatus.UNAUTHORIZED,
  },
  HEADERS_NOT_FOUND: {
    code: HttpStatus.UNAUTHORIZED,
    message: EXCEPTION_MESSAGE.HEADERS_NOT_FOUND,
    httpStatusCode: HttpStatus.UNAUTHORIZED,
  },
  PERMISSION_DENINED: {
    code: HttpStatus.FORBIDDEN,
    message: EXCEPTION_MESSAGE.PERMISSION_DENINED,
    httpStatusCode: HttpStatus.FORBIDDEN,
  },
  PROTECTED_ROUTE: {
    code: HttpStatus.FORBIDDEN,
    message: EXCEPTION_MESSAGE.PROTECTED_ROUTE,
    httpStatusCode: HttpStatus.FORBIDDEN,
  },
  API_KEY_INVALID: {
    code: HttpStatus.FORBIDDEN,
    message: EXCEPTION_MESSAGE.API_KEY_INVALID,
    httpStatusCode: HttpStatus.FORBIDDEN,
  },
  INVALID_OR_EXPIRED_ACCESS_TOKEN: {
    code: 6001, // call refresh token code
    message: EXCEPTION_MESSAGE.INVALID_OR_EXPIRED_ACCESS_TOKEN,
    httpStatusCode: HttpStatus.UNAUTHORIZED,
  },
  INVALID_OR_EXPIRED_REFRESH_TOKEN: {
    code: 6002, // logout user code
    message: EXCEPTION_MESSAGE.INVALID_OR_EXPIRED_REFRESH_TOKEN,
    httpStatusCode: HttpStatus.UNAUTHORIZED,
  },
  BASE_64_DECODING_FAILED: {
    code: 5002,
    message: EXCEPTION_MESSAGE.SOMETHING_WENT_WRONG,
    httpStatusCode: HttpStatus.INTERNAL_SERVER_ERROR,
  },
  BASE_64_ENCODING_FAILED: {
    code: 5001,
    message: EXCEPTION_MESSAGE.SOMETHING_WENT_WRONG,
    httpStatusCode: HttpStatus.INTERNAL_SERVER_ERROR,
  },
  ENCRYPTION_FAILED: {
    code: 5101,
    message: EXCEPTION_MESSAGE.SOMETHING_WENT_WRONG,
    httpStatusCode: HttpStatus.INTERNAL_SERVER_ERROR,
  },
  DECRYPTION_FAILED: {
    code: 5102,
    message: EXCEPTION_MESSAGE.SOMETHING_WENT_WRONG,
    httpStatusCode: HttpStatus.INTERNAL_SERVER_ERROR,
  },
  USER_NOT_VERIFIED: {
    code: 6001,
    message: EXCEPTION_MESSAGE.USER_NOT_VERIFIED,
    httpStatusCode: HttpStatus.UNAUTHORIZED,
  },
  INVALID_REFRESH_TOKEN: {
    code: 6002,
    message: EXCEPTION_MESSAGE.INVALID_REFRESH_TOKEN,
    httpStatusCode: HttpStatus.UNAUTHORIZED,
  },
  INVALID_CREDENTIALS: {
    code: 6003,
    message: EXCEPTION_MESSAGE.INVALID_CREDENTIALS,
    httpStatusCode: HttpStatus.UNAUTHORIZED,
  },
  API_KEY_MISSING: {
    code: 6004,
    message: EXCEPTION_MESSAGE.API_KEY_MISSING,
    httpStatusCode: HttpStatus.UNAUTHORIZED,
  },
  OTP_COOLDOWN_ACTIVE: {
    code: 7001,
    message: EXCEPTION_MESSAGE.OTP_COOLDOWN_ACTIVE,
    httpStatusCode: HttpStatus.TOO_MANY_REQUESTS,
  },
  OTP_RATE_LIMIT_REACHED: {
    code: 7002,
    message: EXCEPTION_MESSAGE.OTP_RATE_LIMIT_REACHED,
    httpStatusCode: HttpStatus.TOO_MANY_REQUESTS,
  },
  MAX_ATTEMPTS_REACHED: {
    code: 7003,
    message: EXCEPTION_MESSAGE.MAX_ATTEMPTS_REACHED,
    httpStatusCode: HttpStatus.TOO_MANY_REQUESTS,
  },
};
