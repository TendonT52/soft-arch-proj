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

INSERT INTO "public"."companies" ("cid", "location", "phone", "category", "created_at", "updated_at", "status", "name", "description") VALUES
(33, '099/999 Silom Road, Bangkok', '312-345-6789', 'Innovative and Security', '2023-09-01 04:43:33.575911', '2023-09-01 04:49:32.955583', 'Approve', 'Soft Company', 'The Best Company'),
(100, 'New York', '123-456-7890', 'Technology and Better', '2023-08-31 14:35:05.4137', '2023-08-31 14:35:05.4137', 'Approve', 'Ant Company', 'Empowering Tomorrow, Today.'),
(101, 'New York', '123-456-7890', 'Banking and Finance', '2023-08-31 14:35:05.4137', '2023-08-31 14:35:05.4137', 'Approve', 'B Company', 'Innovating for a Brighter Future.'),
(102, 'New York', '123-456-7890', 'Bank and Technology', '2023-08-31 14:35:05.4137', '2023-08-31 14:35:05.4137', 'Approve', 'C Company', 'Leading the Way to Transformative Change.'),
(103, 'New York', '123-456-7890', 'Innovative and Tech', '2023-08-31 14:35:05.4137', '2023-08-31 14:35:05.4137', 'Approve', 'D Bank Company', 'Building a Better World.'),
(104, 'New York', '123-456-7890', 'Business and Finance', '2023-08-31 14:35:05.4137', '2023-08-31 14:35:05.4137', 'Approve', 'E Company', 'Creating a Better Tomorrow.');

INSERT INTO "public"."schema_migrations" ("version", "dirty") VALUES
(20230901050401, 'f');

INSERT INTO "public"."students" ("sid", "verification_code", "faculty", "major", "year", "created_at", "updated_at", "name", "description") VALUES
(32, 'Vmp5QlB0cnJPcWhab25iblNFMG0=', 'Engineering 1234', 'Computer Engineering 1234', 3, '2023-09-01 04:26:33.710955', '2023-09-01 04:26:33.710955', 'Platoo Tepsut', 'Best student');

INSERT INTO "public"."users" ("id", "email", "password", "role", "verified", "created_at", "updated_at") VALUES
(31, 'admin@gmail.com', '$2a$10$yGkA68nAWsGgP8F0kVtPhekIGsQp.Q0Mgi1vU3ezQGKLHKqW0diBe', 'admin', 't', '2023-09-01 04:23:17.65823', '2023-09-01 04:23:17.65823'),
(32, '6330203521@student.chula.ac.th', '$2a$10$6jr4hNrhzjNTMtjD4VhEne4zu97Kyi0oa8w0fnkpVTQIFlAR.bJ5C', 'student', 't', '2023-09-01 04:26:33.710955', '2023-09-01 04:27:32.488492'),
(33, 'tikhamporntan@gmail.com', '$2a$10$ZMljgBCsG1LjKp7giB2DquD6vaQn0a43SeLcreHO37JGVy3.VLhOm', 'company', 't', '2023-09-01 04:43:33.575911', '2023-09-01 04:49:32.955583'),
(100, 'a@gmail.com', 'fake-password-a', 'company', 't', '2023-08-31 14:35:05.4137', '2023-08-31 14:35:05.4137'),
(101, 'b@gmail.com', 'fake-password-b', 'company', 't', '2023-08-31 14:35:05.4137', '2023-08-31 14:35:05.4137'),
(102, 'c@gmail.com', 'fake-password-c', 'company', 't', '2023-08-31 14:35:05.4137', '2023-08-31 14:35:05.4137'),
(103, 'd@gmail.com', 'fake-password-d', 'company', 't', '2023-08-31 14:35:05.4137', '2023-08-31 14:35:05.4137'),
(104, 'e@gmail.com', 'fake-password-e', 'company', 't', '2023-08-31 14:35:05.4137', '2023-08-31 14:35:05.4137');

ALTER TABLE "public"."companies" ADD FOREIGN KEY ("cid") REFERENCES "public"."users"("id");
ALTER TABLE "public"."students" ADD FOREIGN KEY ("sid") REFERENCES "public"."users"("id");
