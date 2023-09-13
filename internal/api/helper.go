package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type baseResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func internalError(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, baseResponse{
		Code: 500,
		Msg:  "internal error",
	})
}

func jsonResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, baseResponse{
		Code: 0,
		Data: data,
	})
}
