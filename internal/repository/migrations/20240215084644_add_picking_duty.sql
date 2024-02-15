-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE IF EXISTS "DUTY" ADD COLUMN "picking_time_start" timestamp;
ALTER TABLE IF EXISTS "DUTY" ADD COLUMN "picking_time_end" timestamp;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE IF EXISTS "DUTY" DROP COLUMN IF EXISTS "picking_time_start";
ALTER TABLE IF EXISTS "DUTY" DROP COLUMN IF EXISTS "picking_time_end";
-- +goose StatementEnd