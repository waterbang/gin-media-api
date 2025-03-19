package main

import (
	"gin-media-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 注册路由
	routes.SetupRoutes(r)

	// 启动服务器
	r.Run(":8080")
}
