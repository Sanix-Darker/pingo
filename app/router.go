package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var API_VERSION = "/api/v1"

// requestLoggerMiddleware middleware to add logger on each route
func requestLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf(">> received request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func addApiRoutes(router *mux.Router) *mux.Router {
	// GET /pings
	// should be authentificated ang get from the current user.
	router.HandleFunc(
		Format("%s/pings", API_VERSION),
		GetPings,
	).Methods("GET", "OPTIONS")

	// GET /pings/{id}
	// No need for authentification
	router.HandleFunc(
		Format("%s/pings/{id}", API_VERSION),
		GetPingByID,
	).Methods("GET", "OPTIONS")

	// POST /pings
	router.HandleFunc(
		Format("%s/pings", API_VERSION),
		CreatePing,
	).Methods("POST")

	return router
}

// addViewRoutes for templates or statics
func addViewRoutes(router *mux.Router) *mux.Router {
	// pages
	router.HandleFunc("/", IndexHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/dashboard", DashboardHandler).Methods("GET", "OPTIONS")

	// components
	// forms
	router.HandleFunc("/form-ping-create", CreatePingHandler).Methods("GET")
	// views
	router.HandleFunc("/view-ping-list", ViewPingListHandler).Methods("GET")

	// rendering svgs
	router.HandleFunc("/o/{pingPath}", StrikeHandler).Methods("GET", "OPTIONS")

	return router
}

// setMiddlewares add middlewares to the router
func setMiddlewares(router *mux.Router) *mux.Router {
	// all middlewares
	router.Use(requestLoggerMiddleware)
	router.PathPrefix(
		"/static/",
	).Handler(
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("../static/")),
		),
	)

	return router
}

func BuildRouter() *mux.Router {
	router := mux.NewRouter()

	router = setMiddlewares(router)
	router = addViewRoutes(router)
	router = addApiRoutes(router)

	return router
}

// ListRoutes List all routes (for debug pupourses)
func ListRoutes(router *mux.Router) {
	if err := router.Walk(
		func(
			route *mux.Route,
			router *mux.Router,
			ancestors []*mux.Route,
		) error {
			pathTemplate, err := route.GetPathTemplate()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("> route:", pathTemplate)
			return nil
		}); err != nil {
		fmt.Println(err)
	}

}
