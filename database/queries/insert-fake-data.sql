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

-- CATEGORY POST
INSERT INTO categories_posts (category_id, post_id) VALUES
(1,1);

-- COMMENTS
INSERT INTO comments (content, post_id, user_id) VALUES
('If you read this comment, lets learn english together', 1, 1);
