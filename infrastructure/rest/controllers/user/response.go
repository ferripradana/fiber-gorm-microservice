package user

import "time"

type ResponseUser struct {
	ID        int       `json:"id" example:"1099"`
	UserName  string    `json:"username" example:"BossonH"`
	Email     string    `json:"email" example:"some@mail.com"`
	FirstName string    `json:"first_name" example:"John"`
	LastName  string    `json:"last_name" example:"Doe"`
	Status    bool      `json:"status" example:"false"`
	CreatedAt time.Time `json:"created_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
}

type PaginationResultUser struct {
	Data       *[]ResponseUser
	Total      int64
	Limit      int64
	Current    int64
	NextCursor uint
	PrevCursor uint
	NumPages   int64
}
