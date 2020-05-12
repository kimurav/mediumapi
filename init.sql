CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(50),
    username VARCHAR(50),
    bio VARCHAR(500),
    password VARCHAR(50)
);

CREATE TABLE article(
    id SERIAL PRIMARY KEY,
    slug VARCHAR(100),
    user_id INT REFERENCES users(id),
    title VARCHAR(50),
    body TEXT
);



INSERT INTO users(email, username, bio, password)
VALUES('manav@uoguelph.ca', 'manavp1996', 'hey its me', '131231j3');
