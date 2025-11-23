-- Drop user_modules first because it depends on modules
DROP TABLE IF EXISTS user_modules;

-- Drop modules table
DROP TABLE IF EXISTS modules;

-- Drop indexes on users (optional, they will be dropped with the table automatically)
DROP INDEX IF EXISTS idx_users_uuid;
DROP INDEX IF EXISTS idx_users_name_lower;

-- Drop users table
DROP TABLE IF EXISTS users;