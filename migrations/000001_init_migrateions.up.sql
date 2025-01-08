CREATE TABLE gender (
    id SERIAL PRIMARY KEY,
    gender VARCHAR(50) NOT NULL
);

CREATE TABLE category (
    id SERIAL PRIMARY KEY,
    category VARCHAR(100) NOT NULL
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    encrypted_password TEXT,
    image TEXT,
    name VARCHAR(100),
    lastname VARCHAR(100),
    date_of_birth DATE,
    gender_id INT REFERENCES gender(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE news (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    category_id INT REFERENCES category(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    content TEXT NOT NULL,
    image TEXT
);

CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    news_id INT REFERENCES news(id) ON DELETE CASCADE,
    parent_id INT REFERENCES comments(id) ON DELETE CASCADE,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    date_post TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
