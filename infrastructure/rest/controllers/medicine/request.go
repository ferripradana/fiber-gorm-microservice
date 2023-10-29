package medicine

type NewMedicineRequest struct {
	Name        string `json:"name" gorm:"unique" validate:"required"`
	Description string `json:"description" validate:"required"`
	Laboratory  string `json:"laboratory" validate:"required"`
	EanCode     string `json:"ean_code" gorm:"unique" validate:"required"`
}
