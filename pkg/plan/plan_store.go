package plan

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"time"
)

type store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *store {
	return &store{
		db: db,
	}
}

func (s *store) GetAllPreviewPlan() ([]PlanPreview, error) {
	rows, err := s.db.Query(getAllPreviewPlanSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []PlanPreview
	for rows.Next() {
		var row PlanPreview
		err = rows.Scan(&row.Id, &row.Name, &row.UserId, &row.Topic, &row.TopicEn, &row.Goal)
		if err != nil {
			return nil, err
		}

		data = append(data, row)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *store) CanAccessPlanDetails(planName, username string) (bool, error) {
	var planId int
	row := s.db.QueryRow(canAccessPlanDetailsSQL, planName, username)
	err := row.Scan(&planId)

	if err == sql.ErrNoRows {
		slog.Error("CanAccessPlanDetails(): no row were returned!")
		return false, err
	}
	if err != nil {
		slog.Error(err.Error())
		return false, fmt.Errorf("CanAccessPlanDetails() unknown error")
	}
	return true, nil
}

func (s *store) CanEditPlan(planName, username string) (bool, error) {
	var planId int
	row := s.db.QueryRow(canEditPlanSQL, planName, username)
	err := row.Scan(&planId)

	if err == sql.ErrNoRows {
		slog.Error("CanEditPlan(): no row were returned!")
		return false, err
	}
	if err != nil {
		slog.Error(err.Error())
		return false, fmt.Errorf("CanEditPlan() unknown error")
	}
	return true, nil
}

func (s *store) GetPlanDetails(planName, userRole string, username string) (PlanDetails, error) {
	var pd PlanDetails
	var planRow *sql.Row
	if userRole == "admin" || userRole == "viewer" {
		planRow = s.db.QueryRow(getPlanDetailsForAdminViewSQL, planName)
	} else {
		planRow = s.db.QueryRow(getPlanDetailsSQL, planName, username)
	}

	err := planRow.Scan(
		&pd.PlanId,
		&pd.Name,
		&pd.Topic,
		&pd.TopicEn,
		&pd.ReadinessWillingness,
		&pd.ReadinessWillingnessUpdatedAt,
		&pd.ReadinessWillingnessUpdatedBy,
		&pd.IrGoalType,
		&pd.IrGoalTypeUpdatedAt,
		&pd.IrGoalTypeUpdatedBy,
		&pd.IrGoalDetails,
		&pd.IrGoalDetailsUpdatedAt,
		&pd.IrGoalDetailsUpdatedBy,
		&pd.ProposedActivity,
		&pd.ProposedActivityUpdatedAt,
		&pd.ProposedActivityUpdatedBy,
		&pd.PlanNote,
		&pd.PlanNoteUpdatedAt,
		&pd.PlanNoteUpdatedBy,
		&pd.ContactPerson,
		&pd.ContactPersonUpdatedAt,
		&pd.ContactPersonUpdatedBy,
		&pd.UpdatedAt,
		&pd.UpdatedBy,
	)
	if err == sql.ErrNoRows {
		slog.Error("GetPlanDetails() general: no row were returned!")
		return PlanDetails{}, err
	}
	if err != nil {
		slog.Error(err.Error())
		return PlanDetails{}, fmt.Errorf("GetPlanDetails() general: unknown error")
	}

	// If reach here it means the user match the plan so no need to check username in the query
	rows, err := s.db.Query(getPlanScoreDetailsSQL, planName)
	if err != nil {
		return PlanDetails{}, err
	}
	defer rows.Close()

	var rowsData []AssessmentScoreRow
	for rows.Next() {
		var row AssessmentScoreRow
		err = rows.Scan(&row.PlanId, &row.CriteriaId, &row.CriteriaOrder, &row.UserRole, &row.Year, &row.Score, &row.CreatedAt, &row.CriteriaCategory, &row.CriteriaDisplay)
		if err != nil {
			slog.Error(err.Error(), "field", "scan AssessmentScoreRow")
			return PlanDetails{}, err
		}
		rowsData = append(rowsData, row)
	}
	err = rows.Err()
	if err != nil {
		return PlanDetails{}, err
	}

	// Add AssessmentCriteria to pd
	var assessmentCriteriaList []AssessmentCriteria
	for index, row := range rowsData {
		if index >= 7 {
			break
		}
		assessmentCriteriaList = append(assessmentCriteriaList, AssessmentCriteria{
			CriteriaId:  row.CriteriaId,
			OrderNumber: row.CriteriaOrder,
			Category:    row.CriteriaCategory,
			Display:     row.CriteriaDisplay,
		})
	}
	pd.AssessmentCriteria = assessmentCriteriaList

	// Add score details by year
	var scores []AssessmentScore
	for _, row := range rowsData {
		scores = append(scores, AssessmentScore{
			PlanId:        row.PlanId,
			CriteriaOrder: row.CriteriaOrder,
			UserRole:      row.UserRole,
			Year:          row.Year,
			Score:         row.Score,
			CreatedAt:     row.CreatedAt,
		})
	}
	pd.AssessmentScore = scores

	return pd, nil
}

func (s *store) EditPlan(planName, username string) (bool, error) {
	now := time.Now()
	log.Println("==now", now)
	// var planId int
	// row := s.db.QueryRow(canEditPlanSQL, planName, username)
	// err := row.Scan(&planId)

	// if err == sql.ErrNoRows {
	// 	slog.Error("CanEditPlan(): no row were returned!")
	// 	return false, err
	// }
	// if err != nil {
	// 	slog.Error(err.Error())
	// 	return false, fmt.Errorf("CanEditPlan() unknown error")
	// }
	return true, nil
}
