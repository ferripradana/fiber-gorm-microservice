package medicine

import (
	domainMedicine "fiber-gorm-microservice/domain/medicine"
	"time"
)

type Medicine struct {
	ID          int       `json:"id" example:"1" gorm:"primaryKey"`
	Name        string    `json:"name" example:"Bodrexin" gorm:"unique"`
	Description string    `json:"description" example:"Some Description"`
	EANCode     string    `json:"ean_code" example:"9900000124" gorm:"unique"`
	Laboratory  string    `json:"laboratory" example:"Bayer"`
	CreatedAt   time.Time `json:"created_at" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt   time.Time `json:"updated_at" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
}

func (*Medicine) TableName() string {
	return "medicines"
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
