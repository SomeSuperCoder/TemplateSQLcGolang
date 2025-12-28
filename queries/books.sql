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
  name = COALESCE(sqlc.narg('name'), name),
  author = COALESCE(sqlc.narg('author'), author),
  price = COALESCE(sqlc.narg('price'), price)
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: Like :one
UPDATE books
SET
  likes = likes + 1
WHERE id = $1
    AND (@dummy::boolean IS NULL OR true)
RETURNING likes, 1 as _dummy;
