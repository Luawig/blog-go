package handler

import (
	"blog-go/internal/model"
	"blog-go/internal/repository"
	"blog-go/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateCategory - Creates a category
// @Summary Create a category
// @Tags category
// @Accept json
// @Produce json
// @Param category body model.Category true "Category"
// @Success 200 {object} utils.Response
// @Router /api/category [post]
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

// GetCategory - Gets a single category by ID
// @Summary Get a category
// @Tags category
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /api/category/{id} [get]
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

// GetCategoryList - Gets a list of categories
// @Summary List categories
// @Tags category
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response
// @Router /api/categories [get]
func GetCategoryList(c *gin.Context) {
	categories, code := repository.GetCategoryList()
	if code != utils.Success {
		utils.ResponseError(c, code)
		return
	}

	utils.ResponseSuccess(c, categories)
}

// UpdateCategory - Updates a category
// @Summary Update a category
// @Tags category
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body model.Category true "Category"
// @Success 200 {object} utils.Response
// @Router /api/category/{id} [put]
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

// DeleteCategory - Deletes a category
// @Summary Delete a category
// @Tags category
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} utils.Response
// @Router /api/category/{id} [delete]
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
