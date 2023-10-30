package user

import domainUser "fiber-gorm-microservice/domain/user"

type UserRepository interface {
	GetOneByMap(userMap map[string]interface{}) (*domainUser.User, error)
	Create(newUser *domainUser.User) (*domainUser.User, error)
	GetById(id int) (*domainUser.User, error)
	Delete(id int) error
	Update(id int, userMap map[string]interface{}) (*domainUser.User, error)
	GetAll(page int64, limit int64) (*PaginationResultUser, error)
}
