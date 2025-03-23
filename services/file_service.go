package services

import "gin-media-api/models"

// GetFileList 获取文件列表
func GetFileList() []models.FileItem {
	return []models.FileItem{
		{ID: 1, FileName: "52.4.jpg", Size: 1048576, Url: "https://res-pri-dev.pixcheese.com/048b/145005872/ori/ori/11a2f213-165a-46b7-a260-cdfa71a6cdb4.jpg?Expires=1746028799&OSSAccessKeyId=LTAI5t8F4P46vvCSHYa2jfZo&Signature=cKUiBVVJ4LF19ILspa6DCHJHXr8%3D&et=1742384456", FileType: models.Image},
		{ID: 2, FileName: "图片.jpg", Size: 204800, Url: "https://res-pri-dev.pixcheese.com/887200/50d1d004-4768-4c87-b98e-6c04ec6a1435.png?Expires=1746028799&OSSAccessKeyId=LTAI5t8F4P46vvCSHYa2jfZo&Signature=QAts2BF3mr0OcmxXUCxpk24rBfk%3D&et=0", FileType: models.Image},
		{ID: 3, FileName: "视频.mp4", Size: 512000, Url: "https://res-pri-dev.pixcheese.com/048b%2F145005872%2Fvideo%2F20480b7c-3a57-4195-849d-237780018bca.mp4?Expires=1746028799&OSSAccessKeyId=LTAI5t8F4P46vvCSHYa2jfZo&Signature=LyJUox7o0H0plRGn338SNYK3dP8%3D&et=0", FileType: models.Video},
		{ID: 4, FileName: "324.88.mp4", Size: 512000, Url: "https://res-pri-dev.pixcheese.com/048b%2F145005872%2Fvideo%2Fde60e673-2c55-45da-b3fc-327d420cfdcc.mp4?Expires=1746028799&OSSAccessKeyId=LTAI5t8F4P46vvCSHYa2jfZo&Signature=Y4XbmNkWFSJRfbBWu50yE6hAOi8%3D&et=0", FileType: models.Video},
	}
}
