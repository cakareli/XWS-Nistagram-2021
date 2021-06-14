package repository

import (
	"fmt"
	"gorm.io/gorm"
	"XWS-Nistagram-2021/backend-nistagram/authenticationService/model"
)

type AuthenticationRepository struct {
	Database *gorm.DB
}

func (repository *AuthenticationRepository) RegisterUser(user *model.User) error {
	result := repository.Database.Create(user)
	if result.RowsAffected == 0 {
		return fmt.Errorf("User not registered!")
	}
	fmt.Println("User successfuly registered!")
	return nil
}

func (repository *AuthenticationRepository) UpdateUser(user *model.User) error {
	result := repository.Database.Updates(user)
	if result.RowsAffected == 0 {
		return fmt.Errorf("User not update!")
	}
	fmt.Println("User successfuly updated!")
	return nil
}

func (repository *AuthenticationRepository) FindUserByUsername(username string) (*model.User, error){
	user := &model.User{}
	err := repository.Database.Table("users").First(&user, "username = ?", username).Error
	if  err != nil {
		return nil, err
	}
	return user, nil
}

func (repository *AuthenticationRepository) FindUserByUserId(userId string) (*model.User, error){
	user := &model.User{}
	err := repository.Database.Table("users").First(&user, "user_id = ?", userId).Error
	if  err != nil {
		return nil, err
	}
	return user, nil
}