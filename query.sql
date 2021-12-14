 -- name: GetUserByID :one
 SELECT * FROM users
 WHERE id = ? LIMIT 1;

 -- name: GetAllUsers :many
 SELECT * FROM users
 ORDER BY name;

 -- name: CreateUser :execresult
INSERT INTO users (
  firstname,lastname,email,upassword,isambassador
) VALUES (
  ?, ?, ?, ?, ?
);

-- name: DeleteUserByID :exec
DELETE FROM users
WHERE id = ?;