package user

import (
	domainUser "fiber-gorm-microservice/domain/user"
)

func (u *User) toDomainMapper() *domainUser.User {
	return &domainUser.User{
		ID:           u.ID,
		UserName:     u.UserName,
		Email:        u.Email,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		Status:       u.Status,
		HashPassword: u.HashPassword,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}

func fromDomainMapper(user *domainUser.User) *User {
	return &User{
		ID:           user.ID,
		UserName:     user.UserName,
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Status:       user.Status,
		HashPassword: user.HashPassword,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

func sliceToDomainMapper(users *[]User) *[]domainUser.User {
	userDomain := make([]domainUser.User, len(*users))
	for i, user := range *users {
		userDomain[i] = *user.toDomainMapper()
	}
	return &userDomain
}
