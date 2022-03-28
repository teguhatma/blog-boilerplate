-- name: CreateEntries :one
INSERT INTO entries (
    owner,
    tag_name,
    blog,
    title,
    read_time
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;


-- name: GetEntries :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;


-- name: ListEntries :many
SELECT * FROM entries
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: UpdateEntries :one
UPDATE entries
SET owner = $2, blog = $3, title = $4, read_time = $5
WHERE id = $1
RETURNING *;


-- name: DeleteEntries :exec
DELETE FROM entries
WHERE id = $1;
