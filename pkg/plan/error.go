package plan

type ScoreRequiredError struct{}

func (e ScoreRequiredError) Error() string {
	return "score is required"
}

type ScoreValueOutOfRangeError struct{}

func (e ScoreValueOutOfRangeError) Error() string {
	return "score is ValueOutOfRange 1-10"
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

type SomePlanScoreIsMissing struct{}

func (e SomePlanScoreIsMissing) Error() string {
	return "some plan score is missing"
}

type SomeIrWorkGoalIsMissing struct{}

func (e SomeIrWorkGoalIsMissing) Error() string {
	return "some ir work goal is missing"
}

type SomeProposedActivityIsMissing struct{}

func (e SomeProposedActivityIsMissing) Error() string {
	return "some proposed activity is missing"
}

type SomePlanNoteIsMissing struct{}

func (e SomePlanNoteIsMissing) Error() string {
	return "some plan note is missing"
}

type AdminNoteRequiredError struct{}

func (e AdminNoteRequiredError) Error() string {
	return "adminNote is required"
}
