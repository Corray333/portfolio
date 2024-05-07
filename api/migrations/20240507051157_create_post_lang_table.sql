-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS post_lang (
    post_id BIGINT NOT NULL,
    lang VARCHAR(5) NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    FOREIGN KEY (post_id) REFERENCES posts(post_id),
    PRIMARY KEY (post_id, lang)
);
CREATE INDEX IF NOT EXISTS post_lang_title_idx ON post_lang (title);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS post_lang;
DROP INDEX IF EXISTS post_lang_title_idx;
-- +goose StatementEnd