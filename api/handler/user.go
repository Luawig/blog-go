package handler

import (
	"blog-go/internal/model"
	"blog-go/internal/repository"
	"blog-go/middleware"
	"blog-go/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	var data model.User
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.ResponseInvalidParam(c)
		return
	}

	data.Password, err = encryptUserPassword(data.Password)
	if err != nil {
		utils.ResponseError(c, utils.UnknownErr)
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

	if userID != uint(id) {
		utils.ResponseError(c, utils.ErrorPermissionDenied)
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

func UpdateUserPassword(c *gin.Context) {
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

	if userID != uint(id) {
		utils.ResponseError(c, utils.ErrorPermissionDenied)
		return
	}

	var data model.User
	err = c.ShouldBindJSON(&data)
	if err != nil {
		utils.ResponseInvalidParam(c)
		return
	}

	data.Password, err = encryptUserPassword(data.Password)
	if err != nil {
		utils.ResponseError(c, utils.UnknownErr)
		return
	}

	code := repository.UpdateUserPassword(id, &data)
	if code != utils.Success {
		utils.ResponseError(c, code)
		return
	}

	utils.ResponseSuccess(c, nil)
}

func DeleteUser(c *gin.Context) {
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

	if userID != uint(id) {
		utils.ResponseError(c, utils.ErrorPermissionDenied)
		return
	}

	code := repository.DeleteUser(id)
	if code != utils.Success {
		utils.ResponseError(c, code)
		return
	}

	utils.ResponseSuccess(c, nil)
}

func encryptUserPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func Login(c *gin.Context) {
	var loginInfo struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := c.ShouldBindJSON(&loginInfo)
	if err != nil {
		utils.ResponseInvalidParam(c)
		return
	}

	user, code := repository.GetUserWithPasswordByUsername(loginInfo.Username)
	if code != utils.Success {
		utils.ResponseError(c, code)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password))
	if err != nil {
		utils.ResponseError(c, utils.ErrorPasswordWrong)
		return
	}

	token, err := middleware.GenerateToken(user.ID, user.Username)
	if err != nil {
		utils.ResponseError(c, utils.UnknownErr)
		return
	}

	utils.ResponseSuccess(c, token)
}
