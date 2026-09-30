import { Test, TestingModule } from '@nestjs/testing';
import { HashingService } from './hashing.service';
import * as bcrypt from 'bcrypt';

describe('HashingService', () => {
  let hashingService: HashingService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [HashingService],
    }).compile();

    hashingService = module.get<HashingService>(HashingService);
  });

  describe('hashPassword', () => {
    it('should return a hashed password', async () => {
      const password = 'testPassword';
      const hashSpy = jest.spyOn(bcrypt, 'hash');
      const saltSpy = jest.spyOn(bcrypt, 'genSalt');

      const hashedPassword = await hashingService.hashPassword(password);

      expect(saltSpy).toHaveBeenCalledWith(10);
      expect(hashSpy).toHaveBeenCalled();
      expect(typeof hashedPassword).toBe('string');
      expect(hashedPassword).not.toBe(password);
    });
  });

  describe('matchPassword', () => {
    it('should return true if passwords match', async () => {
      const password = 'testPassword';
      const hashedPassword = await bcrypt.hash(password, 10);
      const compareSpy = jest.spyOn(bcrypt, 'compare');

      const isMatch = await hashingService.matchPassword(
        password,
        hashedPassword,
      );

      expect(compareSpy).toHaveBeenCalledWith(password, hashedPassword);
      expect(isMatch).toBe(true);
    });

    it('should return false if passwords do not match', async () => {
      const password = 'testPassword';
      const hashedPassword = await bcrypt.hash('differentPassword', 10);
      const compareSpy = jest.spyOn(bcrypt, 'compare');

      const isMatch = await hashingService.matchPassword(
        password,
        hashedPassword,
      );

      expect(compareSpy).toHaveBeenCalledWith(password, hashedPassword);
      expect(isMatch).toBe(false);
    });
  });
});
