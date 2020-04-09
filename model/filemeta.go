package model

import "github.com/jinzhu/gorm"

// Filemeta 视频模型
type Filemeta struct {
	gorm.Model
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}
