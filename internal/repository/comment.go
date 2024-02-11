package repository

import (
	"blog-go/internal/db"
	"blog-go/internal/model"
	"blog-go/utils"
	"errors"

	"gorm.io/gorm"
)

// CreateComment adds a comment to the database, and returns a status code.
func CreateComment(comment *model.Comment) int {
	err := db.DB.Create(comment).Error
	if err != nil {
		return utils.UnknownErr
	}
	return utils.Success
}

// GetComment gets a comment's information from the database, and returns the comment and a status code.
func GetComment(id int) (*model.Comment, int) {
	var comment model.Comment
	err := db.DB.Where("id = ?", id).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "username", "email")
		}).
		Preload("Article", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "title", "created_at", "updated_at", "comment_count", "read_count")
		}).
		First(&comment).Error
	if err != nil {
		return nil, utils.UnknownErr
	}
	return &comment, utils.Success
}

// GetCommentList gets a list of comments from the database, and returns the list and a status code.
func GetCommentList(pageSize, pageNum int) ([]*model.Comment, int) {
	var comments []*model.Comment
	err := db.DB.Model(&model.Comment{}).
		Select("id", "content", "created_at", "user_id", "article_id").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "username", "email")
		}).
		Preload("Article", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "title", "created_at", "updated_at", "comment_count", "read_count")
		}).
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&comments).Error
	if err != nil {
		return nil, utils.UnknownErr
	}
	return comments, utils.Success
}

// GetCommentListByArticle gets a list of comments from the database by article, and returns the list and a status code.
func GetCommentListByArticle(articleId, pageSize, pageNum int) ([]*model.Comment, int) {
	var comments []*model.Comment
	err := db.DB.Model(&model.Comment{}).
		Select("ID", "Content", "CreatedAt", "UserID", "ArticleID").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "username", "email")
		}).
		Preload("Article", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "title", "created_at", "updated_at", "comment_count", "read_count")
		}).
		Where("article_id = ?", articleId).
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&comments).Error
	if err != nil {
		return nil, utils.UnknownErr
	}
	return comments, utils.Success
}

// GetCommentUserID gets a comment's user id from the database, and returns the user id and a status code.
func GetCommentUserID(id int) (uint, int) {
	var comment model.Comment
	err := db.DB.Select("user_id").Where("id = ?", id).First(&comment).Error
	if err != nil {
		return 0, utils.UnknownErr
	}
	return comment.UserID, utils.Success
}

// UpdateComment edits a comment in the database, and returns a status code.
func UpdateComment(id int, data *model.Comment) int {
	var comment model.Comment
	err := db.DB.Where("id = ?", id).First(&comment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.ErrorCommentNotExist
		}
		return utils.UnknownErr
	}

	comment.ID = uint(id)
	err = db.DB.Model(&comment).Updates(data).Error
	if err != nil {
		return utils.UnknownErr
	}
	return utils.Success
}

// DeleteComment deletes a comment from the database, and returns a status code.
func DeleteComment(id int) int {
	err := db.DB.Delete(&model.Comment{}, id).Error
	if err != nil {
		return utils.UnknownErr
	}
	return utils.Success
}
