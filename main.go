package main

import (
	"log"
	"net/http"

	"trainer-app/database"
	"trainer-app/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database
	db := database.InitDB()
	defer db.Close()

	// Set up routes
	r := mux.NewRouter()

	// Trainer routes
	r.HandleFunc("/trainers", handlers.CreateTrainer).Methods("POST")
	r.HandleFunc("/trainers/{id}", handlers.GetTrainer).Methods("GET")
	r.HandleFunc("/trainers/{id}", handlers.UpdateTrainer).Methods("PUT")

	// Client routes
	r.HandleFunc("/trainers/{trainer_id}/clients", handlers.GetClients).Methods("GET")
	r.HandleFunc("/trainers/{trainer_id}/clients", handlers.CreateClient).Methods("POST")
	r.HandleFunc("/clients/{id}", handlers.UpdateClient).Methods("PUT")

	// Exercise routes
	r.HandleFunc("/exercises", handlers.CreateExercise).Methods("POST")
	r.HandleFunc("/exercises", handlers.GetExercises).Methods("GET")

	// Assign exercises to clients
	r.HandleFunc("/clients/{client_id}/exercises", handlers.AssignExercise).Methods("POST")

	// Start the server
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
