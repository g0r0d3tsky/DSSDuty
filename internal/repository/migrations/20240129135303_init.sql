-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE EXTENSION IF NOT EXISTS citext;
CREATE TABLE IF NOT EXISTS "USER"
(
    "id"            UUID PRIMARY KEY,
    "created_at"    timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    "email"         citext UNIQUE               NOT NULL,
    "password_hash" bytea                       NOT NULL,
    "activated"     bool                        NOT NULL,
    "full_name"     VARCHAR(255)                NOT NULL,
    "duty_amount"   int,
    "role"          VARCHAR(255)                NOT NULL
);

CREATE TABLE IF NOT EXISTS "DUTY"
(
    "id"       UUID PRIMARY KEY,
    "date"     TIMESTAMP NOT NULL,
    "user_id1" UUID,
    "user_id2" UUID
);

CREATE TABLE IF NOT EXISTS "STIMULATION"
(
    "id"        UUID PRIMARY KEY,
    "user_id"   UUID                        NOT NULL,
    "rewards"   int,
    "sanctions" int,
    "info"      VARCHAR(255),
    "timestamp" timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "EVENT"
(
    "id"        UUID PRIMARY KEY,
    "user_id"   UUID         NOT NULL,
    "type"      VARCHAR(255) NOT NULL,
    "timestamp" TIMESTAMP    NOT NULL
);

ALTER TABLE "EVENT"
    ADD FOREIGN KEY ("user_id") REFERENCES "USER" ("id");

ALTER TABLE "STIMULATION"
    ADD FOREIGN KEY ("user_id") REFERENCES "USER" ("id");

ALTER TABLE "DUTY"
    ADD FOREIGN KEY ("user_id1") REFERENCES "USER" ("id");

ALTER TABLE "DUTY"
    ADD FOREIGN KEY ("user_id2") REFERENCES "USER" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE IF EXISTS "DUTY";
DROP TABLE IF EXISTS "EVENT";
DROP TABLE IF EXISTS "STIMULATION";
DROP TABLE IF EXISTS "USER";
-- +goose StatementEnd