CREATE TABLE links (
  id   BIGINT  NOT NULL AUTO_INCREMENT PRIMARY KEY,
  code TEXT NOT NULL,
  user_id BIGINT NOT NULL
 );
 ALTER TABLE links ADD CONSTRAINT fk_links_user FOREIGN KEY (user_id) REFERENCES users(id);