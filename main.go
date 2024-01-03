package main

import (
	"fmt"

	"log"
	"net/http"
)

func main() {

	router := BuildRouter()

	log.Printf("> Pingo started successfully on %d...\n", PING_PORT)
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%d", PING_PORT),
		router,
	))
}
