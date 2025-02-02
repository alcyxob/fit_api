package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"trainer-app/models"
)

func CreateExercise(w http.ResponseWriter, r *http.Request) {
	var exercise models.Exercise
	json.NewDecoder(r.Body).Decode(&exercise)

	result, err := db.Exec("INSERT INTO exercises (name, description) VALUES (?, ?)", exercise.Name, exercise.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	exercise.ID = int(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exercise)
}

func GetExercises(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, description FROM exercises")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var exercises []models.Exercise
	for rows.Next() {
		var exercise models.Exercise
		rows.Scan(&exercise.ID, &exercise.Name, &exercise.Description)
		exercises = append(exercises, exercise)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exercises)
}

func AssignExercise(w http.ResponseWriter, r *http.Request) {
	params := r.PathValue("client_id")
	clientID, _ := strconv.Atoi(params)

	var clientExercise models.ClientExercise
	json.NewDecoder(r.Body).Decode(&clientExercise)
	clientExercise.ClientID = clientID

	_, err := db.Exec("INSERT INTO client_exercises (client_id, exercise_id) VALUES (?, ?)", clientExercise.ClientID, clientExercise.ExerciseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
