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

type IrGoalTypeRequiredError struct{}

func (e IrGoalTypeRequiredError) Error() string {
	return "irGoalType is required"
}

type IrGoalDetailsRequiredError struct{}

func (e IrGoalDetailsRequiredError) Error() string {
	return "irGoalDetails is required"
}

type ProposedActivityRequiredError struct{}

func (e ProposedActivityRequiredError) Error() string {
	return "proposedActivity is required"
}

type PlanNoteRequiredError struct{}

func (e PlanNoteRequiredError) Error() string {
	return "planNote is required"
}

type ContactPersonRequiredError struct{}

func (e ContactPersonRequiredError) Error() string {
	return "contactPerson is required"
}
