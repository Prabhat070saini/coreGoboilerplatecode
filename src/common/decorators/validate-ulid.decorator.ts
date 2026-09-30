/* eslint-disable */

import {
  registerDecorator,
  ValidationOptions,
  ValidationArguments,
} from 'class-validator';

export function IsUlid(validationOptions?: ValidationOptions) {
  return function (object: Object, propertyName: string) {
    registerDecorator({
      name: 'isUlid',
      target: object.constructor,
      propertyName,
      options: validationOptions,
      validator: {
        validate(value: any, _args: ValidationArguments) {
          if (typeof value !== 'string') return false;
          return /^[0-9A-HJKMNP-TV-Z]{26}$/.test(value);
        },
        defaultMessage(_args: ValidationArguments) {
          return `${propertyName} must be a valid ULID`;
        },
      },
    });
  };
}
