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
  readiness_willingness_updated_by VARCHAR(16),
  ir_goal_type VARCHAR(64),
  ir_goal_type_updated_at TIMESTAMP WITH TIME ZONE,
  ir_goal_type_updated_by VARCHAR(16),
  ir_goal_details TEXT,
  ir_goal_details_updated_at TIMESTAMP WITH TIME ZONE,
  ir_goal_details_updated_by VARCHAR(16),
  proposed_activity TEXT,
  proposed_activity_updated_at TIMESTAMP WITH TIME ZONE,
  proposed_activity_updated_by VARCHAR(16),
  plan_note TEXT,
  plan_note_updated_at TIMESTAMP WITH TIME ZONE,
  plan_note_updated_by VARCHAR(16),
  contact_person TEXT,
  contact_person_updated_at TIMESTAMP WITH TIME ZONE,
  contact_person_updated_by VARCHAR(16),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
  updated_by VARCHAR(16)
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
'Korem ipsum dolor sit amet, consectetur adipiscing elit. Etiam eu turpis molestie, dictum est a, mattis tellus. Sed dignissim, metus nec fringilla accumsan, risus sem sollicitudin lacus, ut interdum tellus elit sed risus. Maecenas eget condimentum velit, sit amet feugiat lectus. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent auctor purus luctus enim egestas, ac scelerisque ante pulvinar. Donec ut rhoncus ex. Suspendisse ac rhoncus nil, eu tempor urna. Curabitur vel bibendum lorem. Morbi convallis convallis diam sit amet lacinia. Aliquam in elementum tellus.', 
'2024-06-30 10:46:35.093141+00', 'user', 'type 1', '2024-06-30 10:46:35.093141+00', 'user', 
'Enhance ThaiHealth''s national policy advocacy for tobacco control by integrating global best practices and strengthening compliance with the WHO Framework Convention on Tobacco Control (FCTC), with a focus on emerging challenges such as e-cigarettes', 
'2024-06-30 10:46:35.093141+00', 'user', 
'2025
Curabitur tempor quis eros tempus lacinia. Nam bibendum pellentesque quam a convallis. Sed ut vulputate nisi. Integer in felis sed leo vestibulum venenatis. Suspendisse quis arcu sem. Aenean feugiat ex eu vestibulum vestibulum. Morbi a eleifend magna. Nam metus lacus, porttitor eu mauris a, blandit ultrices nibh. Mauris sit amet magna non ligula vestibulum eleifend. Nulla varius volutpat turpis sed lacinia. Nam eget mi in purus lobortis eleifend. Sed nec ante dictum sem condimentum ullamcorper quis venenatis nisi. Proin vitae facilisis nisi, ac posuere leo.

2026
Curabitur tempor quis eros tempus lacinia. Nam bibendum pellentesque quam a convallis. Sed ut vulputate nisi. Integer in felis sed leo vestibulum venenatis. Suspendisse quis arcu sem. Aenean feugiat ex eu vestibulum vestibulum. Morbi a eleifend magna. Nam metus lacus, porttitor eu mauris a, blandit ultrices nibh. Mauris sit amet magna non ligula vestibulum eleifend. Nulla varius volutpat turpis sed lacinia. Nam eget mi in purus lobortis eleifend. Sed nec ante dictum sem condimentum ullamcorper quis venenatis nisi. Proin vitae facilisis nisi, ac posuere leo.', 
'2024-06-30 10:46:35.093141+00', 'user', 
'plan note 1 Korem ipsum dolor sit amet, consectetur adipiscing elit. Etiam eu turpis molestie, dictum est a, mattis tellus. Sed dignissim, metus nec fringilla accumsan, risus sem sollicitudin lacus, ut interdum tellus elit sed risus. Maecenas eget condimentum velit, sit amet feugiat lectus. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent auctor purus luctus enim egestas, ac scelerisque ante pulvinar. Donec ut rhoncus ex. Suspendisse ac rhoncus nil, eu tempor urna. Curabitur vel bibendum lorem. Morbi convallis convallis diam sit amet lacinia. Aliquam in elementum tellus.

Curabitur tempor quis eros tempus lacinia. Nam bibendum pellentesque quam a convallis. Sed ut vulputate nisi. Integer in felis sed leo vestibulum venenatis. Suspendisse quis arcu sem. Aenean feugiat ex eu vestibulum vestibulum. Morbi a eleifend magna. Nam metus lacus, porttitor eu mauris a, blandit ultrices nibh. Mauris sit amet magna non ligula vestibulum eleifend. Nulla varius volutpat turpis sed lacinia. Nam eget mi in purus lobortis eleifend. Sed nec ante dictum sem condimentum ullamcorper quis venenatis nisi. Proin vitae facilisis nisi, ac posuere leo.

download report at https://phethai.org/th/about-us', 
'2024-06-30 10:46:35.093141+00', 'user', 
'นายรังสรร มั่นคง (กอล์ฟ) 
ผู้เชี่ยวชาญด้านวิเทศสัมพันธ์  สำนักพัฒนาภาคีสัมพันธ์และวิเทศสัมพันธ์ (สภส.)
Email: Rungsun@thaihealth.or.th', 
'2024-06-30 10:46:35.093141+00', 'user', 'user'), 
('PLAN2', 2,'แผนควบคุมเครื่องดื่มแอลกอฮอล์และสิ่งเสพติด', 'Alcohol and Substance Abuse Control', 'Goal 2',
'ready 2', '2024-06-30 10:46:35.093141+00', 'user', 'type 2', '2024-06-30 10:46:35.093141+00', 'user', 'goal details 2', '2024-06-30 10:46:35.093141+00', 'user', 'activity 2','2024-06-30 10:46:35.093141+00', 'user',
'plan note 2', '2024-06-30 10:46:35.093141+00', 'user', 'contact 2', '2024-06-30 10:46:35.093141+00', 'user', 'user'),
('PLAN3', 3, 'แผนการจัดการความปลอดภัย และปัจจัยเสี่ยงทางสังคม', 'Road Safety and Social Risk Management Plan', 'Goal 3',
'ready 3', '2024-06-30 10:46:35.093141+00', 'user', 'type 3', '2024-06-30 10:46:35.093141+00', 'user', 'goal details 3', '2024-06-30 10:46:35.093141+00', 'user', 'activity 3','2024-06-30 10:46:35.093141+00', 'user',
'plan note 3', '2024-06-30 10:46:35.093141+00', 'user', 'contact 3', '2024-06-30 10:46:35.093141+00', 'user', 'user'),
('PLAN4', 4, 'แผนควบคุมปัจจัยเสี่ยงทางสุขภาพ', 'Health Risk Control Plan', 'Goal 4',
'ready 4', '2024-06-30 10:46:35.093141+00', 'user', 'type 4', '2024-06-30 10:46:35.093141+00', 'user', 'goal details 4', '2024-06-30 10:46:35.093141+00', 'user', 'activity 4','2024-06-30 10:46:35.093141+00', 'user',
'plan note 4', '2024-06-30 10:46:35.093141+00', 'user', 'contact 4', '2024-06-30 10:46:35.093141+00', 'user', 'user'),
('PLAN5', 5, 'แผนสุขภาวะประชากรกลุ่มเฉพาะ', 'Health Promotion Plan for Vulnerable Populations', 'Goal 5',
'ready 5', '2024-06-30 10:46:35.093141+00', 'user', 'type 5', '2024-06-30 10:46:35.093141+00', 'user', 'goal details 5', '2024-06-30 10:46:35.093141+00', 'user', 'activity 5','2024-06-30 10:46:35.093141+00', 'user',
'plan note 5', '2024-06-30 10:46:35.093141+00', 'user', 'contact 5', '2024-06-30 10:46:35.093141+00', 'user', 'user'),
('PLAN6', 6, 'แผนสุขภาวะชุมชน', 'Healthy Community Strengthening Plan', 'Goal 6',
'ready 6', '2024-06-30 10:46:35.093141+00', 'user', 'type 6', '2024-06-30 10:46:35.093141+00', 'user', 'goal details 6', '2024-06-30 10:46:35.093141+00', 'user', 'activity 6','2024-06-30 10:46:35.093141+00', 'user',
'plan note 6', '2024-06-30 10:46:35.093141+00', 'user', 'contact 6', '2024-06-30 10:46:35.093141+00', 'user', 'user'),
('PLAN7', 7, 'แผนสุขภาวะเด็ก เยาวชน และครอบครัว', 'Healthy Child, Youth, and Family Promotion Plan', 'Goal 7',
'ready 7', '2024-06-30 10:46:35.093141+00', 'user', 'type 7', '2024-06-30 10:46:35.093141+00', 'user', 'goal details 7', '2024-06-30 10:46:35.093141+00', 'user', 'activity 7','2024-06-30 10:46:35.093141+00', 'user',
'plan note 7', '2024-06-30 10:46:35.093141+00', 'user', 'contact 7', '2024-06-30 10:46:35.093141+00', 'user', 'user'),
('PLAN8', 8, 'แผนสร้างเสริมสุขภาวะในองค์กร', 'Health Promotion in Organizations Plan', 'Goal 8',
'ready 8', '2024-06-30 10:46:35.093141+00', 'user', 'type 8', '2024-06-30 10:46:35.093141+00', 'user', 'goal details 8', '2024-06-30 10:46:35.093141+00', 'user', 'activity 8','2024-06-30 10:46:35.093141+00', 'user',
'plan note 8', '2024-06-30 10:46:35.093141+00', 'user', 'contact 8', '2024-06-30 10:46:35.093141+00', 'user', 'user'),
('PLAN9', 9, 'แผนส่งเสริมกิจกรรมทางกาย', 'Physical Activity Promotion Plan', 'Goal 9',
'ready 9', '2024-06-30 10:46:35.093141+00', 'user', 'type 9', '2024-06-30 10:46:35.093141+00', 'user', 'goal details 9', '2024-06-30 10:46:35.093141+00', 'user', 'activity 9','2024-06-30 10:46:35.093141+00', 'user',
'plan note 9', '2024-06-30 10:46:35.093141+00', 'user', 'contact 9', '2024-06-30 10:46:35.093141+00', 'user', 'user'),
('PLAN10', 10, 'แผนระบบสื่อและวิถีสุขภาวะทางปัญญา', 'Health Media System and Spiritual Health Pathway Plan', 'Goal 10',
'ready 10', '2024-06-30 10:46:35.093141+00', 'user', 'type 10', '2024-06-30 10:46:35.093141+00', 'user', 'goal details 10', '2024-06-30 10:46:35.093141+00', 'user', 'activity 10','2024-06-30 10:46:35.093141+00', 'user',
'plan note 10', '2024-06-30 10:46:35.093141+00', 'user', 'contact 10', '2024-06-30 10:46:35.093141+00', 'user', 'user'),
('PLAN11', 11, 'แผนสร้างสรรค์โอกาสและนวัตกรรมสุขภาวะ', 'Health Promotion Innovation and Open Grant Plan', 'Goal 11',
'ready 11', '2024-06-30 10:46:35.093141+00', 'user', 'type 11', '2024-06-30 10:46:35.093141+00', 'user', 'goal details 11', '2024-06-30 10:46:35.093141+00', 'user', 'activity 11','2024-06-30 10:46:35.093141+00', 'user',
'plan note 11', '2024-06-30 10:46:35.093141+00', 'user', 'contact 11', '2024-06-30 10:46:35.093141+00', 'user', 'user'),
('PLAN12', 12, 'แผนสนับสนุนการสร้างเสริมสุขภาพ ผ่านระบบบริการสุขภาพ', 'Health Promotion in Health Service System Plan', 'Goal 12',
'ready 12', '2024-06-30 10:46:35.093141+00', 'user', 'type 12', '2024-06-30 10:46:35.093141+00', 'user', 'goal details 12', '2024-06-30 10:46:35.093141+00', 'user', 'activity 12','2024-06-30 10:46:35.093141+00', 'user',
'plan note 12', '2024-06-30 10:46:35.093141+00', 'user', 'contact 12', '2024-06-30 10:46:35.093141+00', 'user', 'user'),
('PLAN14', 13, 'แผนอาหารเพื่อสุขภาวะ', 'Healthy Food Promotion Plan', 'Goal 14',
'ready 14', '2024-06-30 10:46:35.093141+00', 'user', 'type 14', '2024-06-30 10:46:35.093141+00', 'user', 'goal details 14', '2024-06-30 10:46:35.093141+00', 'user', 'activity 14', '2024-06-30 10:46:35.093141+00', 'user',
'plan note 14', '2024-06-30 10:46:35.093141+00', 'user', 'contact 14', '2024-06-30 10:46:35.093141+00', 'user', 'user'),
('PLAN15', 14, 'แผนสร้างเสริมความเข้าใจสุขภาวะ', 'Health Literacy Promotion Plan', 'Goal 15',
'ready 15', '2024-06-30 10:46:35.093141+00', 'user', 'type 15', '2024-06-30 10:46:35.093141+00', 'user', 'goal details 15', '2024-06-30 10:46:35.093141+00', 'user', 'activity 15', '2024-06-30 10:46:35.093141+00', 'user',
'plan note 15', '2024-06-30 10:46:35.093141+00', 'user', 'contact 15', '2024-06-30 10:46:35.093141+00', 'user', 'user')
;

INSERT INTO plan (name, user_id,  topic, topic_en, goal, updated_by) VALUES ('ADMIN', 15, 'ไออาร์', 'IR', 'Goal Admin', 'admin');


-- +goose Down
ALTER TABLE plan DROP COLUMN user_id;
DROP TABLE plan;