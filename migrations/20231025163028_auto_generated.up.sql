-- create "client_creds" table
CREATE TABLE "client_creds" ("id" bigserial NOT NULL, "slug" text NOT NULL, "hashed_secret" text NOT NULL, "is_blocked" boolean NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- create index "idx_client_creds_slug" to table: "client_creds"
CREATE UNIQUE INDEX "idx_client_creds_slug" ON "client_creds" ("slug");
-- create "furnitures" table
CREATE TABLE "furnitures" ("id" bigserial NOT NULL, "type" text NOT NULL, "name" text NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "deleted_at" timestamptz NULL, PRIMARY KEY ("id"));
-- create index "idx_furnitures_deleted_at" to table: "furnitures"
CREATE INDEX "idx_furnitures_deleted_at" ON "furnitures" ("deleted_at");
