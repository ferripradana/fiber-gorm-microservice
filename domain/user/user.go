package user

import "time"

type User struct {
	ID           int       `json:"id"`
	UserName     string    `json:"user_name"`
	Email        string    `json:"email"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Status       bool      `json:"status"`
	HashPassword string    `json:"hash_password"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
