package plan

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/poomipat-k/kmir-backend/pkg/common"
	"github.com/poomipat-k/kmir-backend/pkg/utils"
)

type PlanStore interface {
	GetAllPreviewPlan() ([]PlanPreview, error)
	CanAccessPlanDetails(planName, username string) (bool, error)
	GetPlanDetails(planName, userRole string, username string) (PlanDetails, error)
	CanEditPlan(planName, username string) (bool, error)
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
	userRole, err := utils.GetUserRoleFromRequestHeader(r)
	if userRole == "" {
		slog.Error(err.Error())
		utils.ErrorJSON(w, err, "userRole", http.StatusUnauthorized)
		return
	}
	plans, err := h.store.GetAllPreviewPlan()
	if err != nil {
		utils.ErrorJSON(w, err, "getAllPreviewPlan", http.StatusInternalServerError)
		return
	}
	if userRole == "user" {
		var filterData []PlanPreview
		for _, p := range plans {
			if strings.ToLower(p.Name) != "admin" {
				filterData = append(filterData, p)
			}
		}
		utils.WriteJSON(w, http.StatusOK, filterData)
		return
	}
	utils.WriteJSON(w, http.StatusOK, plans)
}

func (h *PlanHandler) GetPlanDetails(w http.ResponseWriter, r *http.Request) {
	userRole, err := utils.GetUserRoleFromRequestHeader(r)
	if userRole == "" {
		slog.Error(err.Error())
		utils.ErrorJSON(w, err, "userRole", http.StatusUnauthorized)
		return
	}
	username, err := utils.GetUsernameFromRequestHeader(r)
	if err != nil {
		slog.Error(err.Error())
		utils.ErrorJSON(w, err, "username", http.StatusUnauthorized)
		return
	}
	planName := chi.URLParam(r, "planName")
	data, err := h.store.GetPlanDetails(planName, userRole, username)
	if err != nil {
		slog.Error(err.Error())
		utils.ErrorJSON(w, err, "planName", http.StatusBadRequest)
		return
	}
	utils.WriteJSON(w, http.StatusOK, data)
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
	if userRole == "admin" || userRole == "viewer" {
		utils.WriteJSON(w, http.StatusOK, common.CommonSuccessResponse{
			Success: true,
			Message: "allow admin or viewer to access plan details",
		})
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

func (h *PlanHandler) CanEditPlan(w http.ResponseWriter, r *http.Request) {
	username, err := utils.GetUsernameFromRequestHeader(r)
	if err != nil {
		slog.Error(err.Error())
		utils.ErrorJSON(w, err, "username", http.StatusUnauthorized)
		return
	}
	userRole, err := utils.GetUserRoleFromRequestHeader(r)
	if err != nil {
		slog.Error(err.Error())
		utils.ErrorJSON(w, err, "userRole", http.StatusUnauthorized)
		return
	}
	if userRole != "user" {
		utils.ErrorJSON(w, errors.New("no edit permission"), "userRole", http.StatusForbidden)
		return
	}

	planName := chi.URLParam(r, "planName")
	allow, err := h.store.CanEditPlan(planName, username)
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
		Message: "allow to edit plan details",
	})
}

func (h *PlanHandler) EditPlan(w http.ResponseWriter, r *http.Request) {
	var payload EditPlanRequest
	err := utils.ReadJSON(w, r, &payload)
	if err != nil {
		utils.ErrorJSON(w, err, "payload", http.StatusBadRequest)
		return
	}
	username, err := utils.GetUsernameFromRequestHeader(r)
	if err != nil {
		slog.Error(err.Error())
		utils.ErrorJSON(w, err, "username", http.StatusUnauthorized)
		return
	}
	currentPlanData, err := h.store.GetPlanDetails(payload.PlanName, "user", username)
	if err != nil {
		slog.Error(err.Error())
		utils.ErrorJSON(w, err, "planName", http.StatusBadRequest)
		return
	}
	utils.WriteJSON(w, http.StatusOK, currentPlanData)
}
