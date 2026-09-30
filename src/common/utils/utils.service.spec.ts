import { Test, TestingModule } from '@nestjs/testing';
import { UtilsService } from './utils.service';
import { Response } from 'express';
import { IException } from '../exception/exception';
import {
  API_SUCCESS_CODE,
  API_SUCCESS_MESSAGE,
} from '../../core/constants/app.constants';
describe('UtilsService', () => {
  let utilsService: UtilsService;
  let res: Response;
  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [UtilsService],
    }).compile();
    utilsService = module.get<UtilsService>(UtilsService);
    res = {
      status: jest.fn().mockReturnThis(),
      json: jest.fn().mockReturnThis(),
    } as any;
  });
  it('should be defined', () => {
    expect(utilsService).toBeDefined();
  });
  describe('sendRestResponse', () => {
    it('should return an error response if exception is provided', () => {
      const exception: IException = {
        code: 'ERR001',
        message: 'Something went wrong',
        httpStatusCode: 400,
      };
      const output = { exception, message: 'Custom error message' };
      utilsService.sendRestResponse(res, output);
      expect(res.status).toHaveBeenCalledWith(400);
      expect(res.json).toHaveBeenCalledWith({
        code: 'ERR001',
        message: 'Custom error message',
      });
    });
    it('should return a success response if no exception is provided', () => {
      const output = { message: 'Success', outputData: { key: 'value' } };
      utilsService.sendRestResponse(res, output);
      expect(res.status).toHaveBeenCalledWith(200);
      expect(res.json).toHaveBeenCalledWith({
        code: API_SUCCESS_CODE,
        message: 'Success',
        data: { key: 'value' },
      });
    });
    it('should return a default success response when message is not provided', () => {
      const output = { outputData: { key: 'value' } };
      utilsService.sendRestResponse(res, output);
      expect(res.status).toHaveBeenCalledWith(200);
      expect(res.json).toHaveBeenCalledWith({
        code: API_SUCCESS_CODE,
        message: API_SUCCESS_MESSAGE,
        data: { key: 'value' },
      });
    });
    it('should return an empty data field if outputData is not provided', () => {
      const output = { message: 'Success' };
      utilsService.sendRestResponse(res, output);
      expect(res.status).toHaveBeenCalledWith(200);
      expect(res.json).toHaveBeenCalledWith({
        code: API_SUCCESS_CODE,
        message: 'Success',
        data: null,
      });
    });
  });
});
