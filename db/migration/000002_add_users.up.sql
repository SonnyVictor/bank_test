CREATE TABLE "users" (
    "username" varchar PRIMARY KEY,
    "hashed_password" varchar NOT NULL,
    "full_name" varchar NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01',
    "created_at" timestamptz NOT NULL DEFAULT (now ())
);

-- CREATE INDEX ON "accounts" ("owner");
ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

-- CREATE UNIQUE INDEX ON "accounts" ("owner", "currency");
ALTER TABLE "accounts" ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner", "currency");
