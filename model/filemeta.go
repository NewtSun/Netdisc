package model

import "github.com/jinzhu/gorm"

// FileMeta 视频模型
type FileMeta struct {
	gorm.Model
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}
