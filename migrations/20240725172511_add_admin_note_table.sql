-- +goose Up
CREATE TABLE admin_note (
  id SERIAL PRIMARY KEY NOT NULL,
  note TEXT
);

INSERT INTO admin_note (note) VALUES ('<p>-</p>');

-- +goose Down
DROP TABLE admin_note;