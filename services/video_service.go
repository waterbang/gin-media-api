package services

import "gin-media-api/models"

// GetVideoList 获取视频列表
func GetVideoList() []models.Video {
	return []models.Video{
		{ID: 1, Title: "视频1", Thumbnail: "https://example.com/thumb1.jpg", Duration: "10:05", URL: "https://example.com/video1.mp4"},
		{ID: 2, Title: "视频2", Thumbnail: "https://example.com/thumb2.jpg", Duration: "12:30", URL: "https://example.com/video2.mp4"},
	}
}
