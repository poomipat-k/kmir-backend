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
  ir_goal_type VARCHAR(64),
  ir_goal_type_updated_at TIMESTAMP WITH TIME ZONE,
  ir_goal_details TEXT,
  ir_goal_details_updated_at TIMESTAMP WITH TIME ZONE,
  proposed_activity TEXT,
  proposed_activity_updated_at TIMESTAMP WITH TIME ZONE,
  plan_note TEXT,
  plan_note_updated_at TIMESTAMP WITH TIME ZONE,
  contact_person TEXT,
  contact_person_updated_at TIMESTAMP WITH TIME ZONE,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL
);

INSERT 
INTO plan (
  name, user_id, topic, topic_en, goal, readiness_willingness, readiness_willingness_updated_at, ir_goal_type,ir_goal_type_updated_at,
  ir_goal_details, ir_goal_details_updated_at, proposed_activity, proposed_activity_updated_at, plan_note, plan_note_updated_at,
  contact_person, contact_person_updated_at
) 
VALUES 
('PLAN1', 1,'แผนควบคุมยาสูบ', 'Tobacco Control', 'Exchange knowledge with international audiences and participating in the discussion in the global and regional networks', 
'ready 1', '2024-06-30 10:46:35.093141+00', 'type 1', '2024-06-30 10:46:35.093141+00', 'goal details 1', '2024-06-30 10:46:35.093141+00', 'activity 1', '2024-06-30 10:46:35.093141+00', 
'plan note 1', '2024-06-30 10:46:35.093141+00', 'contact 1', '2024-06-30 10:46:35.093141+00'), 
('PLAN2', 2,'แผนควบคุมเครื่องดื่มแอลกอฮอล์และสิ่งเสพติด', 'Alcohol and Substance Abuse Control', 'Goal 2',
'ready 2', '2024-06-30 10:46:35.093141+00', 'type 2', '2024-06-30 10:46:35.093141+00', 'goal details 2', '2024-06-30 10:46:35.093141+00', 'activity 2','2024-06-30 10:46:35.093141+00',
'plan note 2', '2024-06-30 10:46:35.093141+00', 'contact 2', '2024-06-30 10:46:35.093141+00'),
('PLAN3', 3, 'แผนการจัดการความปลอดภัย และปัจจัยเสี่ยงทางสังคม', 'Road Safety and Social Risk Management Plan', 'Goal 3',
'ready 3', '2024-06-30 10:46:35.093141+00', 'type 3', '2024-06-30 10:46:35.093141+00', 'goal details 3', '2024-06-30 10:46:35.093141+00', 'activity 3','2024-06-30 10:46:35.093141+00',
'plan note 3', '2024-06-30 10:46:35.093141+00', 'contact 3', '2024-06-30 10:46:35.093141+00'),
('PLAN4', 4, 'แผนควบคุมปัจจัยเสี่ยงทางสุขภาพ', 'Health Risk Control Plan', 'Goal 4',
'ready 4', '2024-06-30 10:46:35.093141+00', 'type 4', '2024-06-30 10:46:35.093141+00', 'goal details 4', '2024-06-30 10:46:35.093141+00', 'activity 4','2024-06-30 10:46:35.093141+00',
'plan note 4', '2024-06-30 10:46:35.093141+00', 'contact 4', '2024-06-30 10:46:35.093141+00'),
('PLAN5', 5, 'แผนสุขภาวะประชากรกลุ่มเฉพาะ', 'Health Promotion Plan for Vulnerable Populations', 'Goal 5',
'ready 5', '2024-06-30 10:46:35.093141+00', 'type 5', '2024-06-30 10:46:35.093141+00', 'goal details 5', '2024-06-30 10:46:35.093141+00', 'activity 5','2024-06-30 10:46:35.093141+00',
'plan note 5', '2024-06-30 10:46:35.093141+00', 'contact 5', '2024-06-30 10:46:35.093141+00'),
('PLAN6', 6, 'แผนสุขภาวะชุมชน', 'Healthy Community Strengthening Plan', 'Goal 6',
'ready 6', '2024-06-30 10:46:35.093141+00', 'type 6', '2024-06-30 10:46:35.093141+00', 'goal details 6', '2024-06-30 10:46:35.093141+00', 'activity 6','2024-06-30 10:46:35.093141+00',
'plan note 6', '2024-06-30 10:46:35.093141+00', 'contact 6', '2024-06-30 10:46:35.093141+00'),
('PLAN7', 7, 'แผนสุขภาวะเด็ก เยาวชน และครอบครัว', 'Healthy Child, Youth, and Family Promotion Plan', 'Goal 7',
'ready 7', '2024-06-30 10:46:35.093141+00', 'type 7', '2024-06-30 10:46:35.093141+00', 'goal details 7', '2024-06-30 10:46:35.093141+00', 'activity 7','2024-06-30 10:46:35.093141+00',
'plan note 7', '2024-06-30 10:46:35.093141+00', 'contact 7', '2024-06-30 10:46:35.093141+00'),
('PLAN8', 8, 'แผนสร้างเสริมสุขภาวะในองค์กร', 'Health Promotion in Organizations Plan', 'Goal 8',
'ready 8', '2024-06-30 10:46:35.093141+00', 'type 8', '2024-06-30 10:46:35.093141+00', 'goal details 8', '2024-06-30 10:46:35.093141+00', 'activity 8','2024-06-30 10:46:35.093141+00',
'plan note 8', '2024-06-30 10:46:35.093141+00', 'contact 8', '2024-06-30 10:46:35.093141+00'),
('PLAN9', 9, 'แผนส่งเสริมกิจกรรมทางกาย', 'Physical Activity Promotion Plan', 'Goal 9',
'ready 9', '2024-06-30 10:46:35.093141+00', 'type 9', '2024-06-30 10:46:35.093141+00', 'goal details 9', '2024-06-30 10:46:35.093141+00', 'activity 9','2024-06-30 10:46:35.093141+00',
'plan note 9', '2024-06-30 10:46:35.093141+00', 'contact 9', '2024-06-30 10:46:35.093141+00'),
('PLAN10', 10, 'แผนระบบสื่อและวิถีสุขภาวะทางปัญญา', 'Health Media System and Spiritual Health Pathway Plan', 'Goal 10',
'ready 10', '2024-06-30 10:46:35.093141+00', 'type 10', '2024-06-30 10:46:35.093141+00', 'goal details 10', '2024-06-30 10:46:35.093141+00', 'activity 10','2024-06-30 10:46:35.093141+00',
'plan note 10', '2024-06-30 10:46:35.093141+00', 'contact 10', '2024-06-30 10:46:35.093141+00'),
('PLAN11', 11, 'แผนสร้างสรรค์โอกาสและนวัตกรรมสุขภาวะ', 'Health Promotion Innovation and Open Grant Plan', 'Goal 11',
'ready 11', '2024-06-30 10:46:35.093141+00', 'type 11', '2024-06-30 10:46:35.093141+00', 'goal details 11', '2024-06-30 10:46:35.093141+00', 'activity 11','2024-06-30 10:46:35.093141+00',
'plan note 11', '2024-06-30 10:46:35.093141+00', 'contact 11', '2024-06-30 10:46:35.093141+00'),
('PLAN12', 12, 'แผนสนับสนุนการสร้างเสริมสุขภาพ ผ่านระบบบริการสุขภาพ', 'Health Promotion in Health Service System Plan', 'Goal 12',
'ready 12', '2024-06-30 10:46:35.093141+00', 'type 12', '2024-06-30 10:46:35.093141+00', 'goal details 12', '2024-06-30 10:46:35.093141+00', 'activity 12','2024-06-30 10:46:35.093141+00',
'plan note 12', '2024-06-30 10:46:35.093141+00', 'contact 12', '2024-06-30 10:46:35.093141+00'),
('PLAN14', 13, 'แผนอาหารเพื่อสุขภาวะ', 'Healthy Food Promotion Plan', 'Goal 14',
'ready 14', '2024-06-30 10:46:35.093141+00', 'type 14', '2024-06-30 10:46:35.093141+00', 'goal details 14', '2024-06-30 10:46:35.093141+00', 'activity 14', '2024-06-30 10:46:35.093141+00',
'plan note 14', '2024-06-30 10:46:35.093141+00', 'contact 14', '2024-06-30 10:46:35.093141+00'),
('PLAN15', 14, 'แผนสร้างเสริมความเข้าใจสุขภาวะ', 'Health Literacy Promotion Plan', 'Goal 15',
'ready 15', '2024-06-30 10:46:35.093141+00', 'type 15', '2024-06-30 10:46:35.093141+00', 'goal details 15', '2024-06-30 10:46:35.093141+00', 'activity 15', '2024-06-30 10:46:35.093141+00',
'plan note 15', '2024-06-30 10:46:35.093141+00', 'contact 15', '2024-06-30 10:46:35.093141+00')
;

INSERT INTO plan (name, user_id,  topic, topic_en, goal) VALUES ('ADMIN', 15, 'ไออาร์', 'IR', 'Goal Admin');


-- +goose Down
ALTER TABLE plan DROP COLUMN user_id;
DROP TABLE plan;