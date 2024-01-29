-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE IF EXISTS "USER" ADD COLUMN "course" integer;
ALTER TABLE IF EXISTS "USER" ADD CONSTRAINT course_range_check CHECK (course >= 1 AND course <= 10);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE IF EXISTS "USER" DROP COLUMN IF EXISTS "course";
-- +goose StatementEnd