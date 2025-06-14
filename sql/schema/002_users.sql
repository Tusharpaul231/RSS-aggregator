-- +goose Up
ALTER TABLE users
ADD COLUMN IF NOT EXISTS api_key TEXT UNIQUE NOT NULL DEFAULT (
    encode(digest(gen_random_uuid()::text, 'sha256'), 'hex')
);

-- +goose Down
ALTER TABLE users DROP COLUMN IF EXISTS api_key;
