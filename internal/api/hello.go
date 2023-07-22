package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloHandler(c *gin.Context) {

	// use repo sample
	//repo, err := di.NewRepo()
	//if err != nil {
	//	errResponse(c)
	//	return
	//}
	//users, err := repo.GetUsers()

	c.JSON(http.StatusOK, okResponse{OK: true})
}
