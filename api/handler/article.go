package handler

import (
	"blog-go/internal/model"
	"blog-go/internal/repository"
	"blog-go/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
