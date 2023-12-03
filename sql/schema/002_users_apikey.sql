-- +goose Up
ALTER TABLE users
ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT (
    enocde(sha256(random()::text::bytea), 'hex')
  );
-- +goose Down
DROP TABLE users;