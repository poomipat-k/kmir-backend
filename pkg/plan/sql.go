package plan

const getAllPreviewPlanSQL = `SELECT id, name, user_id, topic, topic_en, goal FROM plan ORDER BY id ASC;`
