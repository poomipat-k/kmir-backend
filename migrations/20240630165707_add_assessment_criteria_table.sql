-- +goose Up
CREATE TABLE assessment_criteria (
  id SERIAL PRIMARY KEY NOT NULL,
  version INT NOT NULL,
  category VARCHAR(64) NOT NULL,
  display VARCHAR(255) NOT NULL,
  order_number INT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL
);

-- INSERT INTO assessment_criteria (version, category, display, order_number)
-- VALUES 
-- (1, 'willingness', 'The plan''s international activities are well-aligned with the foundation''s overall strategic objectives.', 1),
-- (1, 'willingness', 'The plan is highly aware of the benefits and importance of engaging in international activities.', 2),
-- (1, 'willingness', 'The plan feels strongly supported by the International Relations Section and Executives in pursuing international activities.', 3),
-- (1, 'capacity', 'The plan uses resources efficiently for international activities.', 4),
-- (1, 'capacity', 'The plan is highly competent in handling international affairs.', 5),
-- (1, 'capacity', 'The quality of the plan''s past and existing international partnerships is excellent', 6),
-- (1, 'capacity', 'The plan effectively measures the outcomes of its international activities.', 7)
-- ;

INSERT INTO assessment_criteria (version, category, display, order_number)
VALUES 
(1, 'willingness', 'กิจกรรมของแผนที่เกี่ยวข้องกับงานด้านต่างประเทศ<u>มีความสอดคล้องอย่างยิ่ง</u>กับเป้าหมายและกลยุทธ์ของงานวิเทศสัมพันธ์ สสส. ', 1),
(1, 'willingness', 'กิจกรรมโดยรวมของแผนแสดงให้เห็น<u>ความสำคัญและคุณค่าของงานด้านต่างประเทศ</u>', 2),
(1, 'willingness', '<u>ทีมวิเทศสัมพันธ์และผู้บริหารระดับสูงกระตุ้นและสนับสนุนให้</u>แผนดำเนินกิจกรรมที่เกี่ยวข้องด้านต่างประเทศอย่างต่อเนื่อง', 3),
(1, 'capacity', 'แผนมีการใช้<u>ทรัพยากร</u> (งบประมาณ กำลังคน เครื่องมือวิชาการ) ในกิจกรรมที่เกี่ยวข้องกับงานด้านต่างประเทศได้<u>อย่างมีประสิทธิภาพ</u>', 4),
(1, 'capacity', 'แผนแสดงให้เห็น<u>ทักษะและความสามารถในระดับสูง</u>ต่อการบริหารจัดการงานด้านต่างประเทศ', 5),
(1, 'capacity', '<u>ประสบการณ์</u>ของแผนในการสร้างและต่อยอดความร่วมมือกับภาคีเครือข่ายระหว่างประเทศตั้งแต่อดีตจนถึงปัจจุบัน<u>มีประสิทธิภาพสูง</u>', 6),
(1, 'capacity', 'แผนมีกลไก<u>การติดตามและวิเคราะห์ผลลัพธ์</u>จากกิจกรรมด้านต่างประเทศได้อย่างเป็นระบบ', 7)
;

-- +goose Down
DROP TABLE assessment_criteria;