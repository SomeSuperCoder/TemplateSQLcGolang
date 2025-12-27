-- name: FindAllBooks :many
SELECT * FROM books ORDER BY created_at DESC;

-- name: InsertBook :one
INSERT INTO books (
  name,
  author,
  price
)
VALUES (
  $1,
  $2,
  $3
)
RETURNING *;

-- name: UpdateBook :one
UPDATE books
SET
  name = COALESCE($2, name),
  author = COALESCE($3, author),
  price = COALESCE($4, price)
WHERE id = $1
RETURNING *;
