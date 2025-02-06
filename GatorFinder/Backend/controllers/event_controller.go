package controllers

import (
	"backend/models" // ✅ Update import path
	"encoding/json"
	"net/http"
)

// Sample events list
var events []models.Event

// @Summary Add a new event
// @Description Adds a new event to the system
// @Tags Events
// @Accept  json
// @Produce  json
// @Param event body models.Event true "Event Data"
// @Success 200 {object} models.Event
// @Router /events/add [post]
func AddEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent models.Event
	_ = json.NewDecoder(r.Body).Decode(&newEvent)
	events = append(events, newEvent)
	json.NewEncoder(w).Encode(map[string]string{"message": "Event added successfully"})
	println("add user")
}

// @Summary Get event details
// @Description Adds a new event to the system
// @Tags  Events
// @Accept  json
// @Produce  json
// @Param event body models.Event true "Event Data"
// @Success 200 {object} models.Event
// @Router /events/get [get]
func GetEvent(w http.ResponseWriter, r *http.Request) {
	println("Fetched Event")

}

// @Summary Delete an event
// @Description Adds a new event to the system
// @Tags Events
// @Accept  json
// @Produce  json
// @Param event body models.Event true "Event Data"
// @Success 200 {object} models.Event
// @Router /events/delete [delete]
func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	println("Delete Event")
}
