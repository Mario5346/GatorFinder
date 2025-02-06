package main

import (
	_ "backend/docs" // Make sure this path is correct
	"backend/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	StartServer()
}
func StartServer() {
	router := mux.NewRouter()

	// Register API Routes
	routes.RegisterEventRoutes(router)

	// Swagger UI route
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Default homepage
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to GatorFinder API"))
	})

	// Get port from env or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
