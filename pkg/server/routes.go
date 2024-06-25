package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/poomipat-k/kmir-backend/pkg/utils"
)

type Server struct{}

func (app *Server) Routes(db *sql.DB) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	// specify who is allowed to connect
	var allowedOrigins []string
	stage := os.Getenv("STAGE")
	uiUrl := os.Getenv("UI_URL")
	if stage == "" {
		log.Fatal("ENV STAGE is required")
	}
	if uiUrl == "" {
		log.Fatal("ENV UI_URL is required")
	}
	if stage == "develop" {
		allowedOrigins = []string{fmt.Sprintf("http://%s", uiUrl)}
	} else if stage == "staging" || stage == "production" {
		allowedOrigins = []string{fmt.Sprintf("http://%s", uiUrl), fmt.Sprintf("https://%s", uiUrl)}
	}
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "withCredentials"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Route("/api/v1", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			utils.WriteJSON(w, http.StatusOK, "API landing Page")
		})

	})

	return mux
}
