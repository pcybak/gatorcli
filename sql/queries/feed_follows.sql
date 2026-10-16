-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)

SELECT iff.*, u.name AS userName, f.name AS feedName
FROM inserted_feed_follow iff
JOIN users u on iff.user_id = u.id
JOIN feeds f on iff.feed_id = f.id;

-- name: GetFeedFollowsForUser :many
SELECT f.name, f.url
FROM feed_follows ff
JOIN feeds f on ff.feed_id = f.id
JOIN users u on ff.user_id = u.id
WHERE u.name = $1;
