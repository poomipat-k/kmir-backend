package plan

import (
	"fmt"
)

const allPlanCount = 14

func validateScore(scores map[string]int) (string, error) {
	for i := 1; i <= 7; i++ {
		key := fmt.Sprintf("q_%d", i)
		val, exist := scores[key]
		if !exist {
			return key, ScoreRequiredError{}
		}
		if val < 1 || val > 10 {
			return key, ScoreValueOutOfRangeError{}
		}
	}
	return "", nil
}

func validateEditPlanPayload(payload EditPlanRequest) (string, error) {
	if payload.IrGoalType != nil && *payload.IrGoalType == "" {
		return "irGoalType", IrGoalTypeRequiredError{}
	}
	if payload.IrGoalDetails != nil && *payload.IrGoalDetails == "" {
		return "irGoalDetails", IrGoalDetailsRequiredError{}
	}
	if payload.ProposedActivity != nil && *payload.ProposedActivity == "" {
		return "proposedActivity", ProposedActivityRequiredError{}
	}
	if payload.PlanNote != nil && *payload.PlanNote == "" {
		return "planNote", PlanNoteRequiredError{}
	}
	if payload.ContactPerson != nil && *payload.ContactPerson == "" {
		return "contactPerson", ContactPersonRequiredError{}
	}
	return "", nil
}

func validateAdminEditPayload(payload AdminEditRequest) (string, error) {
	if payload.AssessmentScore != nil {
		errName, err := validateAdminAssessmentScoreList(*payload.AssessmentScore)
		if err != nil {
			return errName, err
		}
	}
	if payload.ProposedActivity != nil {
		errName, err := validateProposeActivityList(*payload.ProposedActivity)
		if err != nil {
			return errName, err
		}
	}
	if payload.PlanNote != nil {
		errName, err := validatePlanNoteList(*payload.PlanNote)
		if err != nil {
			return errName, err
		}
	}
	if payload.AdminNote != nil && *payload.AdminNote == "" {
		return "adminNote", AdminNoteRequiredError{}
	}
	return "", nil
}

func validateAdminAssessmentScoreList(assessmentScoreList []map[string]int) (string, error) {
	if len(assessmentScoreList) != allPlanCount {
		return "assessmentScore", SomePlanScoreIsMissing{}
	}
	for i, s := range assessmentScoreList {
		errName, err := validateScore(s)
		if err != nil {
			return fmt.Sprintf("plan.id: %d, %s", i+1, errName), err
		}
	}
	return "", nil
}

func validateProposeActivityList(proposedActivities []string) (string, error) {
	if len(proposedActivities) != allPlanCount {
		return "proposedActivity", SomeProposedActivityIsMissing{}
	}
	for i, p := range proposedActivities {
		if p == "" {
			return fmt.Sprintf("plan.id: %d, proposedActivity", i+1), ProposedActivityRequiredError{}
		}
	}
	return "", nil
}

func validatePlanNoteList(planNotes []string) (string, error) {
	if len(planNotes) != allPlanCount {
		return "planNote", SomePlanNoteIsMissing{}
	}
	for i, p := range planNotes {
		if p == "" {
			return fmt.Sprintf("plan.id: %d, planNote", i+1), PlanNoteRequiredError{}
		}
	}
	return "", nil
}
