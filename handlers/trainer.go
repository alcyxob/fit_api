package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"trainer-app/models"
)

var db *sql.DB

// SetDB sets the database connection for the handlers.
func SetDB(database *sql.DB) {
	db = database
}

func CreateTrainer(w http.ResponseWriter, r *http.Request) {
	var trainer models.Trainer
	json.NewDecoder(r.Body).Decode(&trainer)

	result, err := db.Exec("INSERT INTO trainers (name, email) VALUES (?, ?)", trainer.Name, trainer.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	trainer.ID = int(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trainer)
}

func GetTrainer(w http.ResponseWriter, r *http.Request) {
	params := r.PathValue("id")
	id, _ := strconv.Atoi(params)

	var trainer models.Trainer
	err := db.QueryRow("SELECT id, name, email FROM trainers WHERE id = ?", id).Scan(&trainer.ID, &trainer.Name, &trainer.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trainer)
}

func UpdateTrainer(w http.ResponseWriter, r *http.Request) {
	params := r.PathValue("id")
	id, _ := strconv.Atoi(params)

	var trainer models.Trainer
	json.NewDecoder(r.Body).Decode(&trainer)

	_, err := db.Exec("UPDATE trainers SET name = ?, email = ? WHERE id = ?", trainer.Name, trainer.Email, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
