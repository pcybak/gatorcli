-- name: AddFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;


-- name: GetAllFeeds :many
SELECT f.name, f.url, u.name
FROM feeds f
JOIN users u ON f.user_id = u.id;


-- name: GetFeedByURL :one
SELECT id, created_at, updated_at, name, url, user_id
FROM feeds
WHERE url = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = $2,
updated_at = $2
WHERE id = $1;


-- name: GetNextFeedToFetch :one
SELECT f.id, f.url
FROM feeds f
ORDER BY f.last_fetched_at ASC NULLS FIRST
LIMIT 1;
