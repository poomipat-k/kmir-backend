package plan

const getAllPreviewPlanSQL = `SELECT id, name, user_id, topic, topic_en, goal FROM plan ORDER BY id ASC;`

const canAccessPlanDetailsSQL = `SELECT plan.id as plan_id
FROM plan INNER JOIN users ON plan.user_id = users.id 
WHERE plan.name = $1 AND users.username = $2;`

const canEditPlanSQL = `SELECT plan.id as plan_id
FROM plan INNER JOIN users ON plan.user_id = users.id 
WHERE plan.name = $1 AND users.username = $2;`

const getPlanDetailsSQL = `
SELECT
plan.id as plan_id,
plan.name,
plan.topic as topic,
plan.topic_en as topic_en,
plan.topic_short as topic_short,
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
FROM plan 
INNER JOIN users ON users.id = plan.user_id
WHERE plan.name = $1 AND users.username = $2;
`

const getAllAssessmentCriteriaSQL = "SELECT category, id, display, order_number FROM assessment_criteria LIMIT 7;"

const getPlanDetailsForAdminViewSQL = `
SELECT
plan.id as plan_id,
plan.name,
plan.topic as topic,
plan.topic_en as topic_en,
plan.topic_short as topic_short,
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
FROM plan 
INNER JOIN users ON users.id = plan.user_id
WHERE plan.name = $1;
`

const getAllPlanDetailsForAdminDashboardSQL = `
SELECT
plan.id as plan_id,
plan.name,
plan.topic as topic,
plan.topic_en as topic_en,
plan.topic_short as topic_short,

plan.readiness_willingness_updated_at,
plan.readiness_willingness_updated_by,

plan.ir_goal_type_updated_at,
plan.ir_goal_type_updated_by,

plan.ir_goal_details_updated_at,
plan.ir_goal_details_updated_by,
plan.proposed_activity,
plan.proposed_activity_updated_at,
plan.proposed_activity_updated_by,
plan.plan_note,
plan.plan_note_updated_at,
plan.plan_note_updated_by,

plan.contact_person_updated_at,
plan.contact_person_updated_by,
plan.updated_at,
plan.updated_by
FROM plan;
`

const getPlanScoreDetailsSQL = `
SELECT
plan_id,
criteria_id,
criteria_order, 
user_role,
year,
score,
created_at,
criteria_category,
criteria_display
FROM
(
SELECT
assessment_score.plan_id as plan_id,
assessment_score.assessment_criteria_id as criteria_id,
assessment_criteria.order_number as criteria_order,
assessment_score.user_id as user_id,
users.user_role as user_role,
year, 
score, 
assessment_score.created_at as created_at,
assessment_criteria.category as criteria_category,
assessment_criteria.display as criteria_display,
ROW_NUMBER() OVER (
PARTITION BY assessment_score.plan_id, assessment_score.user_id, year 
ORDER BY assessment_score.created_at DESC, assessment_criteria_id ASC) 
as row_num FROM assessment_score 
INNER JOIN plan ON plan.id = assessment_score.plan_id
INNER JOIN assessment_criteria ON assessment_criteria.id = assessment_score.assessment_criteria_id
INNER JOIN users ON users.id = assessment_score.user_id
WHERE plan.name = $1
)
WHERE row_num <= 7;
`
