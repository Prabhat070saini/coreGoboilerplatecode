import { DataSource, DeepPartial } from 'typeorm';
import { Inject, Injectable, Logger } from '@nestjs/common';
import { TYPEORM_DATABASE_PROVIDER } from '../../../common/constants/app.constant';
import { TransactionBaseRepository } from '../../../common/repository/transaction.repository';
import { User } from '../entities/user.entity';
import { EFindUser } from 'src/common/constants/app.enum';
import { IFunctionOutput } from 'src/common/constants/app.interface';
import { exception } from 'src/common/constants/exception';
import { IUserCreate, IFindUser } from '../user.interface';
import { ulid } from 'ulid';
@Injectable()
export class UserRepository extends TransactionBaseRepository<User> {
  private readonly logger = new Logger(UserRepository.name);
  constructor(
    @Inject(TYPEORM_DATABASE_PROVIDER)
    dataSource: DataSource,
  ) {
    super(dataSource, User);
  }
  async findUser(findUser: IFindUser): Promise<IFunctionOutput<User>> {
    try {
      let user: User;
      switch (findUser.valueType) {
        case EFindUser.EMAIL:
          user = await this.findOne({ where: { email: findUser.value } });
          break;
        case EFindUser.PHONE:
          user = await this.findOne({ where: { phone: findUser.value } });
          break;
        case EFindUser.USER_ID:
          user = await this.findOne({ where: { user_id: findUser.value } });
          break;
      }
      return { data: user };
    } catch (error) {
      this.logger.error(`[findUser] ${JSON.stringify(error)}`);
      return { exception: exception.SOMETHING_WENT_WRONG };
    }
  }

  async createUser(createUser: IUserCreate): Promise<IFunctionOutput<User>> {
    try {
      const newUser: DeepPartial<User> = {
        user_id: ulid(),
        name: createUser.name,
        email: createUser.email,
        password: createUser.password,
        created_by: createUser.createdBy || -1,
        phone: createUser.phone || null,
      };
      const user = await this.save({ entity: newUser });
      return { data: user };
    } catch (error) {
      this.logger.error(`[createUser] ${JSON.stringify(error)}`);
      return { exception: exception.SOMETHING_WENT_WRONG };
    }
  }
}
