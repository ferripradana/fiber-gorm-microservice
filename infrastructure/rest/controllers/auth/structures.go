package auth

type LoginRequest struct {
	Email    string `json:"email" gorm:"unique" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
