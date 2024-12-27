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

func CreatePersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade models.Personalidade

	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Create(&personalidade)

	json.NewEncoder(w).Encode(personalidade)
}

func DeletePersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personalidade

	database.DB.Delete(&p, id)
	json.NewEncoder(w).Encode("Personalidade deletada com sucesso")
}

func UpdatePersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personalidade

	database.DB.First(&p, id)
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Save(&p)

	json.NewEncoder(w).Encode(p)
}
