INSERT INTO users (name, nick, email, password) VALUES ('Jose', 'jose', 'jose@gmail.com', '$2a$10$gcMIBUclhmLjorTxHewcWO2DhJJRkE97VAf2JioGjWYUUpE6WNIdO');
INSERT INTO users (name, nick, email, password) VALUES ('Maria', 'maria', 'maria@gmail.com', '$2a$10$gcMIBUclhmLjorTxHewcWO2DhJJRkE97VAf2JioGjWYUUpE6WNIdO');
INSERT INTO users (name, nick, email, password) VALUES ('Joao', 'joao', 'joao@gmail.com', '$2a$10$gcMIBUclhmLjorTxHewcWO2DhJJRkE97VAf2JioGjWYUUpE6WNIdO');

INSERT INTO followers (user_id, follower_id) VALUES (1, 2);
INSERT INTO followers (user_id, follower_id) VALUES (3, 1);
INSERT INTO followers (user_id, follower_id) VALUES (1, 3);

INSERT INTO posts (title, content, author_id) VALUES ('First Post', 'Hello, world from user 1!', 1);
INSERT INTO posts (title, content, author_id) VALUES ('Second Post', 'This is a second post from user 2.', 2);
INSERT INTO posts (title, content, author_id) VALUES ('Third Post', 'This is a third post from user 3.', 3);
