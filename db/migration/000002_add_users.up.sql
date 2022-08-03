CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamp with time zone NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamp with time zone NOT NULL DEFAULT (now())
);
-- COMMENT ON COLUMN "users"."password_changed_at" IS '默认零时区，提示用户修改密码使用';

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

-- CREATE UNIQUE INDEX ON "accounts" ("owner", "currency");
ALTER TABLE "accounts" ADD CONSTRAINT "owner_currency_key" UNIQUE("owner" , "currency");