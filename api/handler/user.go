package handler

import (
	"blog-go/internal/model"
	"blog-go/internal/repository"
	"blog-go/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var data model.User
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.ResponseInvalidParam(c)
		return
	}

	code := repository.CreateUser(&data)
	if code != utils.Success {
		utils.ResponseError(c, code)
		return
	}
	utils.ResponseSuccess(c, nil)
}

func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ResponseInvalidParam(c)
		return
	}

	user, code := repository.GetUser(id)
	if code != utils.Success {
		utils.ResponseError(c, code)
		return
	}

	utils.ResponseSuccess(c, user)
}

func GetUserList(c *gin.Context) {
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

	users, code := repository.GetUserList(pageSize, pageNum)
	if code != utils.Success {
		utils.ResponseError(c, code)
		return
	}

	utils.ResponseSuccess(c, users)
}

func GetUserListByUsername(c *gin.Context) {
	username := c.Param("username")
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

	users, code := repository.GetUserListByUsername(username, pageSize, pageNum)
	if code != utils.Success {
		utils.ResponseError(c, code)
		return
	}

	utils.ResponseSuccess(c, users)
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ResponseInvalidParam(c)
		return
	}

	var data model.User
	err = c.ShouldBindJSON(&data)
	if err != nil {
		utils.ResponseInvalidParam(c)
		return
	}

	code := repository.UpdateUser(id, &data)
	if code != utils.Success {
		utils.ResponseError(c, code)
		return
	}

	utils.ResponseSuccess(c, nil)
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ResponseInvalidParam(c)
		return
	}

	code := repository.DeleteUser(id)
	if code != utils.Success {
		utils.ResponseError(c, code)
		return
	}

	utils.ResponseSuccess(c, nil)
}
