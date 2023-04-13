-- name: CreateQuote :exec
INSERT INTO quotes (id, book_id, content, page, created_at) VALUES ($1, $2, $3, $4, $5);

-- name: ListQuotes :many
SELECT q.*, b.title as book_title, count(*) OVER() AS full_count
FROM  quotes q JOIN books b ON q.book_id = b.id
OFFSET $1 LIMIT $2;