CREATE TABLE orders (
  id bigint unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  transaction_id TEXT NOT NULL,
  user_id bigint unsigned NOT NULL,
  code TEXT NOT null,
  ambassador_email TEXT NOT null,
  firstname TEXT NOT null,
  lastname TEXT NOT null,
  email TEXT NOT null,
  address TEXT NOT null,
  city TEXT NOT null,
  country TEXT NOT null,
  zip TEXT NOT null,
  complete tinyint(1) NOT NULL DEFAULT 0);
  