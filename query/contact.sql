-- name: CreateContact :one
INSERT INTO contact (
   owner,
   github,
   twitter
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: UpdateContact :one
UPDATE contact
SET owner = $2, github = $3, twitter = $4
WHERE id = $1
RETURNING *;

-- name: DeleteContact :exec
DELETE FROM contact
WHERE id = $1;

-- name: GetContact :one
SELECT * FROM contact
WHERE id = $1 LIMIT 1;

-- name: GetAllContact :many
SELECT * FROM contact
ORDER BY id;