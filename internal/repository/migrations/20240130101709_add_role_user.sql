-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE "USER"
    ALTER COLUMN role SET DEFAULT 'WORKER';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE "USER"
    ALTER COLUMN role DROP DEFAULT;
-- +goose StatementEnd