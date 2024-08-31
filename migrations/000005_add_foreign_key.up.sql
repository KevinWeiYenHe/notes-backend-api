ALTER TABLE notes
ADD CONSTRAINT fk_author
FOREIGN KEY (author_id) REFERENCES users(id);