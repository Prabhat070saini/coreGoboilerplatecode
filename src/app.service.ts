import { Injectable } from '@nestjs/common';
import { config } from './config/config';

@Injectable()
export class AppService {
  constructor() {}
  healthCheck(): string {
    return `Server is up and running on port ${config.app.port} and env is ${config.env} `;
  }
}
