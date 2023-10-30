package user

import domainUser "fiber-gorm-microservice/domain/user"

type UserService interface {
	Create(newUser *NewUser) (*domainUser.User, error)
}
