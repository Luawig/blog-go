package handler

import (
	"blog-go/internal/model"
	"blog-go/internal/repository"
	"blog-go/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateComment - Creates a comment
// @Summary Create a comment
// @Tags comment
// @Accept json
// @Produce json
// @Param comment body model.Comment true "Comment"
// @Success 200 {object} utils.Response
// @Router /api/comment [post]
func CreateComment(c *gin.Context) {
	var data model.Comment
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.ResponseInvalidParam(c)
		return
	}

	code := repository.CreateComment(&data)
	if code != utils.Success {
		utils.ResponseError(c, code)
		return
	}
	utils.ResponseSuccess(c, nil)
}

// GetComment - Gets a single comment by ID
// @Summary Get a comment
// @Tags comment
// @Accept json
// @Produce json
// @Param id path int true "Comment ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /api/comment/{id} [get]
func GetComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ResponseInvalidParam(c)
		return
	}

	comment, code := repository.GetComment(id)
	if code != utils.Success {
		utils.ResponseError(c, code)
		return
	}

	utils.ResponseSuccess(c, comment)
}

// GetCommentList - Gets a list of comments with pagination
// @Summary List comments
// @Tags comment
// @Accept json
// @Produce json
// @Param page_size query int false "Page Size"
// @Param page_num query int false "Page Number"
// @Success 200 {object} utils.Response
// @Router /api/comments [get]
func GetCommentList(c *gin.Context) {
	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if err != nil {
		utils.ResponseInvalidParam(c)
		return
	}
	pageNum, err := strconv.Atoi(c.DefaultQuery("page_num", "1"))
	if err != nil {
		utils.ResponseInvalidParam(c)
		return
	}

	if pageSize > 100 {
		pageSize = 100
	}

	comments, code := repository.GetCommentList(pageSize, pageNum)
	if code != utils.Success {
		utils.ResponseError(c, code)
		return
	}

	utils.ResponseSuccess(c, comments)
}

// GetCommentListByArticle - Gets a list of comments for a specific article with pagination
// @Summary List comments by article
// @Tags comment
// @Accept json
// @Produce json
// @Param id path int true "Article ID"
// @Param page_size query int false "Page Size"
// @Param page_num query int false "Page Number"
// @Success 200 {object} utils.Response
// @Router /api/comments/article/{id} [get]
func GetCommentListByArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ResponseInvalidParam(c)
		return
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if err != nil {
		utils.ResponseInvalidParam(c)
		return
	}
	pageNum, err := strconv.Atoi(c.DefaultQuery("page_num", "1"))
	if err != nil {
		utils.ResponseInvalidParam(c)
		return
	}

	if pageSize > 100 {
		pageSize = 100
	}

	comments, code := repository.GetCommentListByArticle(id, pageSize, pageNum)
	if code != utils.Success {
		utils.ResponseError(c, code)
		return
	}

	utils.ResponseSuccess(c, comments)
}

// UpdateComment - Updates a comment by ID
// @Summary Update a comment
// @Tags comment
// @Accept json
// @Produce json
// @Param id path int true "Comment ID"
// @Param comment body model.Comment true "Comment"
// @Success 200 {object} utils.Response
// @Failure 403 "Permission Denied"
// @Router /api/comment/{id} [put]
func UpdateComment(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		utils.ResponseError(c, utils.UnknownErr)
		return
	}
	userID, ok := userID.(uint)
	if !ok {
		utils.ResponseError(c, utils.UnknownErr)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ResponseInvalidParam(c)
		return
	}

	var data model.Comment
	err = c.ShouldBindJSON(&data)
	if err != nil {
		utils.ResponseInvalidParam(c)
		return
	}

	uid, code := repository.GetCommentUserID(id)
	if code != utils.Success {
		utils.ResponseError(c, code)
		return
	}
	if uid != userID {
		utils.ResponseError(c, utils.ErrorPermissionDenied)
		return
	}

	code = repository.UpdateComment(id, &data)
	if code != utils.Success {
		utils.ResponseError(c, code)
		return
	}

	utils.ResponseSuccess(c, nil)
}

// DeleteComment - Deletes a comment by ID
// @Summary Delete a comment
// @Tags comment
// @Accept json
// @Produce json
// @Param id path int true "Comment ID"
// @Success 200 {object} utils.Response
// @Failure 403 "Permission Denied"
// @Router /api/comment/{id} [delete]
func DeleteComment(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		utils.ResponseError(c, utils.UnknownErr)
		return
	}
	userID, ok := userID.(uint)
	if !ok {
		utils.ResponseError(c, utils.UnknownErr)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ResponseInvalidParam(c)
		return
	}

	uid, code := repository.GetCommentUserID(id)
	if code != utils.Success {
		utils.ResponseError(c, code)
		return
	}
	if uid != userID {
		utils.ResponseError(c, utils.ErrorPermissionDenied)
		return
	}

	code = repository.DeleteComment(id)
	if code != utils.Success {
		utils.ResponseError(c, code)
		return
	}

	utils.ResponseSuccess(c, nil)
}
