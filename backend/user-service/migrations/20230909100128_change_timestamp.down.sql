-- Step 1: Rename the new columns to match the original column names
ALTER TABLE users
RENAME COLUMN created_at TO created_at_bigint;

ALTER TABLE users
RENAME COLUMN updated_at TO updated_at_bigint;

ALTER TABLE students
RENAME COLUMN created_at TO created_at_bigint;

ALTER TABLE students
RENAME COLUMN updated_at TO updated_at_bigint;

ALTER TABLE companies
RENAME COLUMN created_at TO created_at_bigint;

ALTER TABLE companies
RENAME COLUMN updated_at TO updated_at_bigint;

-- Step 2: Recreate the old columns if they were dropped (adjust data types as needed)

ALTER TABLE users
ADD COLUMN created_at TIMESTAMP;

ALTER TABLE users
ADD COLUMN updated_at TIMESTAMP;

ALTER TABLE students
ADD COLUMN created_at TIMESTAMP;

ALTER TABLE students
ADD COLUMN updated_at TIMESTAMP;

ALTER TABLE companies
ADD COLUMN created_at TIMESTAMP;

ALTER TABLE companies
ADD COLUMN updated_at TIMESTAMP;

-- Step 3: Drop the new columns with type BIGINT
ALTER TABLE users
DROP COLUMN created_at_bigint;

ALTER TABLE users
DROP COLUMN updated_at_bigint;

ALTER TABLE students
DROP COLUMN created_at_bigint;

ALTER TABLE students
DROP COLUMN updated_at_bigint;

ALTER TABLE companies
DROP COLUMN created_at_bigint;

ALTER TABLE companies
DROP COLUMN updated_at_bigint;
