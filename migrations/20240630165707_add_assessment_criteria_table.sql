-- +goose Up
CREATE TABLE assessment_criteria (
  id SERIAL PRIMARY KEY NOT NULL,
  version INT NOT NULL,
  category VARCHAR(64) NOT NULL,
  display VARCHAR(255) NOT NULL,
  order_number INT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL
);

INSERT INTO assessment_criteria (version, category, display, order_number)
VALUES 
(1, 'willingness', 'The plan''s international activities are well-aligned with the foundation''s overall strategic objectives.', 1),
(1, 'willingness', 'The plan is highly aware of the benefits and importance of engaging in international activities.', 2),
(1, 'willingness', 'The plan feels strongly supported by the International Relations Section and Executives in pursuing international activities.', 3),
(1, 'capacity', 'The plan uses resources efficiently for international activities.', 4),
(1, 'capacity', 'The plan is highly competent in handling international affairs.', 5),
(1, 'capacity', 'The quality of the plan''s past and existing international partnerships is excellent', 6),
(1, 'capacity', 'The plan effectively measures the outcomes of its international activities.', 7)
;

-- +goose Down
DROP TABLE assessment_criteria;