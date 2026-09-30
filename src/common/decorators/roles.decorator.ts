import { SetMetadata } from '@nestjs/common';

export const ROLES_KEY = 'role';

export const Role = (...roles: string[]): MethodDecorator & ClassDecorator => {
  return SetMetadata(ROLES_KEY, roles);
};
