package user

import (
	"encoding/json"
	"fiber-gorm-microservice/domain/errors"
	domainUser "fiber-gorm-microservice/domain/user"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (u *UserRepositoryImpl) GetOneByMap(userMap map[string]interface{}) (*domainUser.User, error) {
	var userData User

	tx := u.DB.Where(userMap).Where("status = ?", true).Limit(1).Find(&userData)
	if tx.Error != nil {
		err := tx.Error
		return &domainUser.User{}, err
	}

	return userData.toDomainMapper(), nil
}

func (u *UserRepositoryImpl) Create(newUser *domainUser.User) (createdUser *domainUser.User, err error) {
	user := fromDomainMapper(newUser)
	tx := u.DB.Create(user)

	if tx.Error != nil {
		byteErr, _ := json.Marshal(tx.Error)
		var newError errors.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return createdUser, err
		}
		switch newError.Number {
		case 1062:
			err = errors.NewAppErrorWithType(errors.ResourceAlreadyExists)
		default:
			err = errors.NewAppErrorWithType(errors.UnknownError)
		}
		return createdUser, err
	}
	createdUser = user.toDomainMapper()
	return createdUser, err
}
