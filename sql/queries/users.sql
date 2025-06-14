-- name: CreateUser :one
-- This query creates a new user in the database.
INSERT INTO users (id, username, email, created_at, updated_at, api_key)
VALUES ($1, $2, $3, NOW(), NOW(),
    encode(sha256(random()::text::bytea), 'hex'))
RETURNING *;

-- name: GetUserByID :one
-- This query retrieves a user by their ID.
SELECT * FROM users WHERE id = $1;