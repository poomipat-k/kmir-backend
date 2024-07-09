package plan

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/poomipat-k/kmir-backend/pkg/utils"
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
		&pd.TopicShort,
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

func (s *store) EditPlan(planName string, payload EditPlanRequest, userRole string, username string, userId int) (string, error) {
	// start transaction
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return "transaction", err
	}
	defer tx.Rollback()

	currentPlanData, err := s.GetPlanDetails(payload.PlanName, "user", username)
	if err != nil {
		slog.Error(err.Error())
		return "load_current_plan", err
	}

	loc, err := utils.GetTimeLocation()
	if err != nil {
		return "location", err
	}
	now := time.Now().In(loc)
	// Check if each params exist
	var sqlParams []string
	sqlValues := []any{}
	totalParamsCount := 0
	if payload.ReadinessWillingness != nil && *payload.ReadinessWillingness != *currentPlanData.ReadinessWillingness {
		sqlParams = append(sqlParams, "readiness_willingness", "readiness_willingness_updated_at", "readiness_willingness_updated_by")
		sqlValues = append(sqlValues, payload.ReadinessWillingness, now, userRole)
		totalParamsCount += 3
	}

	if payload.IrGoalType != nil && *payload.IrGoalType != *currentPlanData.IrGoalType {
		sqlParams = append(sqlParams, "ir_goal_type", "ir_goal_type_updated_at", "ir_goal_type_updated_by")
		sqlValues = append(sqlValues, payload.IrGoalType, now, userRole)
		totalParamsCount += 3
	}
	if payload.IrGoalDetails != nil && *payload.IrGoalDetails != *currentPlanData.IrGoalDetails {
		sqlParams = append(sqlParams, "ir_goal_details", "ir_goal_details_updated_at", "ir_goal_details_updated_by")
		sqlValues = append(sqlValues, payload.IrGoalDetails, now, userRole)
		totalParamsCount += 3
	}
	if payload.ProposedActivity != nil && *payload.ProposedActivity != *currentPlanData.ProposedActivity {
		sqlParams = append(sqlParams, "proposed_activity", "proposed_activity_updated_at", "proposed_activity_updated_by")
		sqlValues = append(sqlValues, payload.ProposedActivity, now, userRole)
		totalParamsCount += 3
	}
	if payload.PlanNote != nil && *payload.PlanNote != *currentPlanData.PlanNote {
		sqlParams = append(sqlParams, "plan_note", "plan_note_updated_at", "plan_note_updated_by")
		sqlValues = append(sqlValues, payload.PlanNote, now, userRole)
		totalParamsCount += 3
	}
	if payload.ContactPerson != nil && *payload.ContactPerson != *currentPlanData.ContactPerson {
		sqlParams = append(sqlParams, "contact_person", "contact_person_updated_at", "contact_person_updated_by")
		sqlValues = append(sqlValues, payload.ContactPerson, now, userRole)
		totalParamsCount += 3
	}
	scoreChanged := false
	if payload.AssessmentScore != nil {
		// add all 7 rows to  assessment_score table
		var addScoreBuilder strings.Builder
		scoreValues := []any{}
		addScoreBuilder.WriteString("INSERT INTO assessment_score (plan_id, user_id, assessment_criteria_id, score, year, created_at) VALUES ")
		// check if there is at least 1 score changed
		for i := 0; i < 7; i++ {
			addScoreBuilder.WriteString(fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d)",
				i*6+1,
				i*6+2,
				i*6+3,
				i*6+4,
				i*6+5,
				i*6+6,
			))
			newScore := payload.AssessmentScore[fmt.Sprintf("q_%d", currentPlanData.AssessmentCriteria[i].CriteriaId)]
			oldScore := getOldScore(currentPlanData, userRole, now.Year(), i+1)
			scoreChanged = scoreChanged || (newScore != oldScore)
			scoreValues = append(scoreValues,
				currentPlanData.PlanId,
				userId,
				currentPlanData.AssessmentCriteria[i].CriteriaId,
				newScore,
				now.Year(),
				now,
			)
			if i < 6 {
				addScoreBuilder.WriteString(", ")
			}
		}

		addScoreBuilder.WriteString(";")
		stmt, err := tx.Prepare(addScoreBuilder.String())
		if err != nil {
			slog.Error("error prepare add insert assessment_score sql", "error", err)
			return "prepare_sql_score", err
		}
		result, err := stmt.ExecContext(ctx, scoreValues...)
		if err != nil {
			slog.Error("execContext on assessment_score sql", "error", err)
			return "exec_sql_score", err
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			slog.Error("execContext on insert assessment_score sql", "error", err)
			return "rows_affected_score", err
		}
		if rowsAffected == 0 {
			return "zero_row_affected_score", errors.New("no plan is updated")
		}
	}
	if totalParamsCount == 0 && !scoreChanged {
		return "no_changes", errors.New("updated plan failed: no new values detected")
	}

	if totalParamsCount > 0 {
		var updateSQLBuilder strings.Builder
		updateSQLBuilder.WriteString("UPDATE plan SET ")
		n := len(sqlParams)
		for i := 0; i < n; i++ {
			updateSQLBuilder.WriteString(sqlParams[i])
			updateSQLBuilder.WriteString(fmt.Sprintf(" = $%d, ", i+1))
		}
		updateSQLBuilder.WriteString(fmt.Sprintf("updated_at = $%d", n+1))
		updateSQLBuilder.WriteString(fmt.Sprintf(" WHERE plan.name = $%d;", n+2))
		updateSQL := updateSQLBuilder.String()

		stmt, err := tx.Prepare(updateSQL)
		if err != nil {
			slog.Error("error prepare add update plan sql", "error", err)
			return "prepare_sql_plan", err
		}
		sqlValues = append(sqlValues, now, planName)
		result, err := stmt.ExecContext(ctx, sqlValues...)
		if err != nil {
			slog.Error("execContext on update plan sql", "error", err)
			return "exec_sql_plan", err
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			slog.Error("execContext on update plan sql", "error", err)
			return "rows_affected_plan", err
		}
		if rowsAffected == 0 {
			return "zero_row_affected_plan", errors.New("no plan is updated")
		}
	}

	err = tx.Commit()
	if err != nil {
		slog.Error("commit error", "err", err.Error())
		return "commit", err
	}
	return "", nil
}

func getOldScore(currentPlanData PlanDetails, userRole string, year int, targetCriteriaOrder int) int {
	oldScore := 0
	if currentPlanData.AssessmentScore != nil {
		for j := 0; j < len(currentPlanData.AssessmentScore); j++ {
			scoreItem := currentPlanData.AssessmentScore[j]
			if scoreItem.PlanId == currentPlanData.PlanId &&
				scoreItem.CriteriaOrder == targetCriteriaOrder &&
				scoreItem.UserRole == userRole &&
				scoreItem.Year == year {
				oldScore = scoreItem.Score
			}
		}
	}
	return oldScore
}
