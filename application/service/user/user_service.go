package user

import domainUser "fiber-gorm-microservice/domain/user"

type UserService interface {
	Create(newUser *NewUser) (*domainUser.User, error)
	GetaAll(page int64, limit int64) (*PaginationResultUser, error)
	GetById(id int) (*domainUser.User, error)
	Delete(id int) error
	Update(id int, userMap map[string]interface{}) (*domainUser.User, error)
}
