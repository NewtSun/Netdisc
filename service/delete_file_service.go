package service

import (
	"Netdisc/model"
	"Netdisc/serializer"
	"os"
)

// DeleteFileService 删除文件及元信息的服务
type DeleteFileService struct {
}

// Delete 删除文件及元信息
func (service *DeleteFileService) Delete(fileSha1 string) serializer.Response {
	var filemeta model.FileMeta
	err := model.DB.Where("file_sha1 = ?", fileSha1).First(&filemeta).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "文件不存在",
			Error: err.Error(),
		}
	}
	os.Remove(filemeta.Location)
	err = model.DB.Delete(&filemeta).Error
	if err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "文件删除失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{}
}
