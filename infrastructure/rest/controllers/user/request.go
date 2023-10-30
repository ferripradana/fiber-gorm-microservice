package user

type NewUserRequest struct {
	UserName  string `json:"username" example:"someUser" gorm:"unique" validate:"required"`
	Email     string `json:"email" example:"mail@mail.com" gorm:"unique" validate:"required,email"`
	FirstName string `json:"first_name" example:"John" validate:"required"`
	LastName  string `json:"last_name" example:"Doe" validate:"required"`
	Password  string `json:"password" example:"Password123" validate:"required"`
	Status    bool   `json:"status" example:"true" validate:"required"`
}
