CREATE TABLE links (
  id   BIGINT  NOT NULL AUTO_INCREMENT PRIMARY KEY,
  code TEXT NOT NULL,
  user_id BIGINT NOT NULL,
   CONSTRAINT fk_links_users FOREIGN KEY (user_id) REFERENCES users(id)
 );
 
 CREATE TABLE users (
  id   BIGINT  NOT NULL AUTO_INCREMENT PRIMARY KEY,
  firstname text NOT Null,
  lastname text NOT NULL,
  email text NOT NULL UNIQUE,
  upassword text NOT NULL,
  isambassador TINYINT
);
