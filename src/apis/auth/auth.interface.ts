export interface ISignInResponse {
  accessToken: string;
  refreshToken: string;
  user_id: string;
}

export interface IGoogleOauthResponse {
  name: string;
  email: string;
  picture: string;
  isValid: boolean;
}

export interface ISendOtpResponse {
  cooldownRemaining?: number;
  attemptsLeft: number;
  identifier: string;
}

export interface IActiveOtp {
  otpHash: string;
  attempts: number;
}

export interface IRefreshTokenResponse {
  accessToken: string;
  refreshToken: string;
}
