package repository

import (
	"blog-go/internal/db"
	"blog-go/internal/model"
	"blog-go/utils"
	"errors"

	"gorm.io/gorm"
)

// CreateArticle adds an article to the database, and returns a status code.
func CreateArticle(article *model.Article) int {
	err := db.DB.Create(article).Error
	if err != nil {
		return utils.UnknownErr
	}
	return utils.Success
}

// GetArticle gets an article's information from the database, and returns the article and a status code.
func GetArticle(id int) (*model.Article, int) {
	var article model.Article
	err := db.DB.Where("id = ?", id).First(&article).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.ErrorArticleNotExist
		}
		return nil, utils.UnknownErr
	}
	return &article, utils.Success
}

// GetArticleList gets a list of articles from the database, and returns the list and a status code.
func GetArticleList(pageSize, pageNum int) ([]model.Article, int) {
	var articles []model.Article
	err := db.DB.Model(&model.Article{}).
		Select("id", "title", "created_at", "updated_at", "comment_count", "read_count").
		Preload("Categories").
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&articles).Error
	if err != nil {
		return nil, utils.UnknownErr
	}
	return articles, utils.Success
}

// GetArticleListByCategory gets a list of articles from the database by category, and returns the list and a status code.
func GetArticleListByCategory(categoryId, pageSize, pageNum int) ([]model.Article, int) {
	var articles []model.Article
	err := db.DB.Select("articles.id", "title", "articles.created_at", "articles.updated_at", "comment_count", "read_count").
		Joins("JOIN article_categories on article_categories.article_id=articles.id").
		Joins("JOIN categories on categories.id=article_categories.category_id").
		Preload("Categories").
		Where("categories.id = ?", categoryId).
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&articles).Error
	if err != nil {
		return nil, utils.UnknownErr
	}
	return articles, utils.Success
}

// GetArticleListByTitle gets a list of articles from the database by title, and returns the list and a status code.
func GetArticleListByTitle(title string, pageSize, pageNum int) ([]model.Article, int) {
	var articles []model.Article
	err := db.DB.Select("articles.id", "title", "articles.created_at", "articles.updated_at", "comment_count", "read_count").
		Joins("JOIN article_categories on article_categories.article_id=articles.id").
		Joins("JOIN categories on categories.id=article_categories.category_id").
		Preload("Categories").
		Where("title like ?", "%"+title+"%").
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&articles).Error
	if err != nil {
		return nil, utils.UnknownErr
	}
	return articles, utils.Success
}

// UpdateArticle updates an article in the database, and returns a status code.
func UpdateArticle(id int, data *model.Article) int {
	var article model.Article
	err := db.DB.Where("id = ?", id).First(&article).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.ErrorArticleNotExist
		}
		return utils.UnknownErr
	}

	data.ID = uint(id)
	err = db.DB.Model(&article).Updates(data).Error
	if err != nil {
		return utils.UnknownErr
	}
	return utils.Success
}

// DeleteArticle deletes an article from the database, and returns a status code.
func DeleteArticle(id int) int {
	if err := db.DB.Where("article_id = ?", id).Delete(&model.Comment{}).Error; err != nil {
		return utils.UnknownErr
	}

	if err := db.DB.Where("id = ?", id).Delete(&model.Article{}).Error; err != nil {
		return utils.UnknownErr
	}

	return utils.Success
}
