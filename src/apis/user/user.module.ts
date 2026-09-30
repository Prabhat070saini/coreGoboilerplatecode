import { Module } from '@nestjs/common';
import { UserService } from './user.service';
import { UserRepository } from './repository/user.repository';
import { UserController } from './user.controller';
@Module({
  imports: [],
  providers: [UserService, UserRepository],
  exports: [UserService],
  controllers: [UserController],
})
export class UserModule {}
