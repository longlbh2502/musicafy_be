-- SQL dump generated using DBML (dbml.dbdiagram.io)
-- Database: PostgreSQL
-- Generated at: 2025-03-29T02:33:28.119Z

CREATE TYPE "gender" AS ENUM (
  'nam',
  'nữ',
  'khác'
);

CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "is_verify" boolean NOT NULL DEFAULT false,
  "email" varchar,
  "gender" gender,
  "licence" varchar,
  "dob" timestamp,
  "active" bool NOT NULL DEFAULT true,
  "avatar" varchar(255),
  "package" varchar,
  "package_expire" timestamptz,
  "updated_at" timestamptz,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "packages" (
  "id" serial PRIMARY KEY,
  "thumb" varchar,
  "title" varchar,
  "code" varchar UNIQUE,
  "level" int DEFAULT 1,
  "description" varchar
);

CREATE TABLE "package_price" (
  "id" serial PRIMARY KEY,
  "title" varchar,
  "code" varchar,
  "recommend" bool DEFAULT false,
  "price" float DEFAULT 0,
  "description" varchar,
  "package" varchar,
  "duration" int DEFAULT 0
);

CREATE TABLE "payment" (
  "id" serial PRIMARY KEY,
  "code" varchar UNIQUE NOT NULL,
  "user" serial
);

CREATE TABLE "transaction" (
  "id" serial PRIMARY KEY,
  "code" varchar,
  "payment" varchar,
  "package" varchar,
  "value" float DEFAULT 0,
  "discount" float DEFAULT 0,
  "is_pair" bool DEFAULT false,
  "pair_at" timestamp,
  "log_preferred" jsonb,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "sessions" (
  "id" serial PRIMARY KEY,
  "username" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "verifies" (
  "id" serial PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "secret_code" varchar NOT NULL,
  "is_used" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "expired_at" timestamptz NOT NULL DEFAULT ((now() + interval '15 minutes'))
);

CREATE TABLE "songs" (
  "id" serial PRIMARY KEY,
  "mask_id" varchar(255) NOT NULL,
  "title" varchar(255) NOT NULL,
  "alias" varchar(255),
  "is_offical" bool DEFAULT true,
  "thumbnail" varchar(255),
  "thumbnailM" varchar(255),
  "duration" int NOT NULL DEFAULT 0,
  "releaseDate" int,
  "album" serial,
  "distributor" varchar(255),
  "hasLyric" bool NOT NULL DEFAULT false,
  "like" int NOT NULL DEFAULT 0,
  "listen" int NOT NULL DEFAULT 0,
  "comment" int NOT NULL DEFAULT 0,
  "audio_file" varchar
);

CREATE TABLE "artists" (
  "id" serial PRIMARY KEY,
  "mask_id" varchar(255) NOT NULL,
  "name" varchar(255),
  "spotlight" boolean,
  "alias" varchar(255),
  "thumbnail" varchar(255),
  "thumbnailM" varchar(255),
  "playlistId" varchar(20),
  "totalFollow" int NOT NULL DEFAULT 0
);

CREATE TABLE "composers" (
  "id" serial PRIMARY KEY,
  "mask_id" varchar(255) NOT NULL,
  "name" varchar(255),
  "spotlight" boolean,
  "alias" varchar(255),
  "thumbnail" varchar(255),
  "cover" varchar(255)
);

CREATE TABLE "albums" (
  "id" serial PRIMARY KEY,
  "mask_id" varchar(255) NOT NULL,
  "title" varchar(255) NOT NULL,
  "is_offical" bool DEFAULT true,
  "thumbnail" varchar(255),
  "sortDescription" text,
  "release_at" int
);

CREATE TABLE "genres" (
  "id" serial PRIMARY KEY,
  "mask_id" varchar(255) NOT NULL,
  "title" varchar(255) NOT NULL,
  "name" varchar(255) NOT NULL,
  "alias" varchar(255)
);

CREATE TABLE "song_file_mp3" (
  "id" serial PRIMARY KEY,
  "title" varchar(255),
  "url" varchar(255) NOT NULL,
  "description" varchar(255),
  "is_vip" bool DEFAULT false,
  "song" serial
);

CREATE TABLE "song_lyric" (
  "id" serial PRIMARY KEY,
  "song" serial UNIQUE,
  "file" varchar(255),
  "data" jsonb
);

CREATE TABLE "play_list" (
  "id" serial PRIMARY KEY,
  "title" varchar(255),
  "users" serial,
  "thumbnail" varchar,
  "sort_description" varchar,
  "created_at" timestamp DEFAULT (now()),
  "is_private" bool DEFAULT false
);

COMMENT ON COLUMN "packages"."thumb" IS 'ảnh cover';

COMMENT ON COLUMN "package_price"."duration" IS 'đơn vị giây';

COMMENT ON COLUMN "payment"."code" IS 'PM_{user.id}_{random}';

COMMENT ON COLUMN "transaction"."code" IS 'mã giao dịch';

ALTER TABLE "users" ADD FOREIGN KEY ("package") REFERENCES "packages" ("code") ON DELETE SET NULL;

ALTER TABLE "package_price" ADD FOREIGN KEY ("package") REFERENCES "packages" ("code") ON DELETE CASCADE;

ALTER TABLE "payment" ADD FOREIGN KEY ("user") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "transaction" ADD FOREIGN KEY ("payment") REFERENCES "payment" ("code") ON DELETE CASCADE;

ALTER TABLE "transaction" ADD FOREIGN KEY ("package") REFERENCES "packages" ("code") ON DELETE SET NULL;

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username") ON DELETE CASCADE;

ALTER TABLE "verifies" ADD FOREIGN KEY ("username") REFERENCES "users" ("username") ON DELETE CASCADE;

CREATE TABLE "songs_artists" (
  "songs_id" serial,
  "artists_id" serial,
  PRIMARY KEY ("songs_id", "artists_id")
);

ALTER TABLE "songs_artists" ADD FOREIGN KEY ("songs_id") REFERENCES "songs" ("id");

ALTER TABLE "songs_artists" ADD FOREIGN KEY ("artists_id") REFERENCES "artists" ("id");


CREATE TABLE "songs_genres" (
  "songs_id" serial,
  "genres_id" serial,
  PRIMARY KEY ("songs_id", "genres_id")
);

ALTER TABLE "songs_genres" ADD FOREIGN KEY ("songs_id") REFERENCES "songs" ("id");

ALTER TABLE "songs_genres" ADD FOREIGN KEY ("genres_id") REFERENCES "genres" ("id");


CREATE TABLE "songs_composers" (
  "songs_id" serial,
  "composers_id" serial,
  PRIMARY KEY ("songs_id", "composers_id")
);

ALTER TABLE "songs_composers" ADD FOREIGN KEY ("songs_id") REFERENCES "songs" ("id");

ALTER TABLE "songs_composers" ADD FOREIGN KEY ("composers_id") REFERENCES "composers" ("id");


ALTER TABLE "songs" ADD FOREIGN KEY ("album") REFERENCES "albums" ("id") ON DELETE SET NULL;

CREATE TABLE "albums_artists" (
  "albums_id" serial,
  "artists_id" serial,
  PRIMARY KEY ("albums_id", "artists_id")
);

ALTER TABLE "albums_artists" ADD FOREIGN KEY ("albums_id") REFERENCES "albums" ("id");

ALTER TABLE "albums_artists" ADD FOREIGN KEY ("artists_id") REFERENCES "artists" ("id");


CREATE TABLE "albums_genres" (
  "albums_id" serial,
  "genres_id" serial,
  PRIMARY KEY ("albums_id", "genres_id")
);

ALTER TABLE "albums_genres" ADD FOREIGN KEY ("albums_id") REFERENCES "albums" ("id");

ALTER TABLE "albums_genres" ADD FOREIGN KEY ("genres_id") REFERENCES "genres" ("id");


ALTER TABLE "song_file_mp3" ADD FOREIGN KEY ("song") REFERENCES "songs" ("id") ON DELETE CASCADE;

ALTER TABLE "song_lyric" ADD FOREIGN KEY ("song") REFERENCES "songs" ("id") ON DELETE CASCADE;

CREATE TABLE "play_list_songs" (
  "play_list_id" serial,
  "songs_id" serial,
  PRIMARY KEY ("play_list_id", "songs_id")
);

ALTER TABLE "play_list_songs" ADD FOREIGN KEY ("play_list_id") REFERENCES "play_list" ("id");

ALTER TABLE "play_list_songs" ADD FOREIGN KEY ("songs_id") REFERENCES "songs" ("id");


ALTER TABLE "play_list" ADD FOREIGN KEY ("users") REFERENCES "users" ("id") ON DELETE CASCADE;
