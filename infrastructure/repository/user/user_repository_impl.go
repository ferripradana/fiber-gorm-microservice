package user

import (
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
