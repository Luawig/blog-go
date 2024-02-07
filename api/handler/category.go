package handler

import (
	"blog-go/internal/model"
	"blog-go/internal/repository"
	"blog-go/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var category model.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		utils.ResponseInvalidParam(c)
		return
	}
	code := repository.CreateCategory(&category)
	if code != utils.Success {
		utils.ResponseError(c, code)
		return
	}

	utils.ResponseSuccess(c, nil)
}

func GetCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ResponseInvalidParam(c)
		return
	}

	category, code := repository.GetCategory(id)
	if code != utils.Success {
		utils.ResponseError(c, code)
		return
	}

	utils.ResponseSuccess(c, category)
}

func GetCategoryList(c *gin.Context) {
	categories, code := repository.GetCategoryList()
	if code != utils.Success {
		utils.ResponseError(c, code)
		return
	}

	utils.ResponseSuccess(c, categories)
}

func UpdateCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ResponseInvalidParam(c)
		return
	}

	var category model.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		utils.ResponseInvalidParam(c)
		return
	}

	code := repository.UpdateCategory(id, &category)
	if code != utils.Success {
		utils.ResponseError(c, code)
		return
	}

	utils.ResponseSuccess(c, nil)
}

func DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ResponseInvalidParam(c)
		return
	}

	code := repository.DeleteCategory(id)
	if code != utils.Success {
		utils.ResponseError(c, code)
		return
	}

	utils.ResponseSuccess(c, nil)
}
