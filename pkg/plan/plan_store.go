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
		err = rows.Scan(&row.Id, &row.Name, &row.UserId, &row.Topic, &row.TopicEn, &row.IrGoalDetails)
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

func (s *store) GetAllPlanDetails(criteriaLen int) ([]AdminDashboardPlanDetailsRow, error) {
	planData, err := s.getAllPlansDetailsForAdmin()
	if err != nil {
		return nil, err
	}

	// get score details
	scoreRowsData, err := s.getAllPlanScoresForAdmin(criteriaLen)
	if err != nil {
		return nil, err
	}

	// put score to each plan details
	for _, scoreRow := range scoreRowsData {
		for index, plan := range planData {
			if plan.PlanId == scoreRow.PlanId {
				planData[index].AssessmentScore = append(planData[index].AssessmentScore, scoreRow)
				break
			}
		}
	}

	return planData, nil
}

func (s *store) getAllPlansDetailsForAdmin() ([]AdminDashboardPlanDetailsRow, error) {
	rows, err := s.db.Query(getAllPlanDetailsForAdminDashboardSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var planData []AdminDashboardPlanDetailsRow
	for rows.Next() {
		var row AdminDashboardPlanDetailsRow
		err = rows.Scan(
			&row.PlanId,
			&row.Name,
			&row.Topic,
			&row.TopicEn,
			&row.TopicShort,
			&row.IrGoalType,
			&row.IrGoalDetails,
			&row.ProposedActivity,
			&row.PlanNote,
			&row.UpdatedAt,
			&row.UpdatedBy,
			&row.IrGoalTypeUpdatedAt,
			&row.IrGoalTypeUpdatedBy,
			&row.IrGoalDetailsUpdatedAt,
			&row.IrGoalDetailsUpdatedBy,
			&row.ProposedActivityUpdatedAt,
			&row.ProposedActivityUpdatedBy,
			&row.PlanNoteUpdatedAt,
			&row.PlanNoteUpdatedBy,
			&row.ContactPersonUpdatedAt,
			&row.ContactPersonUpdatedBy,
		)
		if err != nil {
			slog.Error(err.Error(), "field", "scan AdminDashboardPlanDetailsRow")
			return nil, err
		}
		planData = append(planData, row)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return planData, nil
}

func (s *store) getAllPlanScoresForAdmin(criteriaLen int) ([]AssessmentScoreRow, error) {
	loc, err := utils.GetTimeLocation()
	if err != nil {
		return nil, err
	}
	now := time.Now().In(loc)
	fromDate := time.Date(now.Year(), time.Month(1), 1, 0, 0, 0, 0, loc)
	toDate := time.Date(now.Year()+1, time.Month(1), 1, 0, 0, 0, 0, loc)

	scoreRows, err := s.db.Query(adminGetAllPlanScoreDetailsSQL, criteriaLen, fromDate, toDate)
	if err != nil {
		return nil, err
	}
	defer scoreRows.Close()

	var scoreRowsData []AssessmentScoreRow
	for scoreRows.Next() {
		var row AssessmentScoreRow
		err = scoreRows.Scan(&row.PlanId, &row.CriteriaId, &row.CriteriaOrder, &row.Score, &row.CreatedAt)
		if err != nil {
			slog.Error(err.Error(), "field", "scoreRowsData scan: AssessmentScoreRow")
			return nil, err
		}
		scoreRowsData = append(scoreRowsData, row)
	}
	err = scoreRows.Err()
	if err != nil {
		return nil, err
	}
	return scoreRowsData, nil
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

	// Get all assessment_criteria if not already exists
	if pd.AssessmentCriteria == nil {
		assessmentRows, err := s.db.Query(getAllAssessmentCriteriaSQL)
		if err != nil {
			return PlanDetails{}, err
		}
		defer assessmentRows.Close()

		var assessmentCriteriaData []AssessmentCriteria
		for assessmentRows.Next() {
			var assessmentCri AssessmentCriteria
			err = assessmentRows.Scan(&assessmentCri.Category, &assessmentCri.CriteriaId, &assessmentCri.Display, &assessmentCri.OrderNumber)
			if err != nil {
				slog.Error(err.Error(), "field", "scan assessmentCriteriaRows")
				return PlanDetails{}, err
			}
			assessmentCriteriaData = append(assessmentCriteriaData, assessmentCri)
		}
		err = assessmentRows.Err()
		if err != nil {
			return PlanDetails{}, err
		}
		pd.AssessmentCriteria = assessmentCriteriaData
	}

	return pd, nil
}

func (s *store) GetAssessmentCriteria() ([]AssessmentCriteria, error) {
	assessmentRows, err := s.db.Query(getAllAssessmentCriteriaSQL)
	if err != nil {
		return nil, err
	}
	defer assessmentRows.Close()

	var data []AssessmentCriteria
	for assessmentRows.Next() {
		var assessmentCri AssessmentCriteria
		err = assessmentRows.Scan(&assessmentCri.Category, &assessmentCri.CriteriaId, &assessmentCri.Display, &assessmentCri.OrderNumber)
		if err != nil {
			slog.Error(err.Error(), "field", "GetAssessmentCriteria(): scan assessmentCriteriaRows")
			return nil, err
		}
		data = append(data, assessmentCri)
	}
	err = assessmentRows.Err()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *store) AdminGetScores(fromYear, toYear int, plan string) ([]AssessmentScore, error) {
	sqlStmt := prepareAdminGetScoresSQL(plan)
	sqlValues, err := prepareAdminGetScoresSQLValues(fromYear, toYear, plan)
	if err != nil {
		return nil, err
	}
	stmt, err := s.db.Prepare(sqlStmt)
	if err != nil {
		slog.Error("error prepare AdminGetScores sql", "error", err)
		return nil, err
	}
	rows, err := stmt.Query(sqlValues...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rowData []AssessmentScoreRow
	for rows.Next() {
		var row AssessmentScoreRow
		err = rows.Scan(&row.PlanId, &row.CriteriaId, &row.CriteriaOrder, &row.UserRole, &row.Year, &row.Score, &row.CriteriaCategory, &row.CreatedAt)
		if err != nil {
			slog.Error(err.Error(), "field", "scan AssessmentScoreRow")
			return nil, err
		}
		rowData = append(rowData, row)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	// only send useful params
	var scores []AssessmentScore
	for _, row := range rowData {
		scores = append(scores, AssessmentScore{
			PlanId:           row.PlanId,
			CriteriaOrder:    row.CriteriaOrder,
			CriteriaCategory: row.CriteriaCategory,
			UserRole:         row.UserRole,
			Year:             row.Year,
			Score:            row.Score,
			CreatedAt:        row.CreatedAt,
		})
	}
	return scores, nil
}

func (s *store) GetAdminNote() (string, error) {
	var adminNote string
	row := s.db.QueryRow(getAdminNote)
	err := row.Scan(&adminNote)

	if err == sql.ErrNoRows {
		slog.Error("GetAdminNote(): no row were returned!")
		return "", err
	}
	if err != nil {
		slog.Error(err.Error())
		return "", fmt.Errorf("GetAdminNote() unknown error")
	}
	return adminNote, nil
}

func (s *store) GetOnlyLatestScore() ([]LatestScoreTimestamp, error) {
	rows, err := s.db.Query(getOnlyLatestScoreSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []LatestScoreTimestamp
	for rows.Next() {
		var row LatestScoreTimestamp
		err = rows.Scan(&row.PlanId, &row.UserRole, &row.CreatedAt)
		if err != nil {
			slog.Error(err.Error(), "field", "GetOnlyLatestScore(): scan LatestScoreTimestamp")
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

func prepareAdminGetScoresSQL(plan string) string {
	var sqlStmt string
	if plan != "all" {
		sqlStmt = `
		SELECT
		plan_id,
		criteria_id,
		criteria_order, 
		user_role,
		year,
		score,
		criteria_category,
		created_at
		FROM
		(
		SELECT
		assessment_score.plan_id as plan_id,
		assessment_score.assessment_criteria_id as criteria_id,
		assessment_criteria.order_number as criteria_order,
		assessment_score.user_id as user_id,
		users.user_role as user_role,
		year, 
		score, 
		assessment_criteria.category as criteria_category,
		assessment_score.created_at as created_at,
		ROW_NUMBER() OVER (
		PARTITION BY assessment_score.plan_id, assessment_score.user_id, year 
		ORDER BY assessment_score.created_at DESC, assessment_criteria_id ASC) 
		as row_num FROM assessment_score 
		INNER JOIN plan ON plan.id = assessment_score.plan_id
		INNER JOIN assessment_criteria ON assessment_criteria.id = assessment_score.assessment_criteria_id
		INNER JOIN users ON users.id = assessment_score.user_id
		WHERE plan.name = $1
		)
		WHERE row_num <= 7 AND created_at >= $2 AND created_at < $3
		;
		`
	} else {
		sqlStmt = `
		SELECT
		plan_id,
		criteria_id,
		criteria_order, 
		user_role,
		year,
		score,
		criteria_category,
		created_at
		FROM
		(
		SELECT
		assessment_score.plan_id as plan_id,
		assessment_score.assessment_criteria_id as criteria_id,
		assessment_criteria.order_number as criteria_order,
		assessment_score.user_id as user_id,
		users.user_role as user_role,
		year, 
		score, 
		assessment_criteria.category as criteria_category,
		assessment_score.created_at as created_at,
		ROW_NUMBER() OVER (
		PARTITION BY assessment_score.plan_id, assessment_score.user_id, year 
		ORDER BY assessment_score.created_at DESC, assessment_criteria_id ASC) 
		as row_num FROM assessment_score 
		INNER JOIN plan ON plan.id = assessment_score.plan_id
		INNER JOIN assessment_criteria ON assessment_criteria.id = assessment_score.assessment_criteria_id
		INNER JOIN users ON users.id = assessment_score.user_id
		)
		WHERE row_num <= 7 AND created_at >= $1 AND created_at < $2
		;
		`
	}
	return sqlStmt
}

func prepareAdminGetScoresSQLValues(fromYear, toYear int, plan string) ([]any, error) {
	sqlValues := []any{}
	if plan != "all" {
		sqlValues = append(sqlValues, plan)
	}
	loc, err := utils.GetTimeLocation()
	if err != nil {
		return nil, err
	}
	fromDate := time.Date(fromYear, time.Month(1), 1, 0, 0, 0, 0, loc)
	toDate := time.Date(toYear+1, time.Month(1), 1, 0, 0, 0, 0, loc)
	sqlValues = append(sqlValues, fromDate)
	sqlValues = append(sqlValues, toDate)
	return sqlValues, nil
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
		updateSQLBuilder.WriteString(fmt.Sprintf("updated_at = $%d, updated_by = 'user'", n+1))
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

func (s *store) AdminEdit(payload AdminEditRequest, userId int) (bool, string, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return false, "transaction", err
	}
	defer tx.Rollback()

	criteriaList, err := s.GetAssessmentCriteria()
	if err != nil {
		slog.Error(err.Error())
		return false, "criteriaList", err
	}
	criteriaLen := len(criteriaList)
	planDetails, err := s.GetAllPlanDetails(criteriaLen)
	if err != nil {
		slog.Error(err.Error())
		return false, "Admin update: GetAllPlanDetails", err
	}

	now, err := utils.GetNow()
	if err != nil {
		return false, "now", err
	}

	hasUpdated := false

	if payload.AssessmentScore != nil {
		scoreChanged, errName, err := handleAdminUpdateAssessmentScore(ctx, tx, planDetails, userId, criteriaLen, now, *payload.AssessmentScore)
		if err != nil {
			return false, errName, err
		}
		hasUpdated = hasUpdated || scoreChanged
	}

	planUpdated, errName, err := handleAdminUpdatePlan(ctx, tx, planDetails, now, payload)
	if err != nil {
		return false, errName, err
	}
	hasUpdated = hasUpdated || planUpdated

	if payload.AdminNote != nil {
		adminNoteUpdated, errName, err := s.handleAdminUpdateAdminNote(ctx, tx, *payload.AdminNote)
		if err != nil {
			return false, errName, err
		}
		hasUpdated = hasUpdated || adminNoteUpdated
	}

	err = tx.Commit()
	if err != nil {
		slog.Error("commit error", "err", err.Error())
		return false, "commit", err
	}
	return hasUpdated, "", nil
}

func (s *store) handleAdminUpdateAdminNote(
	ctx context.Context,
	tx *sql.Tx,
	newNote string,
) (bool, string, error) {
	curAdminNote, err := s.GetAdminNote()
	if err != nil {
		return false, "adminUpdate: GetAdminNote", err
	}
	if curAdminNote != newNote {
		const sql = "UPDATE admin_note SET note = $1 WHERE id = 1;"
		result, err := tx.ExecContext(ctx, sql, newNote)
		if err != nil {
			slog.Error("admin update admin_note sql", "error", err)
			return false, "update_admin_note", err
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			slog.Error("rowsAffected update admin_note sql", "error", err)
			return false, "rows_affected_admin_note", err
		}
		if rowsAffected == 0 {
			return false, "zero_row_affected_admin_note", errors.New("update admin_note failed")
		}
		return true, "", nil
	}
	return false, "", nil
}

func handleAdminUpdatePlan(
	ctx context.Context,
	tx *sql.Tx,
	curPlanDetails []AdminDashboardPlanDetailsRow,
	now time.Time,
	payload AdminEditRequest,
) (bool, string, error) {
	changes := getPlanChangesData(payload, curPlanDetails)
	updated := false
	log.Println("===changes", changes)
	for i, updateList := range changes {
		if updateList[0] || updateList[1] || updateList[2] || updateList[3] {
			values := []any{}
			paramCounter := 1
			var updatePlanBuilder strings.Builder
			updatePlanBuilder.WriteString("UPDATE plan SET ")
			needComma := false
			if updateList[0] {
				updatePlanBuilder.WriteString(fmt.Sprintf("proposed_activity = $%d, proposed_activity_updated_at = $%d, proposed_activity_updated_by = 'admin'", paramCounter, paramCounter+1))
				values = append(values, (*payload.ProposedActivity)[i], now)
				paramCounter += 2
				needComma = true
			}
			if updateList[1] {
				if needComma {
					updatePlanBuilder.WriteString(", ")
				}
				needComma = true
				updatePlanBuilder.WriteString(fmt.Sprintf("plan_note = $%d, plan_note_updated_at = $%d, plan_note_updated_by = 'admin'", paramCounter, paramCounter+1))
				values = append(values, (*payload.PlanNote)[i], now)
				paramCounter += 2
			}
			if updateList[2] {
				if needComma {
					updatePlanBuilder.WriteString(", ")
				}
				needComma = true
				updatePlanBuilder.WriteString(fmt.Sprintf("ir_goal_type = $%d, ir_goal_type_updated_at = $%d, ir_goal_type_updated_by = 'admin'", paramCounter, paramCounter+1))
				values = append(values, *(*payload.IrWorkGoal)[i].GoalType, now)
				paramCounter += 2
			}
			if updateList[3] {
				if needComma {
					updatePlanBuilder.WriteString(", ")
				}
				updatePlanBuilder.WriteString(fmt.Sprintf("ir_goal_details = $%d, ir_goal_details_updated_at = $%d, ir_goal_details_updated_by = 'admin'", paramCounter, paramCounter+1))
				values = append(values, *(*payload.IrWorkGoal)[i].GoalDetails, now)
				paramCounter += 2
				needComma = true
			}

			updatePlanBuilder.WriteString(fmt.Sprintf(", updated_at = $%d, updated_by = 'admin' WHERE plan.id = $%d;", paramCounter, paramCounter+1))
			values = append(values, now, curPlanDetails[i].PlanId)

			stmt, err := tx.Prepare(updatePlanBuilder.String())
			if err != nil {
				slog.Error("error prepare admin update plan sql", "error", err)
				return false, "prepare_sql_update_plan", err
			}
			result, err := stmt.ExecContext(ctx, values...)
			if err != nil {
				slog.Error("execContext on admin update plan sql", "error", err)
				return false, "exec_admin_update_plan_sql", err
			}
			rowsAffected, err := result.RowsAffected()
			if err != nil {
				slog.Error("rowsAffected on admin update plan sql", "error", err)
				return false, "rows_affected_admin_update_plan", err
			}
			if rowsAffected == 0 {
				return false, "zero_row_affected_admin_update_plan", errors.New("no plan is updated")
			}
			updated = true
		}
	}
	return updated, "", nil
}

func getPlanChangesData(payload AdminEditRequest, curPlanDetails []AdminDashboardPlanDetailsRow) [][4]bool {
	changes := [][4]bool{} // [][proposedActivity, planNote, irGoalType, irGoalDetails]
	for i, p := range curPlanDetails {
		changes = append(changes, [4]bool{false, false, false, false})
		if payload.ProposedActivity != nil && (*payload.ProposedActivity)[i] != *p.ProposedActivity {
			changes[i][0] = true
		}
		if payload.PlanNote != nil && (*payload.PlanNote)[i] != *p.PlanNote {
			changes[i][1] = true
		}
		if payload.IrWorkGoal != nil {
			if *(*payload.IrWorkGoal)[i].GoalType != *p.IrGoalType {
				log.Println("here type i: ", i)
				changes[i][2] = true
			}
			if (*(*payload.IrWorkGoal)[i].GoalDetails) != *p.IrGoalDetails {
				log.Println("here goal details i: ", i)
				changes[i][3] = true
			}
		}
	}
	return changes
}

func handleAdminUpdateAssessmentScore(
	ctx context.Context,
	tx *sql.Tx,
	curPlanDetails []AdminDashboardPlanDetailsRow,
	userId int,
	criteriaLen int,
	now time.Time,
	newScores []map[string]int,
) (bool, string, error) {
	changesIndex := getScoreChangedPlanIndex(newScores, curPlanDetails, criteriaLen)

	updated := false
	// insert new scores for each plan in changesIndex
	for _, target := range changesIndex {
		// add all 7 rows to  assessment_score table
		var addScoreBuilder strings.Builder
		scoreValues := []any{}
		addScoreBuilder.WriteString("INSERT INTO assessment_score (plan_id, user_id, assessment_criteria_id, score, year, created_at) VALUES ")
		for i := 0; i < criteriaLen; i++ {
			addScoreBuilder.WriteString(fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d)",
				i*6+1,
				i*6+2,
				i*6+3,
				i*6+4,
				i*6+5,
				i*6+6,
			))
			if i < criteriaLen-1 {
				addScoreBuilder.WriteString(", ")
			}
			score := newScores[target][fmt.Sprintf("q_%d", i+1)]
			scoreValues = append(scoreValues,
				curPlanDetails[target].PlanId,
				userId,
				i+1,
				score,
				now.Year(),
				now,
			)
		}

		addScoreBuilder.WriteString(";")
		stmt, err := tx.Prepare(addScoreBuilder.String())
		if err != nil {
			slog.Error("error admin prepare add insert assessment_score sql", "error", err)
			return false, "admin_prepare_sql_score", err
		}
		result, err := stmt.ExecContext(ctx, scoreValues...)
		if err != nil {
			slog.Error("error admin execContext on assessment_score sql", "error", err)
			return false, "admin_exec_sql_score", err
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			slog.Error("admin execContext on insert assessment_score sql", "error", err)
			return false, "admin_rows_affected_score", err
		}
		if rowsAffected == 0 {
			return false, "admin_zero_row_affected_score", errors.New("admin: no score is created")
		}
		updated = true
	}

	return updated, "", nil
}

func getScoreChangedPlanIndex(newScores []map[string]int, curPlanDetails []AdminDashboardPlanDetailsRow, criteriaLen int) []int {
	var changesIndex []int
	for i, s := range newScores {
		changed := hasScoreChanged(s, curPlanDetails[i].AssessmentScore, criteriaLen)
		if changed {
			changesIndex = append(changesIndex, i)
		}
	}
	return changesIndex
}

func hasScoreChanged(newScores map[string]int, oldScores []AssessmentScoreRow, criteriaLen int) bool {
	if len(oldScores) == 0 {
		return true
	}
	for i := 0; i < criteriaLen; i++ {
		newVal := newScores[fmt.Sprintf("q_%d", i+1)]
		if oldScores[i].Score != newVal {
			return true
		}
	}
	return false
}
