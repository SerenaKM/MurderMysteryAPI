package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/serenakm/MurderMysteryAPI/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/health", app.HealthCheck)
	r.Get("/cases/{id}", app.MysteryHandler.HandleGetMysteryByID)

	r.Post("/cases", app.MysteryHandler.HandleCreateMystery)
	r.Delete("/cases/{id}", app.MysteryHandler.HandleDeleteCase)
	return r
}