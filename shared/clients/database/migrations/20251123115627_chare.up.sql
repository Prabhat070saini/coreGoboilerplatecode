CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(100),
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    is_blocked BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by INT DEFAULT -1,
    modified_at TIMESTAMP NULL,
    modified_by INT NULL,
    deleted_at TIMESTAMPTZ NULL,
    deleted_by INT NULL
);

-- Index on UUID
CREATE UNIQUE INDEX IF NOT EXISTS idx_users_uuid ON "users"(uuid);

-- Index on name for faster lookup
CREATE INDEX IF NOT EXISTS idx_users_name_lower ON "users"(LOWER(name));

-- Create modules table
CREATE TABLE modules (
    id BIGSERIAL PRIMARY KEY,           -- ✅ Use BIGSERIAL instead of AUTO_INCREMENT
    code CHAR(6) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    parent_id BIGINT DEFAULT NULL,
    description VARCHAR(255),
    FOREIGN KEY (parent_id) REFERENCES modules(id) ON DELETE CASCADE
);

-- Create user_modules table
CREATE TABLE user_modules (
    user_id BIGINT NOT NULL,
    module_id BIGINT NOT NULL,
    PRIMARY KEY (user_id, module_id),
    FOREIGN KEY (module_id) REFERENCES modules(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);