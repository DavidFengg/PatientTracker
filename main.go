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

	for result.Next() {
		var patient Patient
		err := result.Scan(&patient.ID, &patient.FirstName, &patient.LastName, &patient.Diagnosis, &patient.Physician, &patient.DOV)
		checkErr(err)

		patients = append(patients, patient)
	}

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

	//
	id := r.FormValue("id")
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	diagnosis := r.FormValue("diagnosis")
	physician := r.FormValue("physician")
	dov := r.FormValue("dov")

	stmt, err := db.Prepare("INSERT INTO patient VALUES(?,?,?,?,?,?)")
	checkErr(err)

	_, err = stmt.Exec(id, firstName, lastName, diagnosis, physician, dov)
	checkErr(err)

	fmt.Fprintln(w, id, firstName, lastName)

	// var newPatient Patient
	// json.NewDecoder(r.Body).Decode(&newPatient)
	// newPatient.ID = strconv.Itoa(len(patients) + 1)

	// patients = append(patients, newPatient)

	// json.NewEncoder(w).Encode(newPatient)
}

func updatePatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	stmt, err = db.Prepare("UPDATE patient SET ")

	for i, item := range patients {
		// check if id matches
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

	stmt, err := db.Prepare("DELETE FROM patient WHERE id = ?")
	checkErr(err)

	_, err = stmt.Exec(params["id"])
	checkErr(err)

	// for i, item := range patients {
	// 	if item.ID == params["id"] {
	// 		patients = append(patients[:i], patients[i+1:]...)
	// 		break
	// 	}
	// }
	//
	// json.NewEncoder(w).Encode(patients)
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

	// test data
	// patients = append(patients, 
	// 	Patient{ID: "1", FirstName: "David", LastName: "Feng", Diagnosis: "asd", Physician: "Dr.as", DOV: "asdas"},
	// 	Patient{ID: "2", FirstName: "test", LastName: "t", Diagnosis: "asd", Physician: "Dr.as", DOV: "sfds"},
	// )

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