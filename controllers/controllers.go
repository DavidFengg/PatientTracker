package controllers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"math/rand"

	"github.com/gorilla/mux"
	_"github.com/go-sql-driver/mysql"

	"github.com/davidfengg/PatientTracker/models"
	"github.com/davidfengg/PatientTracker/database"
)

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func GetPatients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enableCORS(&w)

	var patients []models.Patient

	result, err := database.Db.Query("SELECT * from patient")
	checkErr(err)

	defer result.Close();

	var patient models.Patient

	for result.Next() {
		err := result.Scan(&patient.ID, &patient.FirstName, &patient.LastName, &patient.Diagnosis, &patient.Physician, &patient.DOV)
		checkErr(err)

		patients = append(patients, patient)
	}

	json.NewEncoder(w).Encode(patients)
}

func GetPatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enableCORS(&w)

	var patients []models.Patient

	params := mux.Vars(r)

	result, err := database.Db.Query("SELECT * from patient WHERE id = ?", params["id"])
	checkErr(err)

	defer result.Close()

	var patient models.Patient

	for result.Next() {
		err := result.Scan(&patient.ID, &patient.FirstName, &patient.LastName, &patient.Diagnosis, &patient.Physician, &patient.DOV)
		checkErr(err)

		patients = append(patients, patient)
	}

	json.NewEncoder(w).Encode(patients)
}

func CreatePatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	enableCORS(&w)

	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	firstName := keyVal["firstName"]
	lastName := keyVal["lastName"]
	diagnosis := keyVal["diagnosis"]
	physician := keyVal["physician"]
	dov := keyVal["dov"]

	var patient models.Patient
	var randID int
	numRows := 0
	
	// create unique ID
	for unique := true; unique; unique = (numRows != 0) {
		// generate random ID
		randID = rand.Intn(1000000000)

		result, err := database.Db.Query("SELECT id from patient WHERE id = ?", randID)
		checkErr(err)

		// check if ID is unique
		for result.Next() {
			numRows++
			err := result.Scan(&patient.ID)
			checkErr(err)
		}	
	}

	stmt, err := database.Db.Prepare("INSERT INTO patient VALUES(?,?,?,?,?,?)")
	checkErr(err)

	_, err = stmt.Exec(randID, firstName, lastName, diagnosis, physician, dov)
	checkErr(err)

	fmt.Println("New Entity Added")
}

func UpdatePatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	enableCORS(&w)

	params := mux.Vars(r)

	// update each field
	stmt, err := database.Db.Prepare("UPDATE patient SET first_name = ?, last_name = ?, diagnosis = ?, physician = ?, dov = ? WHERE id = ?")
	checkErr(err)

	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	newFirstName := keyVal["firstName"]
	newLastName := keyVal["lastName"]
	newDiagnosis := keyVal["diagnosis"]
	newPhysician := keyVal["physician"]
	newDOV := keyVal["dov"]

	_, err = stmt.Exec(newFirstName, newLastName, newDiagnosis, newPhysician, newDOV, params["id"])
	checkErr(err)
}

func DeletePatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enableCORS(&w)

	params := mux.Vars(r)

	stmt, err := database.Db.Prepare("DELETE FROM patient WHERE id = ?")
	checkErr(err)

	_, err = stmt.Exec(params["id"])
	checkErr(err)

	fmt.Println("Entity deleted")
}