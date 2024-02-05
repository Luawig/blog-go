package handler

import (
	"blog-go/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)

	code := repository.CreateArt(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
