-- CREATE TABLES
-- USERS
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    nickname NVARCHAR(32) UNIQUE NOT NULL CHECK(LENGTH(nickname) <= 32),
    email NVARCHAR(320) UNIQUE NOT NULL CHECK(LENGTH(email) <= 320),
    password TEXT,
    created_at TEXT NOT NULL
);

-- USER SESSIONS 
CREATE TABLE IF NOT EXISTS sessions (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    uuid TEXT NOT NULL,
    expired_at TEXT,
    user_id INTEGER NOT NULL UNIQUE,
    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- POSTS CATEGORIES
CREATE TABLE IF NOT EXISTS categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name NVARCHAR(32) NOT NULL CHECK(LENGTH(name) <= 32) UNIQUE,
    created_at TEXT NOT NULL
);

-- POSTS
CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    title NVARCHAR(100) NOT NULL CHECK(LENGTH(title) <= 100),
    content TEXT NOT NULL,
    user_id INTEGER NOT NULL,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id)
);

-- POST CATEGORIES
CREATE TABLE IF NOT EXISTS posts_categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    category_id INTEGER NOT NULL,
    post_id INTEGER NOT NULL,
    FOREIGN KEY(category_id) REFERENCES categories(id) ON DELETE CASCADE,
    FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE
);

-- POSTS LIKES
CREATE TABLE IF NOT EXISTS posts_votes (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    vote INTEGER NOT NULL CHECK(vote IN(-1, 0, 1)),
    user_id INTEGER NOT NULL,
    post_id INTEGER NOT NULL,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    UNIQUE (user_id, post_id),
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE
);

-- POSTS COMMENTS
CREATE TABLE IF NOT EXISTS posts_comments (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    content TEXT NOT NULL,
    user_id INTEGER NOT NULL,
    post_id INTEGER NOT NULL,
    created_at TEXT NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE
);

-- POSTS COMMENTS LIKES
CREATE TABLE IF NOT EXISTS posts_comments_votes (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    vote INTEGER NOT NULL CHECK(vote IN(-1, 0, 1)),
    user_id INTEGER NOT NULL,
    comment_id INTEGER NOT NULL,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    UNIQUE (user_id, comment_id),
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(comment_id) REFERENCES posts_comments(id) ON DELETE CASCADE
);
