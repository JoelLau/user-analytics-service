-- +goose Up

CREATE TABLE user_logins (
    user_id      BIGINT      NOT NULL,
    logged_in_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

COMMENT ON TABLE  user_logins              IS 'append-only event log of user login events.';
COMMENT ON COLUMN user_logins.user_id      IS 'user identifier - each unique value represents one user.';
COMMENT ON COLUMN user_logins.logged_in_at IS 'provided by the database at insert time; avoid supplying value on insertion';

CREATE INDEX ON user_logins (user_id, logged_in_at);
