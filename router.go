package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// requestLoggerMiddleware middleware to add logger on each route
func requestLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf(">> received request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// responseHeadersMiddleware middleware to add some headers
func responseHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func BuildRouter() *mux.Router {
	router := mux.NewRouter()

	// middlewares
	router.Use(requestLoggerMiddleware)
	router.Use(responseHeadersMiddleware)

	router.HandleFunc("/", GetStatus).Methods("GET", "OPTIONS")

	// GET /pings
	// should be authentificated ang get from the current user.
	router.HandleFunc("/pings", GetPings).Methods("GET", "OPTIONS")
	// GET /pings/{id}
	// No need for authentification
	router.HandleFunc("/pings/{id}", GetPingByID).Methods("GET", "OPTIONS")
	// POST /pings
	router.HandleFunc("/pings", CreatePing).Methods("POST", "OPTIONS")

	return router
}
