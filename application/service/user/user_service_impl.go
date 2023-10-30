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

func (service *UserServiceImpl) GetaAll(page int64, limit int64) (*PaginationResultUser, error) {
	all, err := service.UserRepository.GetAll(page, limit)
	if err != nil {
		return &PaginationResultUser{}, err
	}
	return &PaginationResultUser{
		Data:       all.Data,
		Total:      all.Total,
		Limit:      all.Limit,
		Current:    all.Current,
		NextCursor: all.NextCursor,
		PrevCursor: all.PrevCursor,
		NumPages:   all.NumPages,
	}, nil
}

func (service *UserServiceImpl) GetById(id int) (*domainUser.User, error) {
	return service.UserRepository.GetById(id)
}

func (service *UserServiceImpl) Delete(id int) error {
	return service.UserRepository.Delete(id)
}

func (service *UserServiceImpl) Update(id int, userMap map[string]interface{}) (*domainUser.User, error) {
	return service.UserRepository.Update(id, userMap)
}
