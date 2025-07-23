CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username TEXT,
    password TEXT
);

CREATE TABLE messages (
    sender TEXT,
    receiver TEXT,
    text TEXT
);