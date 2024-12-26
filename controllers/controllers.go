package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/douglastaylorb/api-go-rest/database"
	"github.com/douglastaylorb/api-go-rest/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func GetPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)

	json.NewEncoder(w).Encode(p)
}

func GetPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personalidade

	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)

}
