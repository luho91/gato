-- +goose Up
CREATE TABLE feed_follows(id UUID PRIMARY KEY NOT NULL, created_at TIMESTAMP NOT NULL, updated_at TIMESTAMP NOT NULL, user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE, feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE);

CREATE UNIQUE INDEX users_id_feeds_id_unique ON feed_follows(user_id, feed_id);

-- +goose Down
DROP INDEX IF EXISTS users_id_feeds_id_unique;
DROP TABLE IF EXISTS feed_follows;
