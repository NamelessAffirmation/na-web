-- ############################
-- NamelessAffirmation Official Schema
--
-- https://namelessaffirmation.com
--
-- Licensed Under GNU GPLv3.
-- Copyright 2024. NamelessAffirmation. All Rights Reserved.
-- ############################

-- AUTH

CREATE TABLE IF NOT EXISTS "user_types" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "created_at" timestamp DEFAULT (now()),
  "modified_at" timestamp,
  "is_archived" boolean DEFAULT false,
  "name" text NOT NULL,
  "key" text UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS "users" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "created_at" timestamp DEFAULT (now()),
  "modified_at" timestamp,
  "is_archived" bool DEFAULT false,
  "email" text UNIQUE NOT NULL,
  "is_verified" boolean DEFAULT false,
  "user_type_id" integer,
  "password_hash" text,
  "last_login" timestamp,
  "is_banned" boolean DEFAULT false,
  "ban_reason" text
);

ALTER TABLE "users" ADD FOREIGN KEY ("user_type_id") REFERENCES "user_types" ("id");

-- // USER TYPES
INSERT INTO public.user_types (name, key)
VALUES ('Admin', 'admin')
ON CONFLICT (id) DO NOTHING;

INSERT INTO public.user_types (name, key)
VALUES ('Free', 'free')
ON CONFLICT (id) DO NOTHING;

INSERT INTO public.user_types (name, key)
VALUES ('Pro', 'pro')
ON CONFLICT (id) DO NOTHING;

-- // USERS
INSERT INTO public.users (email, user_type_id)
VALUES ('snowlynxsoftware+admin@gmail.com', 1)
ON CONFLICT (id) DO NOTHING;