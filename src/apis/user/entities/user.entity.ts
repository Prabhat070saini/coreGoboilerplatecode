import {
  Entity,
  PrimaryGeneratedColumn,
  Column,
  CreateDateColumn,
  UpdateDateColumn,
  Index,
  BeforeInsert,
} from 'typeorm';
import { ulid } from 'ulid';
import { ESchema } from 'src/common/constants/app.enum';
@Entity({ name: 'users', schema: ESchema.DBO_SCHEMA })
@Index(['email'], { unique: true })
@Index(['user_id'], { unique: true })
export class User {
  @PrimaryGeneratedColumn()
  id: number;

  @Column({ type: 'varchar', length: 26, unique: true })
  user_id: string;

  @Column({ type: 'varchar', length: 255, nullable: true })
  name?: string;

  @Column({ type: 'varchar', length: 255, unique: true })
  email: string;

  @Column({ type: 'varchar', length: 15, nullable: true })
  phone?: string;

  @Column({ type: 'varchar', length: 15, nullable: true })
  role?: string;

  @Column({ type: 'boolean', default: true })
  is_active: boolean;

  @Column({ type: 'text' })
  password: string;

  @Column({ type: 'int' })
  created_by: number;

  @Column({ type: 'timestamp', nullable: true })
  last_login?: Date;

  @CreateDateColumn({ type: 'timestamp' })
  created_at: Date;

  @UpdateDateColumn({ type: 'timestamp', nullable: true })
  updated_at: Date;
  @Column({ type: 'int', nullable: true })
  updated_by: number;
  @BeforeInsert()
  generateUlid(): string {
    this.user_id = ulid();
    return this.user_id;
  }
}
