-- +goose Up
-- +goose StatementBegin
ALTER TABLE post_lang ADD COLUMN description TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE post_lang DROP COLUMN description;
-- +goose StatementEnd
