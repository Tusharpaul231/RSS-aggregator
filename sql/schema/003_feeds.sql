-- +goose Up
CREATE TABLE IF NOT EXISTS feeds (
    id UUID PRIMARY KEY,
    username TEXT NOT NULL,
    email TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    url TEXT UNIQUE NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE

);

-- +goose Down
DROP TABLE IF EXISTS feeds;
