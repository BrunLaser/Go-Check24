package server

import (
	"Go-Check24/database"
	"encoding/json"
	"fmt"
	"net/http"
)

func measurementGET(w http.ResponseWriter, db database.Database) {
	w.Header().Set("Content-Type", "application/json") //set the header
	measurements, _ := db.GetAllPoints()

	json.NewEncoder(w).Encode(measurements) //write json data
}

func measurementPOST(w http.ResponseWriter, r *http.Request, db database.Database) {
	//check for json header
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Kein JSON Header", http.StatusBadRequest)
		return
	}
	db.CreateBasicTable()
	newM := &database.Measurement{}

	if err := json.NewDecoder(r.Body).Decode(newM); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	db.InsertPoint(*newM)

	//location header
	location := fmt.Sprintf("/messpunkte/%d", newM.ID)
	w.Header().Set("Location", location)
	//Anwort senden
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(*newM)
}

func Measurement(w http.ResponseWriter, r *http.Request, db database.Database) {
	switch r.Method {
	case http.MethodGet:
		measurementGET(w, db)
	case http.MethodPost:
		measurementPOST(w, r, db)
	default:
		http.Error(w, "Unerlaubte Methode", http.StatusMethodNotAllowed)
	}

}
