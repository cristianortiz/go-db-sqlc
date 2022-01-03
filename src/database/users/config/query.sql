 -- name: GetUserByID :one
 SELECT * FROM users
 WHERE id = ? LIMIT 1;
 -- name: GetUserParamsByID :one
 SELECT id,firstname,lastname,email FROM users
 WHERE id = ? LIMIT 1;

 -- name: GetAmbassadors :many
 SELECT id,firstname,lastname,email,isambassador FROM users
 WHERE isambassador = 1 ;
 
 -- name: GetUserByEmail :one
 SELECT * FROM users
 WHERE email = ? LIMIT 1;

 -- name: UserEmailExists :one
 SELECT COUNT(*) FROM users
 WHERE  email = ? LIMIT 1;

 -- name: GetAllUsers :many
 SELECT * FROM users
 ORDER BY name;

 -- name: CreateUser :execresult
INSERT INTO users (
  firstname,lastname,email,upassword,isambassador
) VALUES (
  ?, ?, ?, ?, ?
);
-- name: UpdateUserInfo :execresult
UPDATE users SET
  firstname = ?,lastname= ?,email= ?
WHERE id = ?
 
;
-- name: UpdateUserPassword :execresult
UPDATE users SET
  upassword = ?
WHERE id = ?
 
;
-- name: DeleteUserByID :exec
DELETE FROM users
WHERE id = ?;