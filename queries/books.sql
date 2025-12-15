-- name: FindAllBooks :many
SELECT * FROM books;

-- name: InsertBook :one
INSERT INTO books (id, name, author, price)
VALUES (uuid_generate_v4(), $1, $2, $3)
RETURNING *;
