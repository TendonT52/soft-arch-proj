-- -------------------------------------------------------------
-- TablePlus 5.3.8(500)
--
-- https://tableplus.com/
--
-- Database: golang-auth
-- Generation Time: 2566-09-03 00:42:53.8910
-- -------------------------------------------------------------


DROP TABLE IF EXISTS "public"."companies";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS companies_cid_seq;
DROP TYPE IF EXISTS "public"."statustype";
CREATE TYPE "public"."statustype" AS ENUM ('Approve', 'Reject', 'Pending');

-- Table Definition
CREATE TABLE "public"."companies" (
    "cid" int8 NOT NULL DEFAULT nextval('companies_cid_seq'::regclass),
    "location" varchar NOT NULL,
    "phone" varchar NOT NULL,
    "category" varchar NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "status" "public"."statustype" DEFAULT 'Pending'::statustype,
    "name" varchar NOT NULL,
    "description" varchar NOT NULL,
    PRIMARY KEY ("cid")
);

DROP TABLE IF EXISTS "public"."schema_migrations";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."schema_migrations" (
    "version" int8 NOT NULL,
    "dirty" bool NOT NULL,
    PRIMARY KEY ("version")
);

DROP TABLE IF EXISTS "public"."students";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS students_sid_seq;

-- Table Definition
CREATE TABLE "public"."students" (
    "sid" int8 NOT NULL DEFAULT nextval('students_sid_seq'::regclass),
    "verification_code" varchar,
    "faculty" varchar NOT NULL,
    "major" varchar NOT NULL,
    "year" int4 NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "name" varchar NOT NULL,
    "description" varchar NOT NULL,
    PRIMARY KEY ("sid")
);

DROP TABLE IF EXISTS "public"."users";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS users_id_seq;
DROP TYPE IF EXISTS "public"."roletype";
CREATE TYPE "public"."roletype" AS ENUM ('student', 'company', 'admin');

-- Table Definition
CREATE TABLE "public"."users" (
    "id" int8 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    "email" varchar NOT NULL,
    "password" varchar NOT NULL,
    "role" "public"."roletype" NOT NULL,
    "verified" bool NOT NULL DEFAULT false,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);

ALTER TABLE "public"."companies" ADD FOREIGN KEY ("cid") REFERENCES "public"."users"("id");
ALTER TABLE "public"."students" ADD FOREIGN KEY ("sid") REFERENCES "public"."users"("id");
