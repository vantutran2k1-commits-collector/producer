package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	setupRoutes()

	apiV1 := router.Group("/api/v1")
	apiV1.POST("/collect", c.CommitController.Collect)

	return router
}
