import { Type } from 'class-transformer';
import { IsInt, IsNotEmpty } from 'class-validator';
import { IsUlid } from '../decorators/validate-ulid.decorator';

export class IdValidateDto {
  @IsNotEmpty()
  @IsUlid()
  id: string;
}

export class IntIdValidateDto {
  @IsNotEmpty()
  @IsInt()
  @Type(() => Number)
  id: number;
}
