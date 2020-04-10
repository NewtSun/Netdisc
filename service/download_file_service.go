package service

import (
	"Netdisc/model"
	"Netdisc/serializer"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// DownloadFileService 文件详情的服务
type DownloadFileService struct {
}

// DownloadFile 创建视频
func (service *DownloadFileService) DownloadFile(c *gin.Context) serializer.Response {
	var filemeta model.FileMeta
	// err := model.DB.First(&filemeta, fileSha1).Error
	err := model.DB.Where("file_sha1 = ?", c.Param("fileSha1")).First(&filemeta).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "文件不存在",
			Error: err.Error(),
		}
	}
	f, err := os.Open(filemeta.Location)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}
	c.Writer.Header().Set("Content-Type", "application/octect-stream")
	c.Writer.Header().Set("content-disposition", "attachment; filename=\""+filemeta.FileName+"\"")
	c.Writer.Write(data)
	return serializer.Response{
		Data: serializer.BuildFileMeta(filemeta),
	}
}
