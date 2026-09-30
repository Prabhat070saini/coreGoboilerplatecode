/* eslint-disable @typescript-eslint/no-explicit-any */
import {
  QueryRunner,
  DataSource,
  FindManyOptions,
  FindOneOptions,
  DeleteResult,
  ObjectId,
  FindOptionsWhere,
  EntityTarget,
  DeepPartial,
  SaveOptions,
  UpdateResult,
  SelectQueryBuilder,
} from 'typeorm';
import { QueryDeepPartialEntity } from 'typeorm/query-builder/QueryPartialEntity';

type CriteriaType<T> =
  | string
  | number
  | Date
  | ObjectId
  | string[]
  | number[]
  | Date[]
  | ObjectId[]
  | FindOptionsWhere<T>;

export class TransactionBaseRepository<T> {
  constructor(
    public readonly dataSource: DataSource,
    private entityClass: EntityTarget<T>,
  ) {}
  private getQueryRunner(queryRunner?: QueryRunner): QueryRunner {
    if (queryRunner) return queryRunner;
    return this.dataSource.createQueryRunner();
  }
  async findOne(
    option: FindOneOptions<T>,
    queryRunner?: QueryRunner,
  ): Promise<T> {
    const internalQueryRunner = !queryRunner;
    queryRunner = this.getQueryRunner(queryRunner);
    try {
      return await queryRunner.manager.findOne(this.entityClass, option);
    } finally {
      if (internalQueryRunner && !queryRunner.isReleased) {
        queryRunner.release();
      }
    }
  }
  async find(
    option: FindManyOptions<T>,
    queryRunner?: QueryRunner,
  ): Promise<T[]> {
    const internalQueryRunner = !queryRunner;
    queryRunner = this.getQueryRunner(queryRunner);
    try {
      return await queryRunner.manager.find(this.entityClass, option);
    } finally {
      if (internalQueryRunner && !queryRunner.isReleased) {
        queryRunner.release();
      }
    }
  }
  async delete(
    option:
      | string
      | number
      | Date
      | ObjectId
      | string[]
      | number[]
      | Date[]
      | ObjectId[]
      | FindOptionsWhere<T>,
    queryRunner?: QueryRunner,
  ): Promise<DeleteResult> {
    const internalQueryRunner = !queryRunner;
    queryRunner = this.getQueryRunner(queryRunner);
    try {
      return await queryRunner.manager.delete(this.entityClass, option);
    } finally {
      if (internalQueryRunner && !queryRunner.isReleased) {
        queryRunner.release();
      }
    }
  }
  create(entityLike: DeepPartial<T>, queryRunner?: QueryRunner): T {
    const internalQueryRunner = !queryRunner;
    queryRunner = this.getQueryRunner(queryRunner);
    const entity = queryRunner.manager.create(this.entityClass, entityLike);
    if (internalQueryRunner && !queryRunner.isReleased) {
      queryRunner.release();
    }
    return entity;
  }
  async save(
    data: { entity: any; options?: SaveOptions },
    queryRunner?: QueryRunner,
  ): Promise<T> {
    const internalQueryRunner = !queryRunner;
    queryRunner = this.getQueryRunner(queryRunner);
    try {
      return await queryRunner.manager.save(
        this.entityClass,
        data.entity,
        data.options,
      );
    } finally {
      if (internalQueryRunner && !queryRunner.isReleased) {
        queryRunner.release();
      }
    }
  }
  async update(
    criteria: CriteriaType<T>,
    partialEntity: QueryDeepPartialEntity<T>,
    queryRunner?: QueryRunner,
  ): Promise<UpdateResult> {
    const internalQueryRunner = !queryRunner;
    queryRunner = this.getQueryRunner(queryRunner);
    try {
      return await queryRunner.manager.update(
        this.entityClass,
        criteria,
        partialEntity,
      );
    } finally {
      if (internalQueryRunner && !queryRunner.isReleased) {
        queryRunner.release();
      }
    }
  }
  async count(
    options?: FindManyOptions<T>,
    queryRunner?: QueryRunner,
  ): Promise<number> {
    const internalQueryRunner = !queryRunner;
    queryRunner = this.getQueryRunner(queryRunner);
    try {
      return await queryRunner.manager.count(this.entityClass, options);
    } finally {
      if (internalQueryRunner && !queryRunner.isReleased) {
        queryRunner.release();
      }
    }
  }
  async softDelete(
    criteria: CriteriaType<T>,
    queryRunner?: QueryRunner,
  ): Promise<UpdateResult> {
    const internalQueryRunner = !queryRunner;
    queryRunner = this.getQueryRunner(queryRunner);
    try {
      return await queryRunner.manager.softDelete(this.entityClass, criteria);
    } finally {
      if (internalQueryRunner && !queryRunner.isReleased) {
        queryRunner.release();
      }
    }
  }
  async restore(
    criteria: CriteriaType<T>,
    queryRunner?: QueryRunner,
  ): Promise<UpdateResult> {
    const internalQueryRunner = !queryRunner;
    queryRunner = this.getQueryRunner(queryRunner);
    try {
      return await queryRunner.manager.restore(this.entityClass, criteria);
    } finally {
      if (internalQueryRunner && !queryRunner.isReleased) {
        queryRunner.release();
      }
    }
  }
  async saveMany(
    entities: T[],
    options?: SaveOptions,
    queryRunner?: QueryRunner,
  ): Promise<T[]> {
    const internalQueryRunner = !queryRunner;
    queryRunner = this.getQueryRunner(queryRunner);
    try {
      return await queryRunner.manager.save(
        this.entityClass,
        entities,
        options,
      );
    } finally {
      if (internalQueryRunner && !queryRunner.isReleased) {
        queryRunner.release();
      }
    }
  }
  createMany(entityLikes: DeepPartial<T>[], queryRunner?: QueryRunner): T[] {
    const internalQueryRunner = !queryRunner;
    queryRunner = this.getQueryRunner(queryRunner);
    const entities = queryRunner.manager.create(this.entityClass, entityLikes);
    if (internalQueryRunner && !queryRunner.isReleased) {
      queryRunner.release();
    }
    return entities;
  }
  async findAndCount(
    options?: FindManyOptions<T>,
    queryRunner?: QueryRunner,
  ): Promise<[T[], number]> {
    const internalQueryRunner = !queryRunner;
    queryRunner = this.getQueryRunner(queryRunner);
    try {
      return await queryRunner.manager.findAndCount(this.entityClass, options);
    } finally {
      if (internalQueryRunner && !queryRunner.isReleased) {
        queryRunner.release();
      }
    }
  }
  async createQueryBuilder(
    alias: string,
    queryRunner?: QueryRunner,
  ): Promise<SelectQueryBuilder<T>> {
    const internalQueryRunner = !queryRunner;
    queryRunner = this.getQueryRunner(queryRunner);
    try {
      return queryRunner.manager.createQueryBuilder<T>(this.entityClass, alias);
    } finally {
      if (internalQueryRunner && !queryRunner.isReleased) {
        queryRunner.release();
      }
    }
  }
}
