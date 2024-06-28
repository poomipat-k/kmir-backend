package plan

const getAllPreviewPlanSQL = `SELECT id, name, user_id, topic, topic_en, goal FROM plan ORDER BY id ASC;`

const canAccessPlanDetailsSQL = `SELECT plan.id as plan_id
FROM plan INNER JOIN users ON plan.user_id = users.id 
WHERE plan.name = $1 AND users.username = $2;`
