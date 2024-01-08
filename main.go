package main

import (
	"fmt"
	"strconv"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := BuildRouter()

	// List all routes
	err := router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Route:", pathTemplate)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("> Pingo started successfully on %d...\n", PING_PORT)
	port, _ := strconv.Atoi(PING_PORT)
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		router,
	))
}
