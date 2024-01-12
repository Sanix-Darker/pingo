package main

import (
	"fmt"
	"strconv"

	"log"
	"net/http"
)

func main() {

	router := BuildRouter()

	// ListRoutes(router)
	log.Printf("> Pingo started successfully on %d...\n", PING_PORT)
	port, _ := strconv.Atoi(PING_PORT)
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		router,
	))
}
