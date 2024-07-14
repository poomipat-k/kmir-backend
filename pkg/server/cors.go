package server

import (
	"fmt"
	"log"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func handleCORS(mux *chi.Mux) {
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
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "withCredentials"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
}
