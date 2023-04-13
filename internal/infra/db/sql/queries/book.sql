-- name: CreateBook :exec
INSERT INTO books (id, title, isbn, year, pages, created_at) VALUES ($1, $2, $3, $4, $5, $6);

-- name: ConnectBookAuthor :exec
INSERT INTO books_authors (book_id, author_id) VALUES ($1, $2);

-- name: ConnectBookSubject :exec
INSERT INTO books_subjects (book_id, subject_id) VALUES ($1, $2);
