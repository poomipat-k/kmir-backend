package plan

import (
	"net/http"

	"github.com/poomipat-k/kmir-backend/pkg/utils"
)

type PlanStore interface {
	GetAllPreviewPlan() ([]PlanPreview, error)
}

type PlanHandler struct {
	store PlanStore
}

func NewPlanHandler(s PlanStore) *PlanHandler {
	return &PlanHandler{
		store: s,
	}
}

func (h *PlanHandler) GetAllPreviewPlan(w http.ResponseWriter, r *http.Request) {
	plans, err := h.store.GetAllPreviewPlan()
	if err != nil {
		utils.ErrorJSON(w, err, "getAllPreviewPlan", http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, plans)
}
