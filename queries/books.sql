-- name: FindAllBooks :many
SELECT * FROM books;

-- name: InsertBook :one
INSERT INTO books (
  id,
  name,
  author,
  price
)
VALUES (
  uuid_generate_v4(),
  $1,
  $2,
  $3
)
RETURNING *;

-- name: UpdateBook :one
update books
set
  name = COALESCE($2, name),
  author = COALESCE($3, author),
  price = COALESCE($4, price)
where id = $1
returning *;
