-- +goose Up
-- +goose StatementBegin
ALTER TABLE task ADD COLUMN is_archive BOOLEAN DEFAULT FALSE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE task DROP COLUMN is_archive;
-- +goose StatementEnd