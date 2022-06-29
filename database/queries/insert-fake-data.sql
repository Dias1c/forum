-- INSERT VALUES TO TABLES

-- USERS
INSERT INTO users (nickname, email, password) VALUES
('Dias1c', 'admin@example.com', '1234567890'),
('alice', 'alice@example.com', '1234567890'),
('bob', 'bob@example.com', '1234567890');

-- CATEGORIES
INSERT INTO categories (name) VALUES
('about self'),
('sport'),
('music'),
('car'),
('health'),
('meme');

-- POSTS
INSERT INTO posts (title, content, user_id) VALUES
('About me', 'Hello there! I am Dias1c and I started to learn english. Here I will read, communicate, write posts and comments only in english!', 1);

-- COMMENTS
INSERT INTO comments (content, post_id, user_id) VALUES
('If you read this comment, lets learn english together', 1, 1);


-- SHOW TABLES CONTENTS
SELECT * FROM users;
SELECT * FROM categories;
SELECT * FROM posts;
SELECT * FROM categories_posts;
SELECT * FROM posts_votes;
SELECT * FROM comments;
SELECT * FROM comments_votes;
SELECT * FROM sessions;
