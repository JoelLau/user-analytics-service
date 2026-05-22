-- name: GetUserLogins :many
SELECT id
  , user_id
  , login_time
FROM user_logins;
