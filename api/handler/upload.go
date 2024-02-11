package handler

import (
	"blog-go/utils"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		utils.ResponseError(c, utils.ErrorUploadSaveFile)
		return
	}

	url, code := utils.UploadFile(file)
	if code != utils.Success {
		utils.ResponseError(c, utils.ErrorUploadSaveFile)
		return
	}
	utils.ResponseSuccess(c, url)
}
