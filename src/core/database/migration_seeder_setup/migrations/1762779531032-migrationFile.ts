import { MigrationInterface, QueryRunner } from 'typeorm';

export class MigrationFile1762779531032 implements MigrationInterface {
  name = 'MigrationFile1762779531032';

  public async up(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(
      `CREATE TABLE "dbo"."users" ("id" SERIAL NOT NULL, "user_id" character varying(26) NOT NULL, "name" character varying(255), "email" character varying(255) NOT NULL, "phone" character varying(15), "role" character varying(15), "is_active" boolean NOT NULL DEFAULT true, "password" text NOT NULL, "created_by" integer NOT NULL, "last_login" TIMESTAMP, "created_at" TIMESTAMP NOT NULL DEFAULT now(), "updated_at" TIMESTAMP DEFAULT now(), "updated_by" integer, CONSTRAINT "UQ_96aac72f1574b88752e9fb00089" UNIQUE ("user_id"), CONSTRAINT "UQ_97672ac88f789774dd47f7c8be3" UNIQUE ("email"), CONSTRAINT "PK_a3ffb1c0c8416b9fc6f907b7433" PRIMARY KEY ("id"))`,
    );
    await queryRunner.query(
      `CREATE UNIQUE INDEX "IDX_96aac72f1574b88752e9fb0008" ON "dbo"."users" ("user_id") `,
    );
    await queryRunner.query(
      `CREATE UNIQUE INDEX "IDX_97672ac88f789774dd47f7c8be" ON "dbo"."users" ("email") `,
    );
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(
      `DROP INDEX "dbo"."IDX_97672ac88f789774dd47f7c8be"`,
    );
    await queryRunner.query(
      `DROP INDEX "dbo"."IDX_96aac72f1574b88752e9fb0008"`,
    );
    await queryRunner.query(`DROP TABLE "dbo"."users"`);
  }
}
