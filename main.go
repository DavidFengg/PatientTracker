package main

/*
todo: combine first/last name into seperate struct
*/

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Patient struct {
	ID			string `json:"id"`
	FirstName   string `json:"firstName"`
	LastName	string `json:"lastName"`
	Diagnosis	string `json:"diagnosis"`
	Physician	string `json:"physician"`
	DOV			string `json:"dov"`
}

var patients []Patient

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home")
}

func getPatients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patients)

}

func getPatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range patients {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createPatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newPatient Patient
	json.NewDecoder(r.Body).Decode(&newPatient)
	newPatient.ID = strconv.Itoa(len(patients) + 1)

	patients = append(patients, newPatient)

	json.NewEncoder(w).Encode(newPatient)
}

func updatePatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for i, item := range patients {
		if item.ID == params["id"] {
			// remove patient for update
			patients = append(patients[:i], patients[i+1:]...)

			var newPatient Patient
			json.NewDecoder(r.Body).Decode(&newPatient)
			newPatient.ID = params["id"]
			patients = append(patients, newPatient)

			json.NewEncoder(w).Encode(newPatient)
			return
		}
	}
}

func deletePatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for i, item := range patients {
		if item.ID == params["id"] {
			patients = append(patients[:i], patients[i+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(patients)
}

func main() {
	patients = append(patients, 
		Patient{ID: "1", FirstName: "David", LastName: "Feng", Diagnosis: "asd", Physician: "Dr.as", DOV: "asdas"},
		Patient{ID: "2", FirstName: "test", LastName: "t", Diagnosis: "asd", Physician: "Dr.as", DOV: "sfds"},
	)

	router := mux.NewRouter()

	//routing
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/api/patient", getPatients).Methods("GET")
	router.HandleFunc("/api/patient/{id}", getPatient).Methods("GET")
	router.HandleFunc("/api/patient", createPatient).Methods("POST")
	router.HandleFunc("/api/patient/{id}", updatePatient).Methods("PUT")	
	router.HandleFunc("/api/patient/{id}", deletePatient).Methods("DELETE")	

	log.Fatal(http.ListenAndServe(":8080", router))
}