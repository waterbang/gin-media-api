package models

// File 文件数据结构
type FileItem struct {
	ID       int      `json:"id"`
	FileName string   `json:"fileName"`
	Size     int64    `json:"size"`
	Url      string   `json:"url"`
	FileType FileType `json:"fileType"`
}

type FileType string

const (
	Video FileType = "video"
	Image FileType = "image"
	Other FileType = "other"
)
