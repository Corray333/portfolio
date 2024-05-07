-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS post_tag (
    post_id BIGINT NOT NULL,
    tag VARCHAR(25) NOT NULL,
    FOREIGN KEY (post_id) REFERENCES posts(post_id),
    PRIMARY KEY (post_id, tag)
);
CREATE INDEX IF NOT EXISTS post_tag_post_id_idx ON post_tag (post_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS post_tag;
DROP INDEX IF EXISTS post_tag_post_id_idx;
-- +goose StatementEnd