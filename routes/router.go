package routes

import (
	"gin-media-api/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 配置路由
func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/files", handlers.GetFileList)
	}
}
