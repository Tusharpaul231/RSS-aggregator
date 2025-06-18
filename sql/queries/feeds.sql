-- name: CreateFeed :one
-- This query creates a new user in the database.
INSERT INTO feeds (id, username, email, created_at, updated_at, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    NOW(),
    NOW(),
    $4,
    $5
)
RETURNING *;
