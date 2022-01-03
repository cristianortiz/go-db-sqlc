-- name: GetAllProducts :many
 SELECT * FROM products
 ORDER BY title ASC;

 -- name: CreateNewProduct :execresult
INSERT INTO products (
  title,description,image,price
) VALUES (
  ?, ?, ?, ?
);