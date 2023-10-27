package medicine

import "time"

type Medicine struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	EANCode     string    `json:"ean_code"`
	Laboratory  string    `json:"laboratory"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
