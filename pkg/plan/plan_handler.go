package plan

import (
	"errors"
	"fmt"
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
	EditPlan(planName string, payload EditPlanRequest, userRole string, username string, userId int) (string, error)
	GetAllPlanDetails(criteriaLen int) ([]AdminDashboardPlanDetailsRow, error)
	AdminGetScores(fromYear, toYear int, plan string) ([]AssessmentScore, error)
	GetAssessmentCriteria() ([]AssessmentCriteria, error)
	GetAdminNote() (string, error)
	GetOnlyLatestScore() ([]LatestScoreTimestamp, error)
	AdminEdit(payload AdminEditRequest, userId int) (bool, string, error)
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

func (h *PlanHandler) GetAllPlanDetails(w http.ResponseWriter, r *http.Request) {
	criteriaList, err := h.store.GetAssessmentCriteria()
	if err != nil {
		slog.Error(err.Error())
		utils.ErrorJSON(w, err, "", http.StatusInternalServerError)
		return
	}
	planDetails, err := h.store.GetAllPlanDetails(len(criteriaList))
	if err != nil {
		slog.Error(err.Error())
		utils.ErrorJSON(w, err, "", http.StatusInternalServerError)
		return
	}
	adminNote, err := h.store.GetAdminNote()
	if err != nil {
		slog.Error(err.Error())
		utils.ErrorJSON(w, err, "admin note", http.StatusInternalServerError)
		return
	}

	latestScores, err := h.store.GetOnlyLatestScore()
	if err != nil {
		slog.Error(err.Error())
		utils.ErrorJSON(w, err, "latest scores", http.StatusInternalServerError)
		return
	}

	response := AdminAllPlansDetailsResponse{
		AssessmentCriteria: criteriaList,
		PlanDetails:        planDetails,
		AdminNote:          adminNote,
		LatestScores:       latestScores,
	}
	utils.WriteJSON(w, http.StatusOK, response)
}

func (h *PlanHandler) AdminGetScores(w http.ResponseWriter, r *http.Request) {
	var payload AdminGetScoresRequest
	err := utils.ReadJSON(w, r, &payload)
	if err != nil {
		utils.ErrorJSON(w, err, "payload", http.StatusBadRequest)
		return
	}
	// validate payload
	data, err := h.store.AdminGetScores(payload.FromYear, payload.ToYear, payload.Plan)
	if err != nil {
		slog.Error(err.Error())
		utils.ErrorJSON(w, err, "store", http.StatusInternalServerError)
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

func (h *PlanHandler) UserEditPlan(w http.ResponseWriter, r *http.Request) {
	var payload EditPlanRequest
	err := utils.ReadJSON(w, r, &payload)
	if err != nil {
		utils.ErrorJSON(w, err, "payload", http.StatusBadRequest)
		return
	}
	userId, err := utils.GetUserIdFromRequestHeader(r)
	if err != nil {
		slog.Error(err.Error())
		utils.ErrorJSON(w, err, "userId", http.StatusUnauthorized)
		return
	}
	username, err := utils.GetUsernameFromRequestHeader(r)
	if err != nil {
		slog.Error(err.Error())
		utils.ErrorJSON(w, err, "username", http.StatusUnauthorized)
		return
	}
	// validation
	name, err := validateEditPlanPayload(payload)
	if err != nil {
		slog.Error(err.Error())
		utils.ErrorJSON(w, err, name, http.StatusBadRequest)
		return
	}
	if payload.AssessmentScore != nil {
		name, err := validateScore(payload.AssessmentScore)
		if err != nil {
			slog.Error(err.Error())
			utils.ErrorJSON(w, err, name, http.StatusBadRequest)
			return
		}
	}

	errName, err := h.store.EditPlan(payload.PlanName, payload, "user", username, userId)
	if err != nil {
		slog.Error(err.Error())
		utils.ErrorJSON(w, err, errName, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, common.CommonSuccessResponse{
		Success: true,
		Message: "update plan success",
	})
}

func (h *PlanHandler) AdminEdit(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.GetUserIdFromRequestHeader(r)
	if err != nil {
		slog.Error(err.Error())
		utils.ErrorJSON(w, err, "userId", http.StatusUnauthorized)
		return
	}

	var payload AdminEditRequest
	err = utils.ReadJSON(w, r, &payload)
	if err != nil {
		utils.ErrorJSON(w, err, "payload", http.StatusBadRequest)
		return
	}
	errName, err := validateAdminEditPayload(payload)
	if err != nil {
		slog.Error(err.Error())
		utils.ErrorJSON(w, err, errName, http.StatusBadRequest)
		return
	}

	updated, errName, err := h.store.AdminEdit(payload, userId)
	if err != nil {
		slog.Error(err.Error())
		utils.ErrorJSON(w, err, fmt.Sprintf("store: %s", errName), http.StatusBadRequest)
		return
	}

	if !updated {
		utils.WriteJSON(w, http.StatusOK, common.CommonSuccessResponse{
			Success: false,
			Message: "nothing changed",
		})
		return
	}
	utils.WriteJSON(w, http.StatusOK, common.CommonSuccessResponse{
		Success: true,
		Message: "admin update successfully",
	})

}
