package plan

import "time"

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
	ReadinessWillingness          *string              `json:"readinessWillingness,omitempty"`
	ReadinessWillingnessUpdatedAt *time.Time           `json:"readinessWillingnessUpdatedAt,omitempty"`
	ReadinessWillingnessUpdatedBy *string              `json:"readinessWillingnessUpdatedBy,omitempty"`
	AssessmentCriteria            []AssessmentCriteria `json:"assessmentCriteria,omitempty"`
	AssessmentScore               *AssessmentScore     `json:"assessmentScore,omitempty"`
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

type AssessmentCriteria struct {
	Id          int    `json:"id,omitempty"`
	OrderNumber int    `json:"orderNumber,omitempty"`
	Category    string `json:"category,omitempty"`
	Display     string `json:"display,omitempty"`
}

type AssessmentScore struct {
	YearSummary []YearScore `json:"yearSummary,omitempty"`
}

type YearScore struct {
	Year  int          `json:"year,omitempty"`
	User  ScoreSummary `json:"user,omitempty"`
	Admin ScoreSummary `json:"admin,omitempty"`
}

type ScoreSummary struct {
	Scores    []Score    `json:"scores,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}

type Score struct {
	CriteriaId int `json:"criteriaId,omitempty"`
	Score      int `json:"score"`
}
