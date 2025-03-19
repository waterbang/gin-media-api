package models

// File 文件数据结构
type File struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Size int64  `json:"size"`
	Path string `json:"path"`
}
