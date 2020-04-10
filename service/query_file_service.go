package service

import (
	"Netdisc/model"
	"Netdisc/serializer"
)

// QueryFileService 视频列表的服务
type QueryFileService struct {
}

// List 创建列表
func (service *QueryFileService) List() serializer.Response {
	var filemetas []model.FileMeta
	err := model.DB.Find(&filemetas).Error
	if err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库连接错误",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildFileMetas(filemetas),
	}
}
