package service

import (
	"Netdisc/model"
	"Netdisc/serializer"
)

// UpdateVideoService 更新视频的服务
type UpdateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info  string `form:"info" json:"info" binding:"min=0,max=300"`
}

// Update 创建视频
func (service *UpdateVideoService) Update(fileSha1 string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, fileSha1).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "视频不存在",
			Error: err.Error(),
		}
	}
	video.Title = service.Title
	video.Info = service.Info
	err = model.DB.Save(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  50002,
			Msg:   "视频更新失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
