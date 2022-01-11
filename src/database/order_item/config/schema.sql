CREATE TABLE order_items (
  id bigint  NOT NULL AUTO_INCREMENT PRIMARY KEY,
  order_id bigint  NOT NULL,
  product_title TEXT NOT NULL,
  price double NOT NULL,
  quantity bigint NOT NULL,
  admin_revenue DOUBLE NOT NULL,
  ambassador_revenue DOUBLE NOT NULL);
  ALTER TABLE order_items ADD CONSTRAINT fk_orderItems_order FOREIGN KEY (order_id) REFERENCES orders(id);