
-- +migrate Up
CREATE TABLE TODO(
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    duedate TIMESTAMP
);
-- +migrate Down
