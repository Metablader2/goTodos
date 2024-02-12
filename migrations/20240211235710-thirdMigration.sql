
-- +migrate Up
ALTER TABLE USERS
ADD CONSTRAINT email_unique UNIQUE (email);
-- +migrate Down
