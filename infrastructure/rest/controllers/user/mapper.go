package user

import (
	"fiber-gorm-microservice/application/service/user"
	domainUser "fiber-gorm-microservice/domain/user"
)

func toUserServiceMapper(request *NewUserRequest) *user.NewUser {
	return &user.NewUser{
		UserName:  request.UserName,
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Password:  request.Password,
		Status:    request.Status,
	}
}

func domainToResponseMapper(user *domainUser.User) *ResponseUser {
	return &ResponseUser{
		ID:        user.ID,
		UserName:  user.UserName,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
