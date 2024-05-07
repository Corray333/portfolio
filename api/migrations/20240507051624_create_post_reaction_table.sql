-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS post_reaction (
    post_id BIGINT NOT NULL,
    type INTEGER NOT NULL,
    number INTEGER NOT NULL DEFAULT 0,
    FOREIGN KEY (post_id) REFERENCES posts(post_id),
    PRIMARY KEY (post_id, type)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS post_reaction;
-- +goose StatementEnd