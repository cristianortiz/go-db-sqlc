-- name: GetAllOrderItems :many
 SELECT * FROM order_items;
-- name: GetItemsByOrderID :many
 SELECT * FROM order_items WHERE order_id=?;