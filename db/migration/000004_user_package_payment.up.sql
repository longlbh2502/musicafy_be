-- Thêm các cột mới vào bảng users
ALTER TABLE "users" 
  ADD COLUMN IF NOT EXISTS "package" varchar,
  ADD COLUMN IF NOT EXISTS "package_expire" timestamptz;

ALTER TABLE "users" ALTER COLUMN "package" DROP NOT NULL;
UPDATE "users" SET "package" = NULL;


-- Tạo bảng packages
CREATE TABLE "packages" (
  "id" serial PRIMARY KEY,
  "thumb" varchar,
  "title" varchar,
  "code" varchar UNIQUE,
  "level" int DEFAULT 1,
  "description" varchar
);

-- Tạo bảng package_price
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

-- Tạo bảng payment
CREATE TABLE "payment" (
  "id" serial PRIMARY KEY,
  "code" varchar NOT NULL UNIQUE,
  "user" serial
);

-- Tạo bảng transaction
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

-- Thêm các foreign key
ALTER TABLE "users" ADD FOREIGN KEY ("package") REFERENCES "packages" ("code") ON DELETE SET NULL;

ALTER TABLE "package_price" ADD FOREIGN KEY ("package") REFERENCES "packages" ("code") ON DELETE CASCADE;

ALTER TABLE "payment" ADD FOREIGN KEY ("user") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "transaction" ADD FOREIGN KEY ("payment") REFERENCES "payment" ("code") ON DELETE CASCADE;
ALTER TABLE "transaction" ADD FOREIGN KEY ("package") REFERENCES "packages" ("code") ON DELETE SET NULL;