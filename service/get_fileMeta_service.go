package service

import (
	"Netdisc/model"
	"Netdisc/serializer"
)

// GetFileMetaService 文件详情的服务
type GetFileMetaService struct {
}

// GetFileMeta 创建视频
func (service *GetFileMetaService) GetFileMeta(fileSha1 string) serializer.Response {
	var filemeta model.FileMeta
	// err := model.DB.First(&filemeta, fileSha1).Error
	err := model.DB.Where("file_sha1 = ?", fileSha1).First(&filemeta).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "文件不存在",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildFileMeta(filemeta),
	}
}
