CREATE TYPE "priority_enum" AS ENUM (
  'low',
  'medium',
  'high'
);

CREATE TABLE "users" (
  "user_id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz DEFAULT 'now()'
);

CREATE TABLE "tasks" (
  "task_id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "task_name" varchar NOT NULL,
  "description" text,
  "due_date" timestamptz,
  "priority" priority_enum,
  "created_at" timestamptz DEFAULT 'now()'
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "tasks" ("user_id");

ALTER TABLE "tasks" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");