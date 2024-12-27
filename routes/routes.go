package routes

import (
	"log"
	"net/http"

	"github.com/douglastaylorb/api-go-rest/controllers"
	"github.com/douglastaylorb/api-go-rest/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalidades", controllers.GetPersonalidades).Methods("GET")
	r.HandleFunc("/personalidades/{id}", controllers.GetPersonalidade).Methods("GET")
	r.HandleFunc("/personalidades", controllers.CreatePersonalidade).Methods("POST")
	r.HandleFunc("/personalidades/{id}", controllers.DeletePersonalidade).Methods("DELETE")
	r.HandleFunc("/personalidades/{id}", controllers.UpdatePersonalidade).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
