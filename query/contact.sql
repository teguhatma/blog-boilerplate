-- name: CreateContact :one
INSERT INTO contact (
   owner,
   email,
   github,
   twitter
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: ListContact :many
SELECT * FROM contact
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: UpdateContact :one
UPDATE contact
SET owner = $2, email = $3, github = $4, twitter = $5
WHERE id = $1
RETURNING *;

-- name: DeleteContact :exec
DELETE FROM contact
WHERE id = $1;