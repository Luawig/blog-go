package handler

import (
	"blog-go/internal/model"
	"blog-go/internal/repository"
	"blog-go/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateArticle - Creates an article
// @Summary Create an article
// @Tags article
// @Accept json
// @Produce json
// @Param article body model.Article true "Article"
// @Success 200 {object} utils.Response
// @Router /api/article [post]
func CreateArticle(c *gin.Context) {
	var article model.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		utils.ResponseInvalidParam(c)
		return
	}

	code := repository.CreateArticle(&article)
	if code != utils.Success {
		utils.ResponseError(c, code)
		return
	}

	utils.ResponseSuccess(c, article)
}

// GetArticle - Retrieves an article based on its ID
// @Summary Retrieve an article
// @Tags article
// @Accept json
// @Produce json
// @Param id path int true "Article ID"
// @Success 200 {object} utils.Response
// @Router /api/article/{id} [get]
func GetArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ResponseInvalidParam(c)
		return
	}

	article, code := repository.GetArticle(id)
	if code != utils.Success {
		utils.ResponseError(c, code)
		return
	}

	utils.ResponseSuccess(c, article)
}

// GetArticleList - Retrieves a list of articles with pagination
// @Summary Retrieve list of articles
// @Tags article
// @Accept json
// @Produce json
// @Param page_size query int false "Page Size" default(10)
// @Param page_num query int false "Page Number" default(1)
// @Success 200 {object} utils.Response
// @Router /api/articles [get]
func GetArticleList(c *gin.Context) {
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

	articles, code := repository.GetArticleList(pageSize, pageNum)
	if code != utils.Success {
		utils.ResponseError(c, code)
		return
	}

	utils.ResponseSuccess(c, articles)
}

// GetArticleListByCategory - Retrieves a list of articles by category with pagination
// @Summary Retrieve articles by category
// @Tags article
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param page_size query int false "Page Size" default(10)
// @Param page_num query int false "Page Number" default(1)
// @Success 200 {object} utils.Response
// @Router /api/articles/category/{id} [get]
func GetArticleListByCategory(c *gin.Context) {
	categoryId, err := strconv.Atoi(c.Param("id"))
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

	articles, code := repository.GetArticleListByCategory(categoryId, pageSize, pageNum)
	if code != utils.Success {
		utils.ResponseError(c, code)
		return
	}

	utils.ResponseSuccess(c, articles)
}

// GetArticleListByTitle - Retrieves a list of articles by title with pagination
// @Summary Retrieve articles by title
// @Tags article
// @Accept json
// @Produce json
// @Param title path string true "Article Title"
// @Param page_size query int false "Page Size" default(10)
// @Param page_num query int false "Page Number" default(1)
// @Success 200 {array} utils.Response
// @Router /api/articles/{title} [get]
func GetArticleListByTitle(c *gin.Context) {
	title := c.Param("title")
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

	articles, code := repository.GetArticleListByTitle(title, pageSize, pageNum)
	if code != utils.Success {
		utils.ResponseError(c, code)
		return
	}

	utils.ResponseSuccess(c, articles)
}

// UpdateArticle - Updates an article based on its ID
// @Summary Update an article
// @Tags article
// @Accept json
// @Produce json
// @Param id path int true "Article ID"
// @Param article body model.Article true "Article Update"
// @Success 200 {object} utils.Response
// @Router /api/article/{id} [put]
func UpdateArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ResponseInvalidParam(c)
		return
	}

	var article model.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		utils.ResponseInvalidParam(c)
		return
	}

	code := repository.UpdateArticle(id, &article)
	if code != utils.Success {
		utils.ResponseError(c, code)
		return
	}

	utils.ResponseSuccess(c, nil)
}

// DeleteArticle - Deletes an article based on its ID
// @Summary Delete an article
// @Tags article
// @Accept json
// @Produce json
// @Param id path int true "Article ID"
// @Success 200 {object} utils.Response
// @Router /api/article/{id} [delete]
func DeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ResponseInvalidParam(c)
		return
	}

	code := repository.DeleteArticle(id)
	if code != utils.Success {
		utils.ResponseError(c, code)
		return
	}

	utils.ResponseSuccess(c, nil)
}
