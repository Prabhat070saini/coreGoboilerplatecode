import { ETokenType } from './app.enum';

export type JwtExpires = `${number}${'s' | 'm' | 'h' | 'd'}` | number;

export interface JwtConfig {
  secret: string;
  expiresIn: JwtExpires; // e.g., '1h', '7d', etc.
  type: ETokenType;
}

export interface TokenPayload {
  ref?: string;
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  payload?: any;
  type: ETokenType;
}
export interface IException {
  code: number;
  message: string;
  httpStatusCode?: number;
}

export interface ISuccess<T> {
  code: number;
  message: string;
  httpStatusCode: number;
  data?: T;
}

export interface IServiceOutput<T> {
  success?: ISuccess<T>;
  exception?: IException;
}

export interface IFunctionOutput<T> {
  data?: T;
  exception?: IException;
}
