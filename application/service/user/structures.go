package user

import domainUser "fiber-gorm-microservice/domain/user"

type NewUser struct {
	UserName  string `example:"UserName"`
	Email     string `example:"some@mail.com"`
	FirstName string `example:"John"`
	LastName  string `example:"Doe"`
	Password  string `example:"SomeHashPass"`
	Status    bool   `example:"true"`
}

type PaginationResultUser struct {
	Data       *[]domainUser.User
	Total      int64
	Limit      int64
	Current    int64
	NextCursor uint
	PrevCursor uint
	NumPages   int64
}
