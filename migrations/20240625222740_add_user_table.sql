-- +goose Up
CREATE TABLE users (
  id SERIAL PRIMARY KEY NOT NULL,
  username VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(128) NOT NULL,
  user_role VARCHAR(64) DEFAULT 'applicant' NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL
);
-- +goose Down
DROP TABLE users;