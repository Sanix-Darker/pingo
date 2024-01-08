package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// GetStatus
func GetStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(struct {
		Status  string `json:"status"`
		Version string `json:"version"`
	}{
		Status:  "Ok",
		Version: VERSION,
	})
}

// GetPings
func GetPings(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Println("> Get Pings...")

	// TODO: add mongo code here
	var pings []Ping
	json.NewEncoder(w).Encode(pings)
}

// GetPingByID
func GetPingByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	// TODO: add mongo code here
	var pings []Ping

	for _, ping := range pings {
		if ping.ID == id {
			json.NewEncoder(w).Encode(ping)
			return
		}
	}

	log.Println("> Get Pings...")

	//TODO: serve the SVG here
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(struct {
		Error string `json:"error"`
	}{
		Error: "Ping not found",
	})
}

// CreatePing handler to create a ping
func CreatePing(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var ping Ping
	err := json.NewDecoder(r.Body).Decode(&ping)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(struct {
			Error string `json:"error"`
		}{
			Error: "Invalid request body",
		})
		return
	}

	ping.ID = uuid.New().String()

	log.Println("> Create Ping...")
	// Save the ping to MongoDB
	collection := MongoClient().Database(MONGO_DB).Collection(PING_COLLECTION)
	_, err = collection.InsertOne(nil, ping)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct {
			Error string `json:"error"`
		}{
			Error: "Failed to save ping",
		})
		return
	}

	// TODO: add mongo code here
	var pings []Ping
	// Append the ping to the in-memory slice
	pings = append(pings, ping)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ping)
}

// IndexHandler
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	data := struct {
		Title string
	}{
		Title: "Pingo",
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
