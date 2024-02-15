-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE IF EXISTS "DUTY" ADD COLUMN "amount" integer;
ALTER TABLE IF EXISTS "DUTY" ADD CONSTRAINT course_range_check CHECK (amount >= 1 AND amount <= 2);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE IF EXISTS "DUTY" DROP COLUMN IF EXISTS "amount";
-- +goose StatementEnd