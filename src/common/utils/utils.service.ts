import { Injectable } from '@nestjs/common';
import { Response } from 'express';
import { IServiceOutput } from '../constants/app.interface';
import * as crypto from 'crypto';
import { ulid } from 'ulid';

@Injectable()
export class UtilsService {
  defaultMessage = 'ERROR';
  sendRestResponse(res: Response, output: IServiceOutput<any>): any {
    if (output?.exception) {
      const { code, message, httpStatusCode } = output.exception;
      return res
        .status(httpStatusCode || 500)
        .json({ code, message: message || this.defaultMessage });
    }

    return res.status(output?.success?.httpStatusCode || 200).json({
      code: output?.success?.code || 200,
      message: output?.success?.message || 'Success',
      data: output?.success?.data || null,
    });
  }

  getDurationAndUnit(seconds: number): { duration: number; unit: string } {
    if (seconds < 60) {
      return { duration: seconds, unit: 'Seconds' };
    } else if (seconds < 3600) {
      // Less than an hour
      const minutes = Math.floor(seconds / 60);
      return { duration: minutes, unit: 'Minutes' };
    } else {
      const hours = Math.floor(seconds / 3600);
      return { duration: hours, unit: 'Hours' };
    }
  }
  generatePassword(): string {
    const uppercase = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
    const lowercase = 'abcdefghijklmnopqrstuvwxyz';
    const digits = '0123456789';
    const specialChars = '!@#$^*';
    const allCharacters = uppercase + lowercase + digits + specialChars;

    // Ensure password length is between 8 and 24
    const passwordLength = Math.floor(Math.random() * (24 - 8 + 1)) + 8;

    // Collect at least one character of each type
    let password = [
      uppercase[Math.floor(Math.random() * uppercase.length)],
      lowercase[Math.floor(Math.random() * lowercase.length)],
      digits[Math.floor(Math.random() * digits.length)],
      specialChars[Math.floor(Math.random() * specialChars.length)],
    ];

    // Fill the rest of the password with random characters
    for (let i = password.length; i < passwordLength; i++) {
      password.push(
        allCharacters[Math.floor(Math.random() * allCharacters.length)],
      );
    }

    // Shuffle the password to ensure randomness
    password = password.sort(() => Math.random() - 0.5);

    return password.join('');
  }
  randomNumeric(length = 6): string {
    if (!Number.isInteger(length) || length <= 0)
      throw new Error('randomNumeric() requires a positive integer length.');
    try {
      const min = Math.pow(10, length - 1);
      const max = Math.pow(10, length) - 1;
      return crypto.randomInt(min, max).toString();
    } catch (err) {
      throw new Error(`randomNumeric() failed: ${(err as Error).message}`);
    }
  }

  generateUlId(): string {
    return ulid();
  }
}
