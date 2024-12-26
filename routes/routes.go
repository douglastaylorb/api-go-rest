package routes

import (
	"log"
	"net/http"

	"github.com/douglastaylorb/api-go-rest/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/personalidades", controllers.GetPersonalidades).Methods("GET")
	r.HandleFunc("/personalidades/{id}", controllers.GetPersonalidade).Methods("GET")
	r.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", r))
}
