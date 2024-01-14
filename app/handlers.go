package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

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
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(struct {
			Error string `json:"error"`
		}{
			Error: "Failed to parse form data",
		})
		return
	}

	var ping Ping

	// Extract form values
	ping.ID = uuid.New().String()
	ping.Key = r.Form.Get("key")
	ping.CreatedAt = time.Now().Format("%d-%M-%Y %h:%m:%s")
	ping.UpdatedAt = ping.CreatedAt

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

	// TODO: add mongo code here to create a new ping
	// TODO: return the template view_ping_list
}

// CreatePingHandler the form to create a ping
func ViewPingListHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(
		template.ParseFiles(
			TEMPLATES_PATH + "/components/view_ping_list.html",
		),
	)

	pings := []Ping{
		Ping{
			CommonFields: CommonFields{
				ID: "an-id",
			},
			Key: "special-key",
		},
		Ping{
			CommonFields: CommonFields{
				ID: "the-id",
			},
			Key: "dok-key",
		},
	}

	tmplData := struct {
		Pings []Ping
	}{
		Pings: pings,
	}

	err := tmpl.Execute(w, tmplData)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// CreatePingHandler the form to create a ping
func CreatePingHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(
		template.ParseFiles(
			TEMPLATES_PATH + "/components/form_ping_create.html",
		),
	)

	err := tmpl.Execute(w, struct{}{})
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// DashboardHandler
func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(
		template.ParseFiles(TEMPLATES_PATH + "/dashboard.html"),
	)

	tmplData := struct {
		Title string
	}{
		Title: "Dashboard",
	}

	err := tmpl.Execute(w, tmplData)

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// IndexHandler for the index
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(
		template.ParseFiles(TEMPLATES_PATH + "/index.html"),
	)

	tmplData := struct {
		Title string
	}{
		Title: "Pingo",
	}

	err := tmpl.Execute(w, tmplData)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

// StrikeHandler: the handler for serving an svg for the link generated.
func StrikeHandler(w http.ResponseWriter, r *http.Request) {
	// Assuming the URL pattern is http://host/o/{pingKey}.svg
	urlPath := r.URL.Path
	urlParts := strings.Split(urlPath, "/")

	if len(urlParts) >= 3 {
		pingKey := urlParts[2]
		//TODO:
		// 	- add a new strike from the provided pingkey
		http.ServeFile(
			w,
			r,
			Format("%s/%s", ASSETS_PATH, pingKey), // the path of the pingKey
		)
	} else {
		// Handle invalid URL path
		http.NotFound(w, r)
	}
}
