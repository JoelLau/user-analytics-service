-- +goose Up
CREATE TABLE user_logins (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    login_time TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    UNIQUE(user_id, login_time)
);
