package models

// Event struct represents an event
type Event struct {
	ID               string `json:"id"`
	User             string `json:"username"`
	EventName        string `json:"name"`
	EventDescription string `json:"description"`
	Date             string `json:"date"`
}
