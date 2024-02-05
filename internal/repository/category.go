package repository

import (
	"blog-go/internal/db"
	"blog-go/internal/model"
	"blog-go/pkg/utils"
	"errors"

	"gorm.io/gorm"
)

// CheckCategoryName checks if a category name empty or exists in the database, and returns a status code.
func CheckCategoryName(name string) int {
	if name == "" {
		return utils.ErrorCategoryNameEmpty
	}
	var category model.Category
	err := db.DB.Where("name = ?", name).First(&category).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.Success
		}
		return utils.UnknownErr
	}
	return utils.ErrorCategoryNameUsed
}

// CreateCategory adds a category to the database, and returns a status code.
func CreateCategory(category *model.Category) int {
	if code := CheckCategoryName(category.Name); code != utils.Success {
		return code
	}
	err := db.DB.Create(category).Error
	if err != nil {
		return utils.UnknownErr
	}
	return utils.Success
}

// GetCategory gets a category's information from the database, and returns the category and a status code.
func GetCategory(id int) (*model.Category, int) {
	var category model.Category
	err := db.DB.Where("id = ?", id).First(&category).
		Preload("Articles", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "title", "created_at", "updated_at", "comment_count", "read_count")
		}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.ErrorCategoryNotExist
		}
		return nil, utils.UnknownErr
	}
	return &category, utils.Success
}

// GetCategoryList gets a list of categories from the database, and returns the list and a status code.
func GetCategoryList() ([]model.Category, int) {
	var categories []model.Category
	err := db.DB.Find(&categories).
		Preload("Articles", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "title", "created_at", "updated_at", "comment_count", "read_count")
		}).Error
	if err != nil {
		return nil, utils.UnknownErr
	}
	return categories, utils.Success
}

// UpdateCategory edits a category in the database, and returns a status code.
func UpdateCategory(id int, data *model.Category) int {
	var category model.Category
	err := db.DB.Where("id = ?", id).First(&category).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.ErrorCategoryNotExist
		}
		return utils.UnknownErr
	}

	if code := CheckCategoryName(data.Name); code != utils.Success {
		return code
	}

	err = db.DB.Model(&category).Updates(data).Error
	if err != nil {
		return utils.UnknownErr
	}
	return utils.Success
}

// DeleteCategory deletes a category from the database, and returns a status code.
func DeleteCategory(id int) int {
	var category model.Category
	err := db.DB.Where("id = ?", id).First(&category).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.ErrorCategoryNotExist
		}
		return utils.UnknownErr
	}

	err = db.DB.Model(&category).Association("Articles").Clear()
	if err != nil {
		return utils.UnknownErr
	}

	err = db.DB.Delete(&category).Error
	if err != nil {
		return utils.UnknownErr
	}
	return utils.Success
}
