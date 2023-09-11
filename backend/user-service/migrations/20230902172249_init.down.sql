-- Revert foreign key constraints
ALTER TABLE "public"."students" DROP CONSTRAINT IF EXISTS "students_sid_fkey";
ALTER TABLE "public"."companies" DROP CONSTRAINT IF EXISTS "companies_cid_fkey";

-- Drop tables
DROP TABLE IF EXISTS "public"."companies";
DROP TABLE IF EXISTS "public"."students";
DROP TABLE IF EXISTS "public"."users";

-- Drop sequences and types
DROP SEQUENCE IF EXISTS students_sid_seq;
DROP SEQUENCE IF EXISTS companies_cid_seq;
DROP SEQUENCE IF EXISTS users_id_seq;
DROP TYPE IF EXISTS "public"."roletype";
DROP TYPE IF EXISTS "public"."statustype";