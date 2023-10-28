package auth

type LoginRequest struct {
	Email    string `json:"email" gorm:"unique" binding:"required"`
	Password string `json:"password" binding:"required"`
}
