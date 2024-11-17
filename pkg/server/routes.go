package server

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	appMiddleware "github.com/poomipat-k/kmir-backend/pkg/middleware"
	"github.com/poomipat-k/kmir-backend/pkg/plan"
	"github.com/poomipat-k/kmir-backend/pkg/user"
	"github.com/poomipat-k/kmir-backend/pkg/utils"
)

type Server struct{}

func (app *Server) Routes(db *sql.DB) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	handleCORS(mux)

	userStore := user.NewStore(db)
	userHandler := user.NewUserHandler(userStore)

	planStore := plan.NewStore(db)
	planHandler := plan.NewPlanHandler(planStore)

	mux.Route("/api/v1", func(r chi.Router) {
		// Todo: remove this endpoint before go prod
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			utils.WriteJSON(w, http.StatusOK, "KMIR API landing Page")
		})
		// Todo: remove this endpoint before go prod
		// r.Post("/hash-password", userHandler.GenerateHashedPassword)

		r.Post("/auth/login", userHandler.Login)
		r.Post("/auth/logout", userHandler.Logout)
		r.Post("/auth/refresh-token", userHandler.RefreshAccessToken)
		r.Get("/auth/current", appMiddleware.IsLoggedIn(userHandler.GetCurrentUser))

		r.Get("/plan/preview/all", appMiddleware.IsLoggedIn(planHandler.GetAllPreviewPlan))
		r.Get("/plan/access/{planName}", appMiddleware.IsLoggedIn(planHandler.CanAccessPlanDetails))
		r.Get("/plan/details/{planName}", appMiddleware.IsLoggedIn(planHandler.GetPlanDetails))
		r.Get("/plan/edit/{planName}", appMiddleware.IsLoggedIn(planHandler.CanEditPlan))

		r.Get("/admin/plans", appMiddleware.IsAdminOrViewer(planHandler.GetAllPlanDetails))
		r.Post("/admin/scores", appMiddleware.IsAdminOrViewer(planHandler.AdminGetScores))
		r.Patch("/admin/dashboard", appMiddleware.IsAdmin(planHandler.AdminEdit))

		r.Patch("/plan", appMiddleware.IsUser(planHandler.UserEditPlan))

	})

	return mux
}
