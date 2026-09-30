import { Injectable } from '@nestjs/common';
import axios, { AxiosHeaders, AxiosInstance, AxiosRequestConfig } from 'axios';
import { TracingService } from '../tracing.middleware.ts/tracing.service';

@Injectable()
export class HttpService {
  private readonly axiosInstance: AxiosInstance;

  constructor(private readonly tracingService: TracingService) {
    this.axiosInstance = axios.create();

    // Global request interceptor
    this.axiosInstance.interceptors.request.use(
      (config) => {
        const tracingId = this.tracingService.getTracingId();
        if (tracingId) {
          if (
            config.headers &&
            typeof (config.headers as any).set === 'function'
          ) {
            (config.headers as any).set('tracing_id', tracingId);
          } else {
            config.headers = AxiosHeaders.from({
              ...config.headers,
              tracing_id: tracingId,
            });
          }
        }

        return config;
      },
      (error) => {
        console.error(
          `[HTTP Request Error]: ${error.message}`,
          HttpService.name,
          error.stack,
        );
        return Promise.reject(error);
      },
    );

    // Global response interceptor
    this.axiosInstance.interceptors.response.use(
      (response) => {
        return response;
      },
      (error) => {
        console.error(
          ` [HTTP Response Error]: ${error.message}`,
          HttpService.name,
          error.stack,
        );
        return Promise.reject(error);
      },
    );
  }

  // Generic request builder
  private async request<T>(
    method: 'get' | 'post' | 'put' | 'patch' | 'delete',
    url: string,
    {
      body,
      headers,
      params,
    }: {
      body?: any;
      headers?: Record<string, string>;
      params?: Record<string, any>;
    } = {},
  ): Promise<T> {
    const config: AxiosRequestConfig = {
      method,
      url,
      data: body,
      headers,
      params,
    };

    const response = await this.axiosInstance.request<T>(config);
    return response.data;
  }

  // === Exposed convenience methods ===
  get<T>(
    url: string,
    opts?: { headers?: Record<string, string>; params?: Record<string, any> },
  ): Promise<T> {
    return this.request<T>('get', url, opts);
  }

  post<T>(
    url: string,
    body?: any,
    opts?: { headers?: Record<string, string>; params?: Record<string, any> },
  ): Promise<T> {
    return this.request<T>('post', url, { body, ...opts });
  }

  put<T>(
    url: string,
    body?: any,
    opts?: { headers?: Record<string, string>; params?: Record<string, any> },
  ): Promise<T> {
    return this.request<T>('put', url, { body, ...opts });
  }

  patch<T>(
    url: string,
    body?: any,
    opts?: { headers?: Record<string, string>; params?: Record<string, any> },
  ): Promise<T> {
    return this.request<T>('patch', url, { body, ...opts });
  }

  delete<T>(
    url: string,
    opts?: { headers?: Record<string, string>; params?: Record<string, any> },
  ): Promise<T> {
    return this.request<T>('delete', url, opts);
  }
}
