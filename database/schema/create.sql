DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nickname    TEXT UNIQUE,
	fistname    TEXT,
	lastname    TEXT,
	password    TEXT
	-- created_time string
);

DROP TABLE IF EXISTS questions;
CREATE TABLE questions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT,
    text TEXT,
    user_id INTEGER,
    FOREIGN KEY(user_id) REFERENCES users(id)
);

-- INSERT DATA TO USERS
INSERT INTO users (nickname, fistname, lastname, password) VALUES 
('AliceNick','Alice','Alicevich','alicesuper'),
('BobNick','Bob','Bobovich','bobsuper'),
('ClaraNick','Clara','Clarovich','clarasuper');

-- INSERT DATA QUESTIONS
INSERT INTO questions (title, text, user_id) VALUES 
('How change background in HTML?', 'I have html with this dat...', 3);


-- SHOW ALL USERS
SELECT * FROM users;
SELECT * FROM questions;