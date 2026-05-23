-- name: InsertUserLogin :exec
INSERT INTO user_logins (user_id, logged_in_at)
VALUES (@user_id, @logged_in_at);

-- name: GetDailyUniqueUsers :one
SELECT COUNT(DISTINCT user_id)
FROM user_logins
WHERE logged_in_at >= sqlc.arg(day)
  AND logged_in_at <  sqlc.arg(day) + INTERVAL '1 day';
