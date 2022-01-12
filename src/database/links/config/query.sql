-- name: GetAllLinks :many
 SELECT * FROM links
 LEFT JOIN users ON links.user_id=users.id;