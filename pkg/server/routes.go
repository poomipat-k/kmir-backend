package server

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	appMiddleware "github.com/poomipat-k/kmir-backend/pkg/middleware"
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

	mux.Route("/api/v1", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			utils.WriteJSON(w, http.StatusOK, "API landing Page")
		})
		// Todo: remove this endpoint before go prod
		r.Post("/hash-password", userHandler.GenerateHashedPassword)

		r.Post("/auth/login", userHandler.Login)
		r.Get("/auth/current", appMiddleware.IsLoggedIn(userHandler.GetCurrentUser))
		r.Post("/auth/logout", userHandler.Logout)
	})

	return mux
}
