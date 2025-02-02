package models

type Exercise struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ClientExercise struct {
	ClientID   int `json:"client_id"`
	ExerciseID int `json:"exercise_id"`
}
