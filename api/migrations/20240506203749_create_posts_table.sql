-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS posts (
    post_id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY (INCREMENT 1 START 1 CACHE 1),
    cover TEXT NOT NULL,
    created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM CURRENT_TIMESTAMP),
    updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM CURRENT_TIMESTAMP),
    type INTEGER NOT NULL DEFAULT 1,
    CONSTRAINT posts_pk PRIMARY KEY (post_id)
);
CREATE INDEX IF NOT EXISTS posts_created_at_idx ON posts (created_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS posts;
DROP INDEX IF EXISTS posts_created_at_idx;
-- +goose StatementEnd
