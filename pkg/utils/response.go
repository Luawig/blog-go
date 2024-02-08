package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	response := gin.H{
		"status":  Success,
		"message": GetMsg(Success),
	}
	if data != nil {
		response["data"] = data
	}
	c.JSON(http.StatusOK, response)
}

func ResponseInvalidParam(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  ErrorInvalidParam,
		"message": GetMsg(ErrorInvalidParam),
	})
}

func ResponseError(c *gin.Context, code int) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  code,
		"message": GetMsg(code),
	})
}

func ResponseAuthWrong(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"status":  ErrorTokenWrong,
		"message": GetMsg(ErrorTokenWrong),
	})
}
