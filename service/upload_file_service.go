package service

import (
	"Netdisc/model"
	"Netdisc/serializer"
)

// UploadFileService 视频投稿的服务
type UploadFileService struct {
	// Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	// Info  string `form:"info" json:"info" binding:"min=0,max=300"`
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

// Create 创建视频
func (service *UploadFileService) Create() serializer.Response {
	File := model.FileMeta{
		FileSha1: service.FileSha1,
		FileName: service.FileName,
		FileSize: service.FileSize,
		Location: service.Location,
		UploadAt: service.UploadAt,
	}
	err := model.DB.Create(&File).Error
	if err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "文件元信息保存失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildFileMeta(File),
	}
}
