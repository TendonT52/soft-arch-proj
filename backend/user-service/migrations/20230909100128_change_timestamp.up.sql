-- Step 1: Create new columns with type BIGINT
-- Add new columns to temporarily store the Unix timestamps as BIGINT.
ALTER TABLE users
ADD COLUMN created_at_bigint BIGINT;

ALTER TABLE users
ADD COLUMN updated_at_bigint BIGINT;

ALTER TABLE students
ADD COLUMN created_at_bigint BIGINT;

ALTER TABLE students
ADD COLUMN updated_at_bigint BIGINT;

ALTER TABLE companies
ADD COLUMN created_at_bigint BIGINT;

ALTER TABLE companies
ADD COLUMN updated_at_bigint BIGINT;

-- Step 2: Update the new columns with Unix timestamps
-- Convert and copy the existing data from the old columns to the new BIGINT columns.
UPDATE users
SET created_at_bigint = EXTRACT(EPOCH FROM created_at::TIMESTAMP);

UPDATE users
SET updated_at_bigint = EXTRACT(EPOCH FROM updated_at::TIMESTAMP);

UPDATE students
SET created_at_bigint = EXTRACT(EPOCH FROM created_at::TIMESTAMP);

UPDATE students
SET updated_at_bigint = EXTRACT(EPOCH FROM updated_at::TIMESTAMP);

UPDATE companies
SET created_at_bigint = EXTRACT(EPOCH FROM created_at::TIMESTAMP);

UPDATE companies
SET updated_at_bigint = EXTRACT(EPOCH FROM updated_at::TIMESTAMP);

-- Step 3: Drop the old columns (if desired)
-- If you want to keep the old columns, you can skip this step.
ALTER TABLE users
DROP COLUMN created_at;

ALTER TABLE users
DROP COLUMN updated_at;

ALTER TABLE students
DROP COLUMN created_at;

ALTER TABLE students
DROP COLUMN updated_at;

ALTER TABLE companies
DROP COLUMN created_at;

ALTER TABLE companies
DROP COLUMN updated_at;

-- Step 4: Rename the new columns to match the original column names
ALTER TABLE users
RENAME COLUMN created_at_bigint TO created_at;

ALTER TABLE users
RENAME COLUMN updated_at_bigint TO updated_at;

ALTER TABLE students
RENAME COLUMN created_at_bigint TO created_at;

ALTER TABLE students
RENAME COLUMN updated_at_bigint TO updated_at;

ALTER TABLE companies
RENAME COLUMN created_at_bigint TO created_at;

ALTER TABLE companies
RENAME COLUMN updated_at_bigint TO updated_at;