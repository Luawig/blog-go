package handler

import (
	"blog-go/pkg/oss"
	"blog-go/pkg/utils"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		utils.ResponseError(c, utils.ErrorUploadSaveFile)
		return
	}

	url, code := oss.UploadFile(file)
	if code != utils.Success {
		utils.ResponseError(c, utils.ErrorUploadSaveFile)
		return
	}
	utils.ResponseSuccess(c, url)
}
