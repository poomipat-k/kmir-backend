-- +goose Up
CREATE TABLE plan (
  id SERIAL PRIMARY KEY NOT NULL,
  name VARCHAR(255) UNIQUE NOT NULL,
  user_id INT NOT NULL REFERENCES users (id),
  topic VARCHAR(255) NOT NULL,
  topic_en VARCHAR(255) NOT NULL,
  goal VARCHAR(512) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL
);

INSERT 
INTO plan (name, user_id, topic, topic_en, goal) 
VALUES 
('PLAN1', 1,'แผนควบคุมยาสูบ', 'Tobacco Control', 'Exchange knowledge with international audiences and participating in the discussion in the global and regional networks'), 
('PLAN2', 2,'แผนควบคุมเครื่องดื่มแอลกอฮอล์และสิ่งเสพติด', 'Alcohol and Substance Abuse Control', 'Goal 2'),
('PLAN3', 3, 'แผนการจัดการความปลอดภัย และปัจจัยเสี่ยงทางสังคม', 'Road Safety and Social Risk Management Plan', 'Goal 3'),
('PLAN4', 4, 'แผนควบคุมปัจจัยเสี่ยงทางสุขภาพ', 'Health Risk Control Plan', 'Goal 4'),
('PLAN5', 5, 'แผนสุขภาวะประชากรกลุ่มเฉพาะ', 'Health Promotion Plan for Vulnerable Populations', 'Goal 5'),
('PLAN6', 6, 'แผนสุขภาวะชุมชน', 'Healthy Community Strengthening Plan', 'Goal 6'),
('PLAN7', 7, 'แผนสุขภาวะเด็ก เยาวชน และครอบครัว', 'Healthy Child, Youth, and Family Promotion Plan', 'Goal 7'),
('PLAN8', 8, 'แผนสร้างเสริมสุขภาวะในองค์กร', 'Health Promotion in Organizations Plan', 'Goal 8'),
('PLAN9', 9, 'แผนส่งเสริมกิจกรรมทางกาย', 'Physical Activity Promotion Plan', 'Goal 9'),
('PLAN10', 10, 'แผนระบบสื่อและวิถีสุขภาวะทางปัญญา', 'Health Media System and Spiritual Health Pathway Plan', 'Goal 10'),
('PLAN11', 11, 'แผนสร้างสรรค์โอกาสและนวัตกรรมสุขภาวะ', 'Health Promotion Innovation and Open Grant Plan', 'Goal 11'),
('PLAN12', 12, 'แผนสนับสนุนการสร้างเสริมสุขภาพ ผ่านระบบบริการสุขภาพ', 'Health Promotion in Health Service System Plan', 'Goal 12'),
('PLAN14', 13, 'แผนอาหารเพื่อสุขภาวะ', 'Healthy Food Promotion Plan', 'Goal 14'),
('PLAN15', 14, 'แผนสร้างเสริมความเข้าใจสุขภาวะ', 'Health Literacy Promotion Plan', 'Goal 15'),
('ADMIN', 15, 'ไออาร์', 'IR', 'Goal Admin')
;


-- +goose Down
ALTER TABLE plan DROP COLUMN user_id;
DROP TABLE plan;