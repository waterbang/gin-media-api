package handlers

import (
	"gin-media-api/services"
	"gin-media-api/utils"

	"github.com/gin-gonic/gin"
)

// GetVideoList 获取视频列表 API
func GetVideoList(c *gin.Context) {
	videos := services.GetVideoList()
	utils.SuccessResponse(c, videos)
}
