package auth

type LoginRequest struct {
	Email    string `json:"email" gorm:"unique" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
