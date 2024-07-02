-- +goose Up
CREATE TABLE plan (
  id SERIAL PRIMARY KEY NOT NULL,
  name VARCHAR(255) UNIQUE NOT NULL,
  user_id INT NOT NULL REFERENCES users (id),
  topic VARCHAR(255) NOT NULL,
  topic_en VARCHAR(255) NOT NULL,
  goal VARCHAR(512) NOT NULL,
  readiness_willingness TEXT,
  readiness_willingness_updated_at TIMESTAMP WITH TIME ZONE,
  readiness_willingness_updated_by INT REFERENCES users (id),
  ir_goal_type VARCHAR(64),
  ir_goal_type_updated_at TIMESTAMP WITH TIME ZONE,
  ir_goal_type_updated_by INT REFERENCES users (id),
  ir_goal_details TEXT,
  ir_goal_details_updated_at TIMESTAMP WITH TIME ZONE,
  ir_goal_details_updated_by INT REFERENCES users (id),
  proposed_activity TEXT,
  proposed_activity_updated_at TIMESTAMP WITH TIME ZONE,
  proposed_activity_updated_by INT REFERENCES users (id),
  plan_note TEXT,
  plan_note_updated_at TIMESTAMP WITH TIME ZONE,
  plan_note_updated_by INT REFERENCES users (id),
  contact_person TEXT,
  contact_person_updated_at TIMESTAMP WITH TIME ZONE,
  contact_person_updated_by INT REFERENCES users (id),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
  updated_by INT NOT NULL REFERENCES users (id)
);

INSERT 
INTO plan (
  name, user_id, topic, topic_en, goal, readiness_willingness, readiness_willingness_updated_at, readiness_willingness_updated_by,
  ir_goal_type,ir_goal_type_updated_at, ir_goal_type_updated_by, ir_goal_details, ir_goal_details_updated_at, ir_goal_details_updated_by,
  proposed_activity, proposed_activity_updated_at, proposed_activity_updated_by, plan_note, plan_note_updated_at, plan_note_updated_by,
  contact_person, contact_person_updated_at, contact_person_updated_by, updated_by
) 
VALUES 
('PLAN1', 1,'แผนควบคุมยาสูบ', 'Tobacco Control', 'Exchange knowledge with international audiences and participating in the discussion in the global and regional networks', 
'ready 1', '2024-06-30 10:46:35.093141+00', 1, 'type 1', '2024-06-30 10:46:35.093141+00', 1, 'goal details 1', '2024-06-30 10:46:35.093141+00', 1, 'activity 1', '2024-06-30 10:46:35.093141+00', 1, 
'plan note 1', '2024-06-30 10:46:35.093141+00', 1, 'contact 1', '2024-06-30 10:46:35.093141+00', 1, 1), 
('PLAN2', 2,'แผนควบคุมเครื่องดื่มแอลกอฮอล์และสิ่งเสพติด', 'Alcohol and Substance Abuse Control', 'Goal 2',
'ready 2', '2024-06-30 10:46:35.093141+00', 2, 'type 2', '2024-06-30 10:46:35.093141+00', 2, 'goal details 2', '2024-06-30 10:46:35.093141+00', 2, 'activity 2','2024-06-30 10:46:35.093141+00', 2,
'plan note 2', '2024-06-30 10:46:35.093141+00', 2, 'contact 2', '2024-06-30 10:46:35.093141+00', 2, 2),
('PLAN3', 3, 'แผนการจัดการความปลอดภัย และปัจจัยเสี่ยงทางสังคม', 'Road Safety and Social Risk Management Plan', 'Goal 3',
'ready 3', '2024-06-30 10:46:35.093141+00', 3, 'type 3', '2024-06-30 10:46:35.093141+00', 3, 'goal details 3', '2024-06-30 10:46:35.093141+00', 3, 'activity 3','2024-06-30 10:46:35.093141+00', 3,
'plan note 3', '2024-06-30 10:46:35.093141+00', 3, 'contact 3', '2024-06-30 10:46:35.093141+00', 3, 3),
('PLAN4', 4, 'แผนควบคุมปัจจัยเสี่ยงทางสุขภาพ', 'Health Risk Control Plan', 'Goal 4',
'ready 4', '2024-06-30 10:46:35.093141+00', 4, 'type 4', '2024-06-30 10:46:35.093141+00', 4, 'goal details 4', '2024-06-30 10:46:35.093141+00', 4, 'activity 4','2024-06-30 10:46:35.093141+00', 4,
'plan note 4', '2024-06-30 10:46:35.093141+00', 4, 'contact 4', '2024-06-30 10:46:35.093141+00', 4, 4),
('PLAN5', 5, 'แผนสุขภาวะประชากรกลุ่มเฉพาะ', 'Health Promotion Plan for Vulnerable Populations', 'Goal 5',
'ready 5', '2024-06-30 10:46:35.093141+00', 5, 'type 5', '2024-06-30 10:46:35.093141+00', 5, 'goal details 5', '2024-06-30 10:46:35.093141+00', 5, 'activity 5','2024-06-30 10:46:35.093141+00', 5,
'plan note 5', '2024-06-30 10:46:35.093141+00', 5, 'contact 5', '2024-06-30 10:46:35.093141+00', 5, 5),
('PLAN6', 6, 'แผนสุขภาวะชุมชน', 'Healthy Community Strengthening Plan', 'Goal 6',
'ready 6', '2024-06-30 10:46:35.093141+00', 6, 'type 6', '2024-06-30 10:46:35.093141+00', 6, 'goal details 6', '2024-06-30 10:46:35.093141+00', 6, 'activity 6','2024-06-30 10:46:35.093141+00', 6,
'plan note 6', '2024-06-30 10:46:35.093141+00', 6, 'contact 6', '2024-06-30 10:46:35.093141+00', 6, 6),
('PLAN7', 7, 'แผนสุขภาวะเด็ก เยาวชน และครอบครัว', 'Healthy Child, Youth, and Family Promotion Plan', 'Goal 7',
'ready 7', '2024-06-30 10:46:35.093141+00', 7, 'type 7', '2024-06-30 10:46:35.093141+00', 7, 'goal details 7', '2024-06-30 10:46:35.093141+00', 7, 'activity 7','2024-06-30 10:46:35.093141+00', 7,
'plan note 7', '2024-06-30 10:46:35.093141+00', 7, 'contact 7', '2024-06-30 10:46:35.093141+00', 7, 7),
('PLAN8', 8, 'แผนสร้างเสริมสุขภาวะในองค์กร', 'Health Promotion in Organizations Plan', 'Goal 8',
'ready 8', '2024-06-30 10:46:35.093141+00', 8, 'type 8', '2024-06-30 10:46:35.093141+00', 8, 'goal details 8', '2024-06-30 10:46:35.093141+00', 8, 'activity 8','2024-06-30 10:46:35.093141+00', 8,
'plan note 8', '2024-06-30 10:46:35.093141+00', 8, 'contact 8', '2024-06-30 10:46:35.093141+00', 8, 8),
('PLAN9', 9, 'แผนส่งเสริมกิจกรรมทางกาย', 'Physical Activity Promotion Plan', 'Goal 9',
'ready 9', '2024-06-30 10:46:35.093141+00', 9, 'type 9', '2024-06-30 10:46:35.093141+00', 9, 'goal details 9', '2024-06-30 10:46:35.093141+00', 9, 'activity 9','2024-06-30 10:46:35.093141+00', 9,
'plan note 9', '2024-06-30 10:46:35.093141+00', 9, 'contact 9', '2024-06-30 10:46:35.093141+00', 9, 9),
('PLAN10', 10, 'แผนระบบสื่อและวิถีสุขภาวะทางปัญญา', 'Health Media System and Spiritual Health Pathway Plan', 'Goal 10',
'ready 10', '2024-06-30 10:46:35.093141+00', 10, 'type 10', '2024-06-30 10:46:35.093141+00', 10, 'goal details 10', '2024-06-30 10:46:35.093141+00', 10, 'activity 10','2024-06-30 10:46:35.093141+00', 10,
'plan note 10', '2024-06-30 10:46:35.093141+00', 10, 'contact 10', '2024-06-30 10:46:35.093141+00', 10, 10),
('PLAN11', 11, 'แผนสร้างสรรค์โอกาสและนวัตกรรมสุขภาวะ', 'Health Promotion Innovation and Open Grant Plan', 'Goal 11',
'ready 11', '2024-06-30 10:46:35.093141+00', 11, 'type 11', '2024-06-30 10:46:35.093141+00', 11, 'goal details 11', '2024-06-30 10:46:35.093141+00', 11, 'activity 11','2024-06-30 10:46:35.093141+00', 11,
'plan note 11', '2024-06-30 10:46:35.093141+00', 11, 'contact 11', '2024-06-30 10:46:35.093141+00', 11, 11),
('PLAN12', 12, 'แผนสนับสนุนการสร้างเสริมสุขภาพ ผ่านระบบบริการสุขภาพ', 'Health Promotion in Health Service System Plan', 'Goal 12',
'ready 12', '2024-06-30 10:46:35.093141+00', 12, 'type 12', '2024-06-30 10:46:35.093141+00', 12, 'goal details 12', '2024-06-30 10:46:35.093141+00', 12, 'activity 12','2024-06-30 10:46:35.093141+00', 12,
'plan note 12', '2024-06-30 10:46:35.093141+00', 12, 'contact 12', '2024-06-30 10:46:35.093141+00', 12, 12),
('PLAN14', 13, 'แผนอาหารเพื่อสุขภาวะ', 'Healthy Food Promotion Plan', 'Goal 14',
'ready 14', '2024-06-30 10:46:35.093141+00', 13, 'type 14', '2024-06-30 10:46:35.093141+00', 13, 'goal details 14', '2024-06-30 10:46:35.093141+00', 13, 'activity 14', '2024-06-30 10:46:35.093141+00', 13,
'plan note 14', '2024-06-30 10:46:35.093141+00', 13, 'contact 14', '2024-06-30 10:46:35.093141+00', 13, 13),
('PLAN15', 14, 'แผนสร้างเสริมความเข้าใจสุขภาวะ', 'Health Literacy Promotion Plan', 'Goal 15',
'ready 15', '2024-06-30 10:46:35.093141+00', 14, 'type 15', '2024-06-30 10:46:35.093141+00', 14, 'goal details 15', '2024-06-30 10:46:35.093141+00', 14, 'activity 15', '2024-06-30 10:46:35.093141+00', 14,
'plan note 15', '2024-06-30 10:46:35.093141+00', 14, 'contact 15', '2024-06-30 10:46:35.093141+00', 14, 14)
;

INSERT INTO plan (name, user_id,  topic, topic_en, goal, updated_by) VALUES ('ADMIN', 15, 'ไออาร์', 'IR', 'Goal Admin', 15);


-- +goose Down
ALTER TABLE plan DROP COLUMN user_id;
DROP TABLE plan;