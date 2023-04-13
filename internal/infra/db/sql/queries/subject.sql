-- name: CreateSubject :exec
INSERT INTO subjects (id, name, description, created_at) VALUES ($1, $2, $3, $4);