package auth

import (
	"errors"
	"fiber-gorm-microservice/application/security/jwt"
	domainError "fiber-gorm-microservice/domain/errors"
	userRepository "fiber-gorm-microservice/infrastructure/repository/user"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthServiceImpl struct {
	UserRepository userRepository.UserRepository
}

func NewAuthServiceImpl(repository userRepository.UserRepository) AuthService {
	return &AuthServiceImpl{
		UserRepository: repository,
	}
}

func (authServ *AuthServiceImpl) Login(loginUser LoginUser) (*SecurityAuthenticatedUser, error) {
	userMap := make(map[string]interface{})
	userMap["email"] = loginUser.Email
	domainUser, err := authServ.UserRepository.GetOneByMap(userMap)
	if err != nil {
		return &SecurityAuthenticatedUser{}, err
	}
	if domainUser.ID == 0 {
		return &SecurityAuthenticatedUser{}, errors.New("email or password does not match")
	}

	isAuthenticated := CheckPasswordHash(loginUser.Password, domainUser.HashPassword)
	if !isAuthenticated {
		return &SecurityAuthenticatedUser{}, errors.New("email or password does not match")
	}

	accessTokenClaims, err := jwt.GenerateJWTToken(domainUser.ID, "access")
	if err != nil {
		return &SecurityAuthenticatedUser{}, err
	}

	refreshTokenClaims, err := jwt.GenerateJWTToken(domainUser.ID, "refresh")
	if err != nil {
		return &SecurityAuthenticatedUser{}, err
	}

	return secAuthUserMapper(
		domainUser,
		&Auth{
			AccessToken:               accessTokenClaims.Token,
			RefreshToken:              refreshTokenClaims.Token,
			ExpirationAccessDateTime:  accessTokenClaims.ExpirationTime,
			ExpirationRefreshDateTime: accessTokenClaims.ExpirationTime,
		},
	), nil
}

func (authServ *AuthServiceImpl) AccessTokenByRefreshToken(refreshToken string) (*SecurityAuthenticatedUser, error) {
	claimsMap, err := jwt.GetClaimsAndVerifyToken(refreshToken, "refresh")
	if err != nil {
		return nil, err
	}
	userMap := map[string]interface{}{"id": claimsMap["id"]}
	domainUser, err := authServ.UserRepository.GetOneByMap(userMap)
	if err != nil {
		return nil, domainError.NewAppErrorImpl(err, domainError.RepositoryError, fiber.StatusInternalServerError)
	}
	if domainUser.ID == 0 {
		return &SecurityAuthenticatedUser{}, domainError.NewAppErrorWithType(domainError.NotFound)
	}

	accessTokenClaims, err := jwt.GenerateJWTToken(domainUser.ID, "access")
	if err != nil {
		return &SecurityAuthenticatedUser{}, domainError.NewAppErrorImpl(err, domainError.UnknownError, fiber.StatusInternalServerError)
	}

	var expTime = int64(claimsMap["exp"].(float64))
	return secAuthUserMapper(
		domainUser,
		&Auth{
			AccessToken:               accessTokenClaims.Token,
			RefreshToken:              refreshToken,
			ExpirationAccessDateTime:  accessTokenClaims.ExpirationTime,
			ExpirationRefreshDateTime: time.Unix(expTime, 0),
		},
	), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
