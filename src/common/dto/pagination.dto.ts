import { Type } from 'class-transformer';
import { IsInt, Min } from 'class-validator';

export class PaginationDto {
  @IsInt()
  @Min(1)
  @Type(() => Number)
  pageSize: number = 10;

  @IsInt()
  @Min(1)
  @Type(() => Number)
  pageIndex: number = 1;
}
