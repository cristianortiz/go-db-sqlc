-- name: GetAllOrders :many
 SELECT * FROM orders;


-- name: GetOrdersByUsers :many
SELECT CONCAT(users.firstname," ",users.lastname) as name, orders.id,user_id,complete
FROM orders
LEFT JOIN users ON orders.user_id=users.id;