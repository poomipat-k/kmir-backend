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
('p1thhealth1', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'PLAN1', 'user'), 
('p2thhealth1', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'PLAN2', 'user'), 
('p3thhealth1', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'PLAN3', 'user'), 
('p4thhealth1', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'PLAN4', 'user'), 
('p5thhealth1', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'PLAN5', 'user'),
('p6thhealth1', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'PLAN6', 'user'),
('p7thhealth1', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'PLAN7', 'user'),
('p8thhealth1', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'PLAN8', 'user'),
('p9thhealth1', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'PLAN9', 'user'),
('p10thhealth1', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'PLAN10', 'user'),
('p11thhealth1', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'PLAN11', 'user'),
('p12thhealth1', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'PLAN12', 'user'),
('p14thhealth1', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'PLAN14', 'user'),
('p15thhealth1', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'PLAN15', 'user'),
('adminthhealth1', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'ADMIN', 'admin'),
('viewerthhealth1', '$2a$10$gI6Po0SLyJDazdS/nYx5FeScn4quG/uKKOTf0eFlMYTQPUeVOl5pK_eitPzQbO', 'viewer', 'viewer')
;


-- +goose Down
DROP TABLE users;