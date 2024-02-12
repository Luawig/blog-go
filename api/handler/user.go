package handler

import (
	"blog-go/internal/model"
	"blog-go/internal/repository"
	"blog-go/middleware"
	"blog-go/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser - Creates a user
// @Summary Create a user
// @Tags user
// @Accept json
// @Produce json
// @Param user body model.User true "User"
// @Success 200 {object} utils.Response
// @Router /api/user [post]
func CreateUser(c *gin.Context) {
	var data model.User
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.ResponseInvalidParam(c)
		return
	}

	_, err = utils.Validate(&data)
	if err != nil {
		utils.ResponseError(c, utils.UnknownErr)
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

// GetUser - Gets a single user by ID
// @Summary Get a user
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} utils.Response
// @Router /api/user/{id} [get]
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

// GetUserList - Gets a list of users with pagination
// @Summary List users
// @Tags user
// @Accept json
// @Produce json
// @Param page_size query int false "Page Size"
// @Param page_num query int false "Page Number"
// @Success 200 {object} utils.Response
// @Router /api/users [get]
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

// GetUserListByUsername - Gets a list of users filtered by username with pagination
// @Summary List users by username
// @Tags user
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Param page_size query int false "Page Size"
// @Param page_num query int false "Page Number"
// @Success 200 {object} utils.Response
// @Router /api/users/{username} [get]
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

// UpdateUser - Updates a user by ID
// @Summary Update a user
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body model.User true "User"
// @Success 200 {object} utils.Response
// @Router /api/user/{id} [put]
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

// UpdateUserPassword - Updates a user's password by ID
// @Summary Update a user's password
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param password body string true "New Password"
// @Success 200 {object} utils.Response
// @Router /api/user/{id}/password [put]
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

// DeleteUser - Deletes a user by ID
// @Summary Delete a user
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} utils.Response
// @Router /api/user/{id} [delete]
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

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login - Authenticates a user and returns a token
// @Summary Login a user
// @Tags auth
// @Accept json
// @Produce json
// @Param login body loginRequest true "Login Information"
// @Success 200 {object} utils.Response
// @Router /api/login [post]
func Login(c *gin.Context) {
	var loginInfo loginRequest
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
