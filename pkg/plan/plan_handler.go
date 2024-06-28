package plan

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/poomipat-k/kmir-backend/pkg/common"
	"github.com/poomipat-k/kmir-backend/pkg/utils"
)

type PlanStore interface {
	GetAllPreviewPlan() ([]PlanPreview, error)
	CanAccessPlanDetails(planName, username string) (bool, error)
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

func (h *PlanHandler) CanAccessPlanDetails(w http.ResponseWriter, r *http.Request) {
	username, err := utils.GetUsernameFromRequestHeader(r)
	if err != nil {
		slog.Error(err.Error())
		utils.ErrorJSON(w, err, "username", http.StatusUnauthorized)
		return
	}
	userRole, err := utils.GetUserRoleFromRequestHeader(r)
	if userRole == "" {
		slog.Error(err.Error())
		utils.ErrorJSON(w, err, "userRole", http.StatusUnauthorized)
		return
	}

	planName := chi.URLParam(r, "planName")
	allow, err := h.store.CanAccessPlanDetails(planName, username)
	if err != nil {
		slog.Error(err.Error())
		utils.ErrorJSON(w, err, "", http.StatusNotFound)
		return
	}
	if !allow {
		utils.WriteJSON(w, http.StatusForbidden, common.CommonSuccessResponse{
			Success: false,
			Message: "user permission denied",
		})
		return
	}
	utils.WriteJSON(w, http.StatusOK, common.CommonSuccessResponse{
		Success: true,
		Message: "allow to access plan details",
	})
}
