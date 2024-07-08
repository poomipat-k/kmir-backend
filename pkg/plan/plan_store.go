package plan

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"strings"
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

func (s *store) EditPlan(planName string, payload EditPlanRequest, userRole string) error {
	// start transaction
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	now := time.Now()
	log.Println("==now", now)
	// Check if each params exist
	var sqlParams []string
	sqlValues := []any{}
	totalParamsCount := 0
	if payload.ReadinessWillingness != nil {
		sqlParams = append(sqlParams, "readiness_willingness", "readiness_willingness_updated_at", "readiness_willingness_updated_by")
		sqlValues = append(sqlValues, payload.ReadinessWillingness, now, userRole)
		totalParamsCount += 3
	}
	if payload.IrGoalType != nil {
		sqlParams = append(sqlParams, "ir_goal_type", "ir_goal_type_updated_at", "ir_goal_type_updated_by")
		sqlValues = append(sqlValues, payload.IrGoalType, now, userRole)
		totalParamsCount += 3
	}
	if payload.IrGoalDetails != nil {
		sqlParams = append(sqlParams, "ir_goal_details", "ir_goal_details_updated_at", "ir_goal_details_updated_by")
		sqlValues = append(sqlValues, payload.IrGoalDetails, now, userRole)
		totalParamsCount += 3
	}
	if payload.ProposedActivity != nil {
		sqlParams = append(sqlParams, "proposed_activity_details", "proposed_activity_details_updated_at", "proposed_activity_details_updated_by")
		sqlValues = append(sqlValues, payload.ProposedActivity, now, userRole)
		totalParamsCount += 3
	}
	if payload.PlanNote != nil {
		sqlParams = append(sqlParams, "plan_note_details", "plan_note_details_updated_at", "plan_note_details_updated_by")
		sqlValues = append(sqlValues, payload.PlanNote, now, userRole)
		totalParamsCount += 3
	}
	if payload.ContactPerson != nil {
		sqlParams = append(sqlParams, "contact_person_details", "contact_person_details_updated_at", "contact_person_details_updated_by")
		sqlValues = append(sqlValues, payload.ContactPerson, now, userRole)
		totalParamsCount += 3
	}
	var updateSQLBuilder strings.Builder
	updateSQLBuilder.WriteString("UPDATE plan SET ")
	n := len(sqlParams)
	for i := 0; i < n; i++ {
		updateSQLBuilder.WriteString(sqlParams[i])
		updateSQLBuilder.WriteString(fmt.Sprintf(" = $%d", i+1))
		if i < n-1 {
			updateSQLBuilder.WriteString(", ")
		}
	}
	updateSQLBuilder.WriteString(fmt.Sprintf(" WHERE plan.name = $%d;", n+1))
	updateSQL := updateSQLBuilder.String()

	log.Println("===updateSQL:", updateSQL)

	stmt, err := tx.Prepare(updateSQL)
	if err != nil {
		slog.Error("error prepare add update plan sql", "error", err)
		return err
	}
	sqlValues = append(sqlValues, planName)
	result, err := stmt.ExecContext(ctx, sqlValues...)
	if err != nil {
		slog.Error("execContext on update plan sql", "error", err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.Error("execContext on update plan sql", "error", err)
		return err
	}
	log.Println("==rowsAffected", rowsAffected)
	if rowsAffected == 0 {
		return errors.New("no plan is updated")
	}

	err = tx.Commit()
	if err != nil {
		slog.Error("commit error", "err", err.Error())
		return err
	}
	return nil
}
