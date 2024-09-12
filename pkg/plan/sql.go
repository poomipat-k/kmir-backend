package plan

const getAllPreviewPlanSQL = `SELECT id, name, user_id, topic, topic_en, ir_goal_details FROM plan ORDER BY id ASC;`

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
plan.ir_goal_type,
plan.ir_goal_details,
plan.proposed_activity,
plan.plan_note,
plan.updated_at,
plan.updated_by,
plan.ir_goal_type_updated_at,
plan.ir_goal_type_updated_by,
plan.ir_goal_details_updated_at,
plan.ir_goal_details_updated_by,
plan.proposed_activity_updated_at,
plan.proposed_activity_updated_by,
plan.plan_note_updated_at,
plan.plan_note_updated_by,
plan.contact_person_updated_at,
plan.contact_person_updated_by
FROM plan
WHERE plan.for_admin = false
ORDER BY plan.id
;
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

const adminGetAllPlanScoreDetailsSQL = `
SELECT
plan_id,
criteria_id,
criteria_order, 
score,
created_at
FROM
(
SELECT
assessment_score.plan_id as plan_id,
assessment_score.assessment_criteria_id as criteria_id,
assessment_criteria.order_number as criteria_order,
users.user_role as user_role,
score, 
assessment_score.created_at as created_at,
ROW_NUMBER() OVER (
PARTITION BY assessment_score.plan_id, assessment_score.user_id, year 
ORDER BY assessment_score.created_at DESC, assessment_criteria_id ASC) 
as row_num FROM assessment_score 
INNER JOIN plan ON plan.id = assessment_score.plan_id
INNER JOIN assessment_criteria ON assessment_criteria.id = assessment_score.assessment_criteria_id
INNER JOIN users ON users.id = assessment_score.user_id
)
WHERE row_num <= $1 
AND user_role = 'admin'
AND created_at >= $2
AND created_at < $3
ORDER BY plan_id ASC, criteria_order ASC
;
`

const getAdminNote = "SELECT note from admin_note LIMIT 1;"

const getOnlyLatestScoreSQL = `
SELECT plan_id, user_role, created_at 
FROM 
(
SELECT assessment_score.plan_id, users.user_role,  assessment_score.created_at,
ROW_NUMBER() OVER (
PARTITION BY assessment_score.plan_id
ORDER BY assessment_score.created_at DESC) 
as row_num
FROM assessment_score
INNER JOIN users ON assessment_score.user_id = users.id
ORDER BY assessment_score.created_at
)
WHERE row_num <= 1;
;`
