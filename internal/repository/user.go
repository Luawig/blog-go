package repository

import (
	"blog-go/internal/db"
	"blog-go/internal/model"
	"blog-go/pkg/utils"
	"errors"

	"gorm.io/gorm"
)

// CheckUsername checks if a user empty or exists in the database, and returns a status code.
func CheckUsername(username string) int {
	if username == "" {
		return utils.ErrorUsernameEmpty
	}
	var user model.User
	err := db.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.Success
		}
		return utils.UnknownErr
	}
	return utils.ErrorUsernameUsed
}

// CheckEmail checks if an email empty or exists in the database, and returns a status code.
func CheckEmail(email string) int {
	if email == "" {
		return utils.ErrorEmailEmpty
	}
	var user model.User
	err := db.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.Success
		}
		return utils.UnknownErr
	}
	return utils.ErrorEmailUsed
}

// CreateUser adds a user to the database, and returns a status code.
func CreateUser(user *model.User) int {
	if code := CheckUsername(user.Username); code != utils.Success {
		return code
	}
	if code := CheckEmail(user.Email); code != utils.Success {
		return code
	}
	if user.Password == "" {
		return utils.ErrorPasswordEmpty
	}
	err := db.DB.Create(user).Error
	if err != nil {
		return utils.UnknownErr
	}
	return utils.Success
}

// GetUser gets a user's information from the database, and returns the user and a status code.
func GetUser(id int) (*model.User, int) {
	var user model.User
	err := db.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.ErrorUserNotExist
		}
		return nil, utils.UnknownErr
	}
	return &user, utils.Success
}

// GetUserList gets a list of users from the database, and returns the list and a status code.
func GetUserList() ([]model.User, int) {
	var users []model.User
	err := db.DB.Find(&users).Error
	if err != nil {
		return nil, utils.UnknownErr
	}
	return users, utils.Success
}

// UpdateUser edits a user in the database, and returns a status code.
func UpdateUser(id int, data *model.User) int {
	var user model.User
	err := db.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.ErrorUserNotExist
		}
		return utils.UnknownErr
	}

	if code := CheckUsername(data.Username); code != utils.Success {
		return code
	}
	if code := CheckEmail(data.Email); code != utils.Success {
		return code
	}
	if data.Password == "" {
		return utils.ErrorPasswordEmpty
	}

	err = db.DB.Model(&user).Updates(data).Error
	if err != nil {
		return utils.UnknownErr
	}
	return utils.Success
}

// DeleteUser deletes a user from the database, and returns a status code.
func DeleteUser(id int) int {
	if err := db.DB.Where("user_id = ?", id).Delete(&model.Comment{}).Error; err != nil {
		return utils.UnknownErr
	}

	if err := db.DB.Where("id = ?", id).Delete(&model.User{}).Error; err != nil {
		return utils.UnknownErr
	}

	return utils.Success
}
