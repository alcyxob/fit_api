package models

type Client struct {
	ID        int    `json:"id"`
	TrainerID int    `json:"trainer_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}
