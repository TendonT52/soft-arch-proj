CREATE TYPE roleType AS ENUM ('student', 'company');

CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL,
    "email" varchar NOT NULL UNIQUE,
    "password" varchar NOT NULL,
    "role" roleType NOT NULL,
    "verified" boolean NOT NULL DEFAULT false,
    "description" varchar NOT NULL
);

CREATE TABLE "students" (
    "sid" bigserial PRIMARY KEY REFERENCES users(id),
    "verification_code" varchar,
    "faculty" varchar NOT NULL,
    "major" varchar NOT NULL,
    "year" int NOT NULL
);

CREATE TABLE "companies" (
    "cid" bigserial PRIMARY KEY REFERENCES users(id),
    "location" varchar NOT NULL,
    "phone" varchar NOT NULL,
    "category" varchar NOT NULL
);