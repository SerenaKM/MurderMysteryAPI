package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/serenakm/MurderMysteryAPI/internal/api"
	"github.com/serenakm/MurderMysteryAPI/internal/store"
	"github.com/serenakm/MurderMysteryAPI/migrations"
)

type Application struct {
	Logger *log.Logger
	MysteryHandler *api.MysteryHandler
	DB *sql.DB
}

func NewApplication() (*Application, error) {
	pgDB, err := store.Open()
	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// stores
	mysteryStore := store.NewPostgresMysteryStore(pgDB)

	// handlers
	mysteryHandler := api.NewMysteryHandler(mysteryStore, logger)

	app := &Application {
		Logger: logger,
		MysteryHandler: mysteryHandler,
		DB: pgDB,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available\n")
}