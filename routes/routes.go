package routes

import (
	"go-api-rest/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.Index).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.Show).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.Store).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.Update).Methods("Put")
	r.HandleFunc("/api/personalidades/{id}", controllers.Delete).Methods("Delete")
	log.Fatal(http.ListenAndServe(":8000", r))
}