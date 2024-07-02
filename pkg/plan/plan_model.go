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
	ReadinessWillingness          *string          `json:"readinessWillingness,omitempty"`
	ReadinessWillingnessUpdatedAt *time.Time       `json:"readinessWillingnessUpdatedAt,omitempty"`
	AssessmentScore               *AssessmentScore `json:"assessmentScore,omitempty"`
}

type AssessmentScore struct {
	CriteriaId int `json:"criteriaId,omitempty"`
}
