package plan

import (
	"time"
)

type PlanPreview struct {
	Id      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	UserId  int    `json:"userId,omitempty"`
	Topic   string `json:"topic,omitempty"`
	TopicEn string `json:"topicEn,omitempty"`
	Goal    string `json:"goal,omitempty"`
}

type PlanDetails struct {
	PlanId                        int                  `json:"planId,omitempty"`
	Name                          string               `json:"name,omitempty"`
	Topic                         string               `json:"topic,omitempty"`
	TopicEn                       string               `json:"topicEn,omitempty"`
	TopicShort                    string               `json:"topicShort,omitempty"`
	ReadinessWillingness          *string              `json:"readinessWillingness,omitempty"`
	ReadinessWillingnessUpdatedAt *time.Time           `json:"readinessWillingnessUpdatedAt,omitempty"`
	ReadinessWillingnessUpdatedBy *string              `json:"readinessWillingnessUpdatedBy,omitempty"`
	AssessmentCriteria            []AssessmentCriteria `json:"assessmentCriteria,omitempty"`
	AssessmentScore               []AssessmentScore    `json:"assessmentScore,omitempty"`
	IrGoalType                    *string              `json:"irGoalType,omitempty"`
	IrGoalTypeUpdatedAt           *time.Time           `json:"irGoalTypeUpdatedAt,omitempty"`
	IrGoalTypeUpdatedBy           *string              `json:"irGoalTypeUpdatedBy,omitempty"`
	IrGoalDetails                 *string              `json:"irGoalDetails,omitempty"`
	IrGoalDetailsUpdatedAt        *time.Time           `json:"irGoalDetailsUpdatedAt,omitempty"`
	IrGoalDetailsUpdatedBy        *string              `json:"irGoalDetailsUpdatedBy,omitempty"`
	ProposedActivity              *string              `json:"proposedActivity,omitempty"`
	ProposedActivityUpdatedAt     *time.Time           `json:"proposedActivityUpdatedAt,omitempty"`
	ProposedActivityUpdatedBy     *string              `json:"proposedActivityUpdatedBy,omitempty"`
	PlanNote                      *string              `json:"planNote,omitempty"`
	PlanNoteUpdatedAt             *time.Time           `json:"planNoteUpdatedAt,omitempty"`
	PlanNoteUpdatedBy             *string              `json:"planNoteUpdatedBy,omitempty"`
	ContactPerson                 *string              `json:"contactPerson,omitempty"`
	ContactPersonUpdatedAt        *time.Time           `json:"contactPersonUpdatedAt,omitempty"`
	ContactPersonUpdatedBy        *string              `json:"contactPersonUpdatedBy,omitempty"`
	UpdatedAt                     *time.Time           `json:"updatedAt,omitempty"`
	UpdatedBy                     *string              `json:"updatedBy,omitempty"`
}

type AdminAllPlansDetailsResponse struct {
	AssessmentCriteria []AssessmentCriteria           `json:"assessmentCriteria,omitempty"`
	PlanDetails        []AdminDashboardPlanDetailsRow `json:"planDetails,omitempty"`
	AdminNote          string                         `json:"adminNote,omitempty"`
	LatestScores       []LatestScoreTimestamp         `json:"latestScores,omitempty"`
}

type LatestScoreTimestamp struct {
	PlanId    int        `json:"planId,omitempty"`
	UserRole  string     `json:"userRole,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}

type AdminDashboardPlanDetailsRow struct {
	PlanId           int                  `json:"planId,omitempty"`
	Name             string               `json:"name,omitempty"`
	Topic            string               `json:"topic,omitempty"`
	TopicEn          string               `json:"topicEn,omitempty"`
	TopicShort       string               `json:"topicShort,omitempty"`
	AssessmentScore  []AssessmentScoreRow `json:"assessmentScore"`
	ProposedActivity *string              `json:"proposedActivity,omitempty"`
	PlanNote         *string              `json:"planNote,omitempty"`
	// Field to detect change logs
	UpdatedAt                     *time.Time `json:"updatedAt,omitempty"`
	UpdatedBy                     *string    `json:"updatedBy,omitempty"`
	ReadinessWillingnessUpdatedAt *time.Time `json:"readinessWillingnessUpdatedAt,omitempty"`
	ReadinessWillingnessUpdatedBy *string    `json:"readinessWillingnessUpdatedBy,omitempty"`
	IrGoalTypeUpdatedAt           *time.Time `json:"irGoalTypeUpdatedAt,omitempty"`
	IrGoalTypeUpdatedBy           *string    `json:"irGoalTypeUpdatedBy,omitempty"`
	IrGoalDetailsUpdatedAt        *time.Time `json:"irGoalDetailsUpdatedAt,omitempty"`
	IrGoalDetailsUpdatedBy        *string    `json:"irGoalDetailsUpdatedBy,omitempty"`
	ProposedActivityUpdatedAt     *time.Time `json:"proposedActivityUpdatedAt,omitempty"`
	ProposedActivityUpdatedBy     *string    `json:"proposedActivityUpdatedBy,omitempty"`
	PlanNoteUpdatedAt             *time.Time `json:"planNoteUpdatedAt,omitempty"`
	PlanNoteUpdatedBy             *string    `json:"planNoteUpdatedBy,omitempty"`
	ContactPersonUpdatedAt        *time.Time `json:"contactPersonUpdatedAt,omitempty"`
	ContactPersonUpdatedBy        *string    `json:"contactPersonUpdatedBy,omitempty"`
}

type AssessmentCriteria struct {
	CriteriaId  int    `json:"criteriaId,omitempty"`
	OrderNumber int    `json:"orderNumber,omitempty"`
	Category    string `json:"category,omitempty"`
	Display     string `json:"display,omitempty"`
}

type AssessmentScore struct {
	PlanId           int        `json:"planId,omitempty"`
	CriteriaOrder    int        `json:"criteriaOrder,omitempty"`
	CriteriaCategory string     `json:"criteriaCategory,omitempty"`
	UserRole         string     `json:"userRole,omitempty"`
	Year             int        `json:"year,omitempty"`
	Score            int        `json:"score,omitempty"`
	CreatedAt        *time.Time `json:"createdAt,omitempty"`
}

type AssessmentScoreRow struct {
	PlanId           int        `json:"planId,omitempty"`
	CriteriaId       int        `json:"criteriaId,omitempty"`
	CriteriaOrder    int        `json:"criteriaOrder,omitempty"`
	UserRole         string     `json:"userRole,omitempty"`
	Year             int        `json:"year,omitempty"`
	Score            int        `json:"score,omitempty"`
	CreatedAt        *time.Time `json:"createdAt,omitempty"`
	CriteriaCategory string     `json:"criteriaCategory,omitempty"`
	CriteriaDisplay  string     `json:"criteriaDisplay,omitempty"`
}

type EditPlanRequest struct {
	PlanName             string         `json:"planName,omitempty"`
	ReadinessWillingness *string        `json:"readinessWillingness,omitempty"`
	AssessmentScore      map[string]int `json:"assessmentScore,omitempty"`
	IrGoalType           *string        `json:"irGoalType,omitempty"`
	IrGoalDetails        *string        `json:"irGoalDetails,omitempty"`
	ProposedActivity     *string        `json:"proposedActivity,omitempty"`
	PlanNote             *string        `json:"planNote,omitempty"`
	ContactPerson        *string        `json:"contactPerson,omitempty"`
}

type AdminEditRequest struct {
	AssessmentScore  []map[string]int `json:"assessmentScore,omitempty"`
	ProposedActivity []string         `json:"proposedActivity,omitempty"`
	PlanNote         []string         `json:"planNote,omitempty"`
	AdminNote        *string          `json:"adminNote,omitempty"`
}

type AdminGetScoresRequest struct {
	FromYear int    `json:"fromYear,omitempty"`
	ToYear   int    `json:"toYear,omitempty"`
	Plan     string `json:"plan,omitempty"`
}
