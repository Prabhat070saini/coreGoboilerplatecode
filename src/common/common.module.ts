import { Global, Module } from '@nestjs/common';
import { LibModule } from './lib/lib.module';
import { TokenModule } from './token/token.module';
import { UtilsModule } from './utils/utils.module';
@Global()
@Module({
  imports: [LibModule, TokenModule, UtilsModule],
  controllers: [],
  providers: [],
})
export class CommonModule {}
