package handlers

import (
	"gin-media-api/services"
	"gin-media-api/utils"

	"github.com/gin-gonic/gin"
)

// GetFileList 获取文件列表 API
func GetFileList(c *gin.Context) {
	files := services.GetFileList()
	utils.SuccessResponse(c, files)
}
