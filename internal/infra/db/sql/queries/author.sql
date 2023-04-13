-- name: CreateAuthor :exec
INSERT INTO authors (id, name, bio, created_at) VALUES ($1, $2, $3, $4);