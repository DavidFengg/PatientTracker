package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"database/sql"
	_"github.com/go-sql-driver/mysql"

	"github.com/dgrijalva/jwt-go"
	// jwt "github.com/dgrijalva/jwt-go"
)

type Patient struct {
	ID			string `json:"id"`
	FirstName   string `json:"firstName"`
	LastName	string `json:"lastName"`
	Diagnosis	string `json:"diagnosis"`
	Physician	string `json:"physician"`
	DOV			string `json:"dov"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//temp
var users = map[string]string {
	"test": "test",
}

var signingKey = []byte("secret")
var patients []Patient

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[creds.Username]

	if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(2 * time.Hour)

	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims {
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie {
		Name: "token",
		Value: tokenString,
		Expires: expirationTime,
	})

}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return signingKey, nil
			})

			checkErr(err)
			fmt.Println("ok")
			if token.Valid {
				endpoint(w, r)
			}
		} else {
			fmt.Println("Not Authorized")
		}
	})
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getPatients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enableCORS(&w)

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
	enableCORS(&w)

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

	fmt.Println("New Entity Added")
}

func updatePatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	enableCORS(&w)

	params := mux.Vars(r)

	// update each field
	stmt, err := db.Prepare("UPDATE patient SET first_name = ?, last_name = ?, diagnosis = ?, physician = ?, dov = ? WHERE id = ?")
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

func deletePatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enableCORS(&w)

	params := mux.Vars(r)

	stmt, err := db.Prepare("DELETE FROM patient WHERE id = ?")
	checkErr(err)

	_, err = stmt.Exec(params["id"])
	checkErr(err)

	fmt.Println("Entity deleted")
}

var db *sql.DB
var err error

func main() {
	// db connection to local mac host
	db, err = sql.Open("mysql", "root:abcd1234@tcp(docker.for.mac.localhost:3306)/rest_api")
	// db, err = sql.Open("mysql", "root:abcd1234@tcp(localhost:3306)/rest_api")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connection made")

	defer db.Close();

	router := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	//routing
	router.HandleFunc("/login", login).Methods("POST")
	router.Handle("/patient", isAuthorized(getPatients)).Methods("GET")
	// router.HandleFunc("/patient", getPatients).Methods("GET")
	router.HandleFunc("/patient/{id}", getPatient).Methods("GET")
	router.HandleFunc("/patient", createPatient).Methods("POST")
	router.HandleFunc("/patient/{id}", updatePatient).Methods("PUT")	
	router.HandleFunc("/patient/{id}", deletePatient).Methods("DELETE")	

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(headers, methods, origins)(router)))
}