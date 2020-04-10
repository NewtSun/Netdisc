package api

import (
	"Netdisc/service"
	"Netdisc/util"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// ReHf 返回页面
func ReHf(c *gin.Context) {
	data, err := ioutil.ReadFile("./static/view/index.html")
	if err != nil {
		io.WriteString(c.Writer, "internel server error")
		return
	}
	io.WriteString(c.Writer, string(data))
}

// UploadFile 文件上传
func UploadFile(c *gin.Context) {
	file, head, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(200, ErrorResponse(err))
		return
	}
	service := service.UploadFileService{
		FileName: head.Filename,
		FileSize: head.Size,
		Location: "./tmp/" + head.Filename,
		UploadAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	dst := fmt.Sprintf("./tmp/%s", head.Filename)
	newFile, err := os.Create(dst)
	if err != nil {
		fmt.Printf("Failed to create file, err:%s\n", err.Error())
		return
	}
	defer newFile.Close()
	_, err = io.Copy(newFile, file)
	if err != nil {
		fmt.Printf("Failed to save data into file, err:%s\n", err.Error())
		return
	}
	newFile.Seek(0, 0)
	fmt.Println(newFile)
	service.FileSha1 = util.FileSha1(newFile)
	fmt.Println(service.FileSha1)
	// 游标重新回到文件头部
	// newFile.Seek(0, 0)
	// fmt.Println(head.Filename, service)
	defer newFile.Close()
	// c.SaveUploadedFile(file, dst)
	res := service.Create()
	c.JSON(200, res)
	// c.JSON(200, gin.H{"message": fmt.Sprintf("'%s' uploaded!", file.Filename)})
}

// GetFileMeta 视频详情接口
func GetFileMeta(c *gin.Context) {
	service := service.GetFileMetaService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetFileMeta(c.Param("fileSha1"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// QueryFile 视频列表接口
func QueryFile(c *gin.Context) {
	service := service.QueryFileService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DownloadFile 文件下载接口
func DownloadFile(c *gin.Context) {
	service := service.DownloadFileService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.DownloadFile(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UpdateFileMeta 更新文件元信息接口
func UpdateFileMeta(c *gin.Context) {
	service := service.UpdateFileMetaService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c.Param("fileSha1"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteFile 删除文件及元信息接口
func DeleteFile(c *gin.Context) {
	service := service.DeleteFileService{}
	// if err := c.ShouldBind(&service); err == nil {
	res := service.Delete(c.Param("fileSha1"))
	c.JSON(200, res)
	// } else {
	// c.JSON(200, ErrorResponse(err))
	// }
}
