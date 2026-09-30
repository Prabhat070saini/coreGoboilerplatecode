import { Injectable } from '@nestjs/common';
import { Inject } from '@nestjs/common';
import { DataSource } from 'typeorm';
import { TYPEORM_DATABASE_PROVIDER } from 'src/common/constants/app.constant';

@Injectable()
export class CommonRepository {
  constructor(
    @Inject(TYPEORM_DATABASE_PROVIDER) public readonly dataSource: DataSource,
  ) {}
}
