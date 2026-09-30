export const xApiKeyHeader = 'x-api-key';
export const TYPEORM_DATABASE_PROVIDER = 'TYPEORM_DATABASE_PROVIDER';

export const SUCCESS_MESSAGE = {
  USER_SIGNUP_SUCCESSFULLY: 'User signup successfully',
  USER_LOGIN_SUCCESSFULLY: 'User login successfully',
  USER_LOGOUT_SUCCESSFULLY: 'User logout successfully',
  USER_REFRESH_TOKEN_SUCCESSFULLY: 'User refresh token successfully',
  SEND_OTP_SUCCESSFULLY: 'OTP sent successfully',
  OTP_VERIFIED_SUCCESSFULLY: 'OTP verified successfully',
};

export const EXCEPTION_MESSAGE = {
  USER_UNAUTHORIZED_REQUEST: 'Authentication refused the request',
  PHONE_NOT_REGISTERED:
    'This Phone number is not registered in our system, You can send invite',
  NO_DATA_FOUND: 'No data found',
  USER_NOT_FOUND: 'User not found',
  BAD_REQUEST: 'Bad request',
  SOMETHING_WENT_WRONG: 'Something went wrong. Please try again later.',
  OTP_EXPIRED: 'Otp expired!, please regenerate otp',
  INVALID_OTP: 'Otp invalid!, please enter correct otp',
  EMAIL_ALREADY_EXIST: 'This email already exist',
  DISABLE_USER: 'User is disabled',
  HEADERS_NOT_FOUND: 'Security headers not found in request',
  PERMISSION_DENINED: 'Permission Denied',
  PROTECTED_ROUTE: 'This route is protected',
  API_KEY_INVALID: 'Invalid API key provided',
  INVALID_OR_EXPIRED_ACCESS_TOKEN: 'Invalid or expired token',
  INVALID_OR_EXPIRED_REFRESH_TOKEN: 'Invalid or expired token',
  USER_NOT_VERIFIED: 'User not verified',
  INVALID_REFRESH_TOKEN: 'Invalid refresh token',
  INVALID_CREDENTIALS: 'Invalid credentials',
  API_KEY_MISSING: 'API key missing',
  OTP_COOLDOWN_ACTIVE:
    'You are sending OTPs too quickly. Please wait before trying again.',
  OTP_RATE_LIMIT_REACHED:
    'You have reached the OTP request limit. Please try again later.',
  MAX_ATTEMPTS_REACHED:
    'Maximum OTP verification attempts reached. Please try again later.',
};
