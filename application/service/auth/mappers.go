package auth

import (
	domainUser "fiber-gorm-microservice/domain/user"
)

func secAuthUserMapper(user *domainUser.User, authInfo *Auth) *SecurityAuthenticatedUser {
	return &SecurityAuthenticatedUser{
		Data: DataUserAuthenticated{
			ID:        user.ID,
			UserName:  user.UserName,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Status:    user.Status,
		},
		Security: DataSecurityAuthenticated{
			JWTAccessToken:            authInfo.AccessToken,
			JWTRefreshToken:           authInfo.RefreshToken,
			ExpirationAccessDateTime:  authInfo.ExpirationAccessDateTime,
			ExpirationRefreshDateTime: authInfo.ExpirationRefreshDateTime,
		},
	}
}

func (registerUser *RegisterUser) toDomainMapper() *domainUser.User {
	return &domainUser.User{
		UserName:  registerUser.UserName,
		Email:     registerUser.Email,
		FirstName: registerUser.FirstName,
		LastName:  registerUser.LastName,
		Status:    false,
	}
}
