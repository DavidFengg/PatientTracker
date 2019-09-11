package main

import (
	"fmt"
	//"encoding/json"
	"log"
	"net/http"
	//"math/rand"
	//"strconv"

	"github.com/gorilla/mux"
)

type Patient struct {
	ID			string `json:"id"`
	FirstName   string `json:"firstName"`
	LastName	string `json:"lastName"`
	Diagnosis	string `json:"diagnosis"`

}

//var patients []Patient

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home")
}

func getFirstName(w http.ResponseWriter, r *http.Request) {
}

func getLastName(w http.ResponseWriter, r *http.Request) {
}

func main() {
	router := mux.NewRouter()

	//routing
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/api/firstname", getFirstName).Methods("GET")
	router.HandleFunc("/api/lastname", getLastName).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}