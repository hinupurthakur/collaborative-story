package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hinupurthakur/collaborative-story/middlewares"
)

func CreateRoutes() http.Handler {
	r := mux.NewRouter()
	r = r.PathPrefix("/api/v1").Subrouter()
	r.Handle("/", Routes(r))
	// wrap all routes
	return middlewares.EnableCORS(r)
}

func Routes(r *mux.Router) *mux.Router {
	r.HandleFunc("/health", HealthCheck).Methods("GET")
	r.HandleFunc("/add", AddNewWord).Methods("POST")
	r.HandleFunc("/stories", GetAllStories).Methods("GET")
	r.HandleFunc("/stories/{id:[0-9]+}", GetStory).Methods("GET")
	return r
}

func WordRoutes(r *mux.Router) *mux.Router {
	return r
}
