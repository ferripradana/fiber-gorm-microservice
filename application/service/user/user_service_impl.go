package user

import (
	"fiber-gorm-microservice/domain/errors"
	domainUser "fiber-gorm-microservice/domain/user"
	userRepository "fiber-gorm-microservice/infrastructure/repository/user"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository userRepository.UserRepository
}

func NewUserServiceImpl(repository userRepository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: repository,
	}
}

func (service *UserServiceImpl) Create(newUser *NewUser) (*domainUser.User, error) {
	user := newUser.toDomainMapper()

	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return &domainUser.User{}, errors.NewAppErrorImpl(err, errors.UnknownError, fiber.StatusInternalServerError)
	}

	user.HashPassword = string(hash)
	return service.UserRepository.Create(user)
}
