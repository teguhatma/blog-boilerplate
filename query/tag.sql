-- name: CreateTag :one
INSERT INTO tag (
   name 
) VALUES (
    $1
) RETURNING *;

-- name: ListTags :many
SELECT * FROM tag
ORDER BY name;

-- name: UpdateTag :one
UPDATE tag
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteTag :exec
DELETE FROM tag
WHERE id = $1;

-- name: GetTag :one
SELECT * FROM tag
WHERE id = $1 LIMIT 1;