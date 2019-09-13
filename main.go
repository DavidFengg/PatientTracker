package main

/*
todo:
- fix update patient func
- generate unique patient id's
- combine first/last name into seperate struct

*/

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"io/ioutil"
	"math/rand"
	//"strconv"
	//"context"

	"github.com/gorilla/mux"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getPatients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var patients []Patient

	result, err := db.Query("SELECT * from patient")
	checkErr(err)

	defer result.Close();

	var patient Patient

	for result.Next() {
		err := result.Scan(&patient.ID, &patient.FirstName, &patient.LastName, &patient.Diagnosis, &patient.Physician, &patient.DOV)
		checkErr(err)

		patients = append(patients, patient)
	}

	json.NewEncoder(w).Encode(patients)
}

func getPatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var patients []Patient

	params := mux.Vars(r)

	result, err := db.Query("SELECT * from patient WHERE id = ?", params["id"])
	checkErr(err)

	defer result.Close()

	var patient Patient

	for result.Next() {
		err := result.Scan(&patient.ID, &patient.FirstName, &patient.LastName, &patient.Diagnosis, &patient.Physician, &patient.DOV)
		checkErr(err)

		patients = append(patients, patient)
	}

	json.NewEncoder(w).Encode(patients)
}

func createPatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	diagnosis := r.FormValue("diagnosis")
	physician := r.FormValue("physician")
	dov := r.FormValue("dov")

	var patient Patient
	var randID int
	numRows := 0
	
	// create unique ID
	for unique := true; unique; unique = (numRows != 0) {
		// generate random ID
		randID = rand.Intn(1000000000)

		result, err := db.Query("SELECT id from patient WHERE id = ?", randID)
		checkErr(err)

		// check if ID is unique
		for result.Next() {
			numRows++
			err := result.Scan(&patient.ID)
			checkErr(err)
		}	
	}

	stmt, err := db.Prepare("INSERT INTO patient VALUES(?,?,?,?,?,?)")
	checkErr(err)

	_, err = stmt.Exec(randID, firstName, lastName, diagnosis, physician, dov)
	checkErr(err)

	fmt.Fprintln(w, randID, firstName, lastName)
}

func updatePatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// updates first name
	stmt, err := db.Prepare("UPDATE patient SET first_name = ? WHERE id = ?")
	checkErr(err)

	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	newFirstName := keyVal["firstName"]

	_, err = stmt.Exec(newFirstName, params["id"])
	checkErr(err)
}

func deletePatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	stmt, err := db.Prepare("DELETE FROM patient WHERE id = ?")
	checkErr(err)

	_, err = stmt.Exec(params["id"])
	checkErr(err)
}

var db *sql.DB
var err error

func main() {
	// db connection
	db, err = sql.Open("mysql", "root:abcd1234@tcp(127.0.0.1:3306)/rest_api")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close();

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