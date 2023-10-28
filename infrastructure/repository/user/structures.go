package user

import (
	domainUser "fiber-gorm-microservice/domain/user"
	"time"
)

type User struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	UserName     string    `json:"user_name" gorm:"column:user_name;unique"`
	Email        string    `json:"email" gorm:"unique"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Status       bool      `json:"status"`
	HashPassword string    `json:"hash_password"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime:mili"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime:mili"`
}

func (*User) name() string {
	return "users"
}

type PaginationResultUser struct {
	Data       *[]domainUser.User `json:"data"`
	Total      int64              `json:"total"`
	Limit      int64              `json:"limit"`
	Current    int64              `json:"current"`
	NextCursor uint               `json:"next_cursor"`
	PrevCursor uint               `json:"prev_cursor"`
	NumPages   int64              `json:"num_pages"`
}
