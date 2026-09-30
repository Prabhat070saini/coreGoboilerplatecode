/* eslint-disable */

import { Injectable, LoggerService } from '@nestjs/common';
import * as winston from 'winston';
import { TracingService } from '../tracing.middleware.ts/tracing.service';
type LogLevel = 'debug' | 'info' | 'warn' | 'error';

@Injectable()
export class CustomLoggerService implements LoggerService {
  private readonly logger: winston.Logger;
  private readonly level: LogLevel;
  private readonly levelPriority: Record<LogLevel, number> = {
    debug: 3,
    info: 2,
    warn: 1,
    error: 0,
  };

  constructor(
    private readonly tracingService: TracingService,
    level: LogLevel,
  ) {
    // Add colors for custom levels

    this.level = level;
    winston.addColors({
      debug: 'magenta',
      info: 'green',
      warn: 'yellow',
      error: 'red',
    });

    this.logger = winston.createLogger({
      levels: this.levelPriority,
      level,
      transports: [
        new winston.transports.Console({
          format: winston.format.combine(
            winston.format.colorize({ all: true }),
            winston.format.timestamp({ format: 'YYYY-MM-DD HH:mm:ss' }),
            winston.format.printf(({ level, message, timestamp }) => {
              const tracingId = this.tracingService.getTracingId();
              const tracingInfo = tracingId ? `[tracing_id: ${tracingId}]` : '';
              const messageString =
                typeof message === 'string' ? message : JSON.stringify(message);
              const time =
                typeof timestamp === 'string'
                  ? timestamp
                  : (timestamp as Date).toISOString();
              return `[${level}] [${time}] ${tracingInfo} ${messageString}`;
            }),
          ),
        }),
      ],
    });
  }

  private shouldLog(msgLevel: LogLevel): boolean {
    return this.levelPriority[msgLevel] <= this.levelPriority[this.level];
  }

  log(message: any, context?: string, trace?: string) {
    if (this.shouldLog('info'))
      this.logger.info(this.formatMessage(message, context, trace));
  }

  info(message: any, context?: string, trace?: string) {
    if (this.shouldLog('info'))
      this.logger.info(this.formatMessage(message, context, trace));
  }

  warn(message: any, context?: string, trace?: string) {
    if (this.shouldLog('warn'))
      this.logger.warn(this.formatMessage(message, context, trace));
  }

  error(message: any, context?: string, trace?: string) {
    if (this.shouldLog('error'))
      this.logger.error(this.formatMessage(message, context, trace));
  }

  debug(message: any, context?: string, trace?: string) {
    if (this.shouldLog('debug'))
      this.logger.debug(this.formatMessage(message, context, trace));
  }

  private formatMessage(message: any, context?: string, trace?: string) {
    let msg = typeof message === 'string' ? message : JSON.stringify(message);
    if (context) msg = `[${context}] ${msg}`;
    if (trace) msg = `${msg} | trace: ${trace}`;
    return msg;
  }
}
