-- name: GetAllProducts :many
 SELECT * FROM products
 ORDER BY title ASC;

 -- name: GetProductByID :one
 SELECT id,category,title,description,price,image FROM products
 WHERE id=?;

 -- name: CreateNewProduct :execresult
INSERT INTO products (
  category,title,description,price,image
) VALUES (
  ?, ?, ?, ?,?
);

-- name: UpdateProduct :execresult
UPDATE products SET
  category = ?,title = ?,description= ?,price = ?,image = ?
WHERE id = ?;

-- name: DeleteProduct :execresult
DELETE FROM products
WHERE id = ?;