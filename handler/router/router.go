package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()
	// todoService := service.NewTODOService(todoDB)
	// todoHandler := handler.NewTODOHandler(todoService)
	healthzHandler := handler.NewHealthzHandler()
	mux.Handle("/healthz", healthzHandler)

	return mux
}
