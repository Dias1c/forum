-- CREATE TABLES
-- USERS
DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    nickname NVARCHAR(32) unique NOT NULL CHECK(LENGTH(nickname) <= 32),
    email NVARCHAR(320) unique NOT NULL CHECK(LENGTH(email) <= 320),
    password TEXT
);

-- USER SESSIONS 
DROP TABLE IF EXISTS sessions;
CREATE TABLE IF NOT EXISTS sessions (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    uuid TEXT NOT NULL,
    expired_at TEXT,
    user_id INTEGER,
    FOREIGN KEY(user_id) REFERENCES users(id)
);

-- POSTS CATEGORIES
DROP TABLE IF EXISTS categories;
CREATE TABLE categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name NVARCHAR(32) NOT NULL CHECK(LENGTH(name) <= 32)
);

-- POSTS
DROP TABLE IF EXISTS posts;
CREATE TABLE posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    title NVARCHAR(100) NOT NULL CHECK(LENGTH(title) <= 100),
    content TEXT NOT NULL,
    user_id INTEGER NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id)
);

-- POST CATEGORIES
DROP TABLE IF EXISTS categories_posts;
CREATE TABLE categories_posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    category_id INTEGER NOT NULL,
    post_id INTEGER NOT NULL,
    FOREIGN KEY(category_id) REFERENCES categories(id),
    FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE
);

-- POSTS LIKES
DROP TABLE IF EXISTS posts_votes;
CREATE TABLE posts_votes (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    vote INTEGER NOT NULL CHECK(vote IN(-1, 0, 1)),
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
    vote INTEGER NOT NULL CHECK(vote IN(-1, 0, 1)),
    user_id INTEGER NOT NULL,
    comment_id INTEGER NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(comment_id) REFERENCES comments(id)
);
