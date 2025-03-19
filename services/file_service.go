package services

import "gin-media-api/models"

// GetFileList 获取文件列表
func GetFileList() []models.File {
	return []models.File{
		{ID: 1, Name: "文档.pdf", Size: 1048576, Path: "/files/doc.pdf"},
		{ID: 2, Name: "图片.jpg", Size: 204800, Path: "/files/image.jpg"},
	}
}
