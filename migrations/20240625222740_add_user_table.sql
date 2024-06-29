-- +goose Up
CREATE TABLE users (
  id SERIAL PRIMARY KEY NOT NULL,
  username VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(128) NOT NULL,
  display_name VARCHAR(255),
  user_role VARCHAR(64) DEFAULT 'user' NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL
);

-- seed users data
INSERT INTO users (username, password, display_name, user_role) 
VALUES 
('user1', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'PLAN1', 'user'), 
('user2', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'PLAN2', 'user'), 
('user3', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'PLAN3', 'user'), 
('user4', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'PLAN4', 'user'), 
('user5', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'PLAN5', 'user'),
('user6', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'PLAN6', 'user'),
('user7', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'PLAN7', 'user'),
('user8', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'PLAN8', 'user'),
('user9', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'PLAN9', 'user'),
('user10', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'PLAN10', 'user'),
('user11', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'PLAN11', 'user'),
('user12', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'PLAN12', 'user'),
('user14', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'PLAN14', 'user'),
('user15', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'PLAN15', 'user'),
('adminir', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'ADMIN IR', 'admin'),
('viewer1', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'viewer', 'viewer')
;


-- +goose Down
DROP TABLE users;