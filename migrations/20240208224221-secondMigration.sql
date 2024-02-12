
-- +migrate Up
CREATE TABLE USERS(
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL,
    password TEXT NOT NULL
);
-- +migrate Down
