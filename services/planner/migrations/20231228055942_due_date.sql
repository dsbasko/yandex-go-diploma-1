-- +goose Up
-- +goose StatementBegin
ALTER TABLE task ADD COLUMN due_date TIMESTAMPTZ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE task DROP COLUMN due_date;
-- +goose StatementEnd
