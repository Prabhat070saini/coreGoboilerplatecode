import { EFindUser } from 'src/common/constants/app.enum';

export interface IUserCreate {
  name: string;
  email: string;
  password: string;
  createdBy?: number;
  phone?: string;
}

export interface IFindUser {
  valueType: EFindUser;
  value: string;
}
