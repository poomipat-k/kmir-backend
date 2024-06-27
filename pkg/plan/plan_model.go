package plan

type PlanPreview struct {
	Id      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	UserId  int    `json:"userId,omitempty"`
	Topic   string `json:"topic,omitempty"`
	TopicEn string `json:"topicEn,omitempty"`
	Goal    string `json:"goal,omitempty"`
}
