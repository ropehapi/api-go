package controllers

import (
	"encoding/json"
	"fmt"
	"go-api-rest/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}

func RetornaPersonalidade(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id:= vars["id"]

	for _, personalidade := range models.Personalidades {
		if strconv.Itoa(personalidade.Id) == id{
			json.NewEncoder(w).Encode(personalidade)
		}
	}
}