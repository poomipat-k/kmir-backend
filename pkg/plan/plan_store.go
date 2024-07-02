package plan

import (
	"database/sql"
	"fmt"
	"log/slog"
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

func (s *store) GetPlanDetails(planName string) (PlanDetails, error) {
	var pd PlanDetails
	planRow := s.db.QueryRow(getPlanDetailsSQL, planName)
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
		slog.Error("GetPlanDetails(): no row were returned!")
		return PlanDetails{}, err
	}
	if err != nil {
		slog.Error(err.Error())
		return PlanDetails{}, fmt.Errorf("GetPlanDetails() unknown error")
	}
	return pd, nil
}
