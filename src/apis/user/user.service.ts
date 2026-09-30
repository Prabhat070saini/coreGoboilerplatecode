import { Injectable, Logger, HttpStatus } from '@nestjs/common';
import { UserRepository } from './repository/user.repository';
import { IServiceOutput } from 'src/common/constants/app.interface';
import { EFindUser } from 'src/common/constants/app.enum';
import { exception } from 'src/common/constants/exception';
@Injectable()
export class UserService {
  private readonly logger = new Logger(UserService.name);
  constructor(private readonly userRepository: UserRepository) {}

  async findUserById(userId: string): Promise<IServiceOutput<any>> {
    const { data: user, exception: findUserExp } =
      await this.userRepository.findUser({
        valueType: EFindUser.USER_ID,
        value: userId,
      });
    if (findUserExp) {
      this.logger.debug(`[findUserById] ${JSON.stringify(findUserExp)}`);
      return { exception: findUserExp };
    }
    if (!user) {
      this.logger.debug(`[findUserById] User not found with user_id ${userId}`);
      return { exception: exception.USER_NOT_FOUND };
    }
    return {
      success: {
        code: HttpStatus.OK,
        message: 'User found successfully',
        data: user,
        httpStatusCode: HttpStatus.OK,
      },
    };
  }
}
