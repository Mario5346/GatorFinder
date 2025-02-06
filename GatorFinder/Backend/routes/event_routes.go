package routes

import (
	"backend/controllers"

	"github.com/gorilla/mux"
)

func RegisterEventRoutes(router *mux.Router) {
	router.HandleFunc("/events/add", controllers.AddEvent).Methods("POST")
	router.HandleFunc("/events/get", controllers.GetEvent).Methods("GET")
	router.HandleFunc("/events/delete", controllers.DeleteEvent).Methods("DELETE")
}
