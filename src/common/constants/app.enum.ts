export enum ETokenType {
  //   TemporaryToken = "temporaryToken",
  AccessToken = 'accessToken',
  RefreshToken = 'refreshToken',
}

export enum ESortDir {
  ASC = 'ASC',
  DESC = 'DESC',
}

export enum ESchema {
  DBO_SCHEMA = 'dbo',
}

export enum EFindUser {
  EMAIL = 'EMAIL',
  PHONE = 'PHONE',
  USER_ID = 'USER_ID',
}

export enum ESendOtpPurpose {
  FORGOT_PASSWORD = 'FORGOT_PASSWORD',
  SIGN_IN = 'SIGN_IN',
  SIGN_UP = 'SIGN_UP',
}
