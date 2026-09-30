import { Controller, UseGuards, Get, Response, Param } from '@nestjs/common';
import { UserService } from './user.service';
import { AuthnGuard } from 'src/common/guards/auth.guard';
import { UtilsService } from 'src/common/utils/utils.service';
import { PermissionGuard } from 'src/common/guards/permission.guard';
import { Role } from 'src/common/decorators/roles.decorator';
import { ApiTags } from '@nestjs/swagger';
import { ApiJwtAndApiKey } from 'src/common/decorators/apiKey-jwt-swagger.decorator';
@UseGuards(AuthnGuard)
@ApiTags('User')
@ApiJwtAndApiKey()
@Controller({ version: '1', path: 'user' })
export class UserController {
  constructor(
    private readonly userService: UserService,
    private readonly utilsService: UtilsService,
  ) {}
  @UseGuards(PermissionGuard)
  @Role('MST002')
  @Get(':userId')
  async getUserById(
    @Response() res,
    @Param('userId') userId: string,
  ): Promise<void> {
    const output = await this.userService.findUserById(userId);
    this.utilsService.sendRestResponse(res, output);
  }
}
