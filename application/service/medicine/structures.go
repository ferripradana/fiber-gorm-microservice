package medicine

import domainMedicine "fiber-gorm-microservice/domain/medicine"

type NewMedicine struct {
	Name        string `json:"name" example:"Biogesic"`
	Description string `json:"description" example:"Paracetamol"`
	EANCode     string `json:"ean_code" example:"7837373"`
	Laboratory  string `json:"laboratory" example:"Bayer"`
}

type PaginationResultMedicine struct {
	Data       *[]domainMedicine.Medicine `json:"data"`
	Total      int64                      `json:"total"`
	Limit      int64                      `json:"limit"`
	Current    int64                      `json:"current"`
	NextCursor uint                       `json:"next_cursor"`
	PrevCursor uint                       `json:"prev_cursor"`
	NumPages   int64                      `json:"num_pages"`
}
