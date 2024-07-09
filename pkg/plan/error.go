package plan

type ScoreRequiredError struct{}

func (e ScoreRequiredError) Error() string {
	return "score is required"
}

type ScoreValueOutOfRangeError struct{}

func (e ScoreValueOutOfRangeError) Error() string {
	return "score is ValueOutOfRange 1-10"
}

type ReadinessWillingnessRequiredError struct{}

func (e ReadinessWillingnessRequiredError) Error() string {
	return "readinessWillingness is required"
}
