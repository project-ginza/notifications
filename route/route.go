package router

import (
	"github.com/gin-gonic/gin"
	"github.com/project-ginza/notifications/request"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", request.Home)

	return router
}
