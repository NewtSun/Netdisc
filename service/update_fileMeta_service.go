package service

import (
	"Netdisc/model"
	"Netdisc/serializer"
)

// UpdateFileMetaService 更新文件元信息的服务
type UpdateFileMetaService struct {
	FileName string `form:"filename" json:"filename" binding:"required,min=2,max=30"`
	// Info  string `form:"info" json:"info" binding:"min=0,max=300"`
}

// Update 更新文件元信息
func (service *UpdateFileMetaService) Update(fileSha1 string) serializer.Response {
	var filemeta model.FileMeta
	err := model.DB.Where("file_sha1 = ?", fileSha1).First(&filemeta).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "文件不存在",
			Error: err.Error(),
		}
	}
	filemeta.FileName = service.FileName
	err = model.DB.Save(&filemeta).Error
	if err != nil {
		return serializer.Response{
			Code:  50002,
			Msg:   "文件元信息更新失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildFileMeta(filemeta),
	}
}
