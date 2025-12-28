-- +goose Up
-- +goose StatementBegin
ALTER TABLE books ADD COLUMN likes BIGINT NOT NULL DEFAULT 0 CHECK (likes >= 0);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE books DROP COLUMN likes;
-- +goose StatementEnd
