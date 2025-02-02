package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"trainer-app/models"
)

func GetClients(w http.ResponseWriter, r *http.Request) {
	params := r.PathValue("trainer_id")
	trainerID, _ := strconv.Atoi(params)

	rows, err := db.Query("SELECT id, name, email FROM clients WHERE trainer_id = ?", trainerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var clients []models.Client
	for rows.Next() {
		var client models.Client
		rows.Scan(&client.ID, &client.Name, &client.Email)
		client.TrainerID = trainerID
		clients = append(clients, client)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clients)
}

func CreateClient(w http.ResponseWriter, r *http.Request) {
	params := r.PathValue("trainer_id")
	trainerID, _ := strconv.Atoi(params)

	var client models.Client
	json.NewDecoder(r.Body).Decode(&client)
	client.TrainerID = trainerID

	result, err := db.Exec("INSERT INTO clients (trainer_id, name, email) VALUES (?, ?, ?)", client.TrainerID, client.Name, client.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	client.ID = int(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(client)
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	params := r.PathValue("id")
	id, _ := strconv.Atoi(params)

	var client models.Client
	json.NewDecoder(r.Body).Decode(&client)

	_, err := db.Exec("UPDATE clients SET name = ?, email = ? WHERE id = ?", client.Name, client.Email, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
