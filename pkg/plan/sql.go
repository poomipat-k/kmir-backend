package plan

const getAllPreviewPlanSQL = `SELECT id, name, user_id, topic, topic_en, goal FROM plan ORDER BY id ASC;`

const canAccessPlanDetailsSQL = `SELECT plan.id as plan_id
FROM plan INNER JOIN users ON plan.user_id = users.id 
WHERE plan.name = $1 AND users.username = $2;`

const getPlanDetailsSQL = `
SELECT
plan.id as plan_id,
plan.name,
plan.topic as topic,
plan.topic_en as topic_en,
plan.readiness_willingness,
plan.readiness_willingness_updated_at,
plan.readiness_willingness_updated_by,
plan.ir_goal_type,
plan.ir_goal_type_updated_at,
plan.ir_goal_type_updated_by,
plan.ir_goal_details,
plan.ir_goal_details_updated_at,
plan.ir_goal_details_updated_by,
plan.proposed_activity,
plan.proposed_activity_updated_at,
plan.proposed_activity_updated_by,
plan.plan_note,
plan.plan_note_updated_at,
plan.plan_note_updated_by,
plan.contact_person,
plan.contact_person_updated_at,
plan.contact_person_updated_by,
plan.updated_at,
plan.updated_by
FROM plan WHERE plan.name = $1;
`
