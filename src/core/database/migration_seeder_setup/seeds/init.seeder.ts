import { DataSource } from 'typeorm';
import { runSeeders } from 'typeorm-extension';
import { Logger } from '@nestjs/common';

export default class InitSeeder {
  private readonly logger = new Logger(InitSeeder.name);
  public async run(
    dataSource: DataSource,
    // factoryManager: SeederFactoryManager,
  ): Promise<void> {
    const seeds = [
      /* add your seeders here */
    ];

    // Only run if seeds array is not empty
    if (seeds.length === 0) {
      this.logger.log('No seeders to run. Skipping seeding...');
      return;
    }

    await runSeeders(dataSource, {
      seeds,
      factories: [],
    });

    this.logger.log('Seeding completed.');
  }
}
