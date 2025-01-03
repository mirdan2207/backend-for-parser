package app

import (
	"database/sql"
	"net/http"
	"testing_po/config"
	"testing_po/internal/auth"
	"testing_po/internal/database"
	"testing_po/internal/parser"
	"testing_po/internal/samples"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type App struct {
	Router *mux.Router
	DB *sql.DB
	Config *config.Config
}

func New(cfg *config.Config) *App {
	db := database.Connect(cfg.DatabaseURL)

	router := mux.NewRouter()

	auth.RegisterRoutes(router, db)
	samples.RegisterRoutes(router, db)
	parser.RegisterRoutes(router)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // You can specify allowed origins here
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	router.Use(corsHandler.Handler)
	
	return &App{
        Router: router,
        DB: db,
        Config: cfg,
    }
}

func (a *App) Run() error {
    return http.ListenAndServe(":"+a.Config.ServerPort, a.Router)
}