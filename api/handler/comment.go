package handler

import (
	"blog-go/internal/model"
	"blog-go/internal/repository"
	"blog-go/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
