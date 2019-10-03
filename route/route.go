package route

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	controllers "github.com/davidfengg/restAPI/controllers"
	login "github.com/davidfengg/restAPI/login"
)

func GetRoutes() {
	router := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	//routing
	router.HandleFunc("/login", login.Login).Methods("POST")
	router.HandleFunc("/patient", controllers.GetPatients).Methods("GET")
	router.HandleFunc("/patient/{id}", controllers.GetPatient).Methods("GET")
	router.HandleFunc("/patient", controllers.CreatePatient).Methods("POST")
	router.HandleFunc("/patient/{id}", controllers.UpdatePatient).Methods("PUT")	
	router.HandleFunc("/patient/{id}", controllers.DeletePatient).Methods("DELETE")	

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(headers, methods, origins)(router)))

}