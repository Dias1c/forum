-- CREATE TABLES
-- USERS
DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    nickname NVARCHAR(20) NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL
);

-- POSTS CATEGORIES
DROP TABLE IF EXISTS categories;
CREATE TABLE categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name NVARCHAR(30) NOT NULL
);

-- POSTS
DROP TABLE IF EXISTS posts;
CREATE TABLE posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    title NVARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    category_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    FOREIGN KEY(category_id) REFERENCES categories(id),
    FOREIGN KEY(user_id) REFERENCES users(id)
);

-- POSTS LIKES
DROP TABLE IF EXISTS posts_votes;
CREATE TABLE posts_votes (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    vote INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    post_id INTEGER NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(post_id) REFERENCES posts(id)
);

-- POSTS COMMENTS
DROP TABLE IF EXISTS comments;
CREATE TABLE comments (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    content TEXT NOT NULL,
    user_id INTEGER NOT NULL,
    post_id INTEGER NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(post_id) REFERENCES posts(id)
);

-- COMMENTS LIKES
DROP TABLE IF EXISTS comments_votes;
CREATE TABLE comments_votes (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    vote INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    comment_id INTEGER NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(comment_id) REFERENCES comments(id)
);


-- INSERT VALUES TO TABLES

-- USERS
INSERT INTO users (nickname, email, password) VALUES
('Dias1c', 'example@mail.com', '1234567890');

SELECT * FROM users;

-- CATEGORIES
INSERT INTO categories (name) VALUES
('about self'),
('sport'),
('music'),
('car'),
('health'),
('meme');

SELECT * FROM categories;

-- POSTS
INSERT INTO posts (title, content, category_id, user_id) VALUES
('About me', 'Hello there! I am Dias1c and I started to learn english. Here I will read, communicate, write posts and comments only in english!', 1, 1);

SELECT * FROM posts;

-- COMMENTS
INSERT INTO comments (content, post_id, user_id) VALUES
('If you read this comment, lets learn english together', 1, 1);

SELECT * FROM comments;


-- SHOW TABLES CONTENTS
SELECT * FROM users;
SELECT * FROM categories;
SELECT * FROM posts;
SELECT * FROM posts_votes;
SELECT * FROM comments;
SELECT * FROM comments_votes;
