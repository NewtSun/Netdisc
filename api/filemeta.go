package api

import (
	"Netdisc/service"
	"Netdisc/util"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// UploadFile 文件上传
func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(200, ErrorResponse(err))
		return
	}
	service := service.UploadFileService{
		FileName: file.Filename,
		FileSize: file.Size,
		Location: "/tmp/" + file.Filename,
		UploadAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	newFile, err := os.Create(service.Location)
	if err != nil {
		fmt.Printf("Failed to create file, err:%s\n", err.Error())
		return
	}
	defer newFile.Close()
	newFile.Seek(0, 0)
	service.FileSha1 = util.FileSha1(newFile)

	// 游标重新回到文件头部
	newFile.Seek(0, 0)

	fmt.Println(file.Filename)
	if err != nil {
		fmt.Printf("Failed to create file, err:%s\n", err.Error())
		return
	}
	defer newFile.Close()
	dst := fmt.Sprintf("/tmp/%s", file.Filename)
	c.SaveUploadedFile(file, dst)
	c.JSON(200, gin.H{"message": fmt.Sprintf("'%s' uploaded!", file.Filename)})
}

// // ShowVideo 视频详情接口
// func ShowVideo(c *gin.Context) {
// 	service := service.ShowVideoService{}
// 	if err := c.ShouldBind(&service); err == nil {
// 		res := service.Show(c.Param("id"))
// 		c.JSON(200, res)
// 	} else {
// 		c.JSON(200, ErrorResponse(err))
// 	}
// }

// // ListVideo 视频列表接口
// func ListVideo(c *gin.Context) {
// 	service := service.ListVideoService{}
// 	if err := c.ShouldBind(&service); err == nil {
// 		res := service.List()
// 		c.JSON(200, res)
// 	} else {
// 		c.JSON(200, ErrorResponse(err))
// 	}
// }

// // UpdateVideo 更新视频接口
// func UpdateVideo(c *gin.Context) {
// 	service := service.UpdateVideoService{}
// 	if err := c.ShouldBind(&service); err == nil {
// 		res := service.Update(c.Param("id"))
// 		c.JSON(200, res)
// 	} else {
// 		c.JSON(200, ErrorResponse(err))
// 	}
// }

// // DeleteVideo 删除视频接口
// func DeleteVideo(c *gin.Context) {
// 	service := service.DeleteVideoService{}
// 	// if err := c.ShouldBind(&service); err == nil {
// 	res := service.Delete(c.Param("id"))
// 	c.JSON(200, res)
// 	// } else {
// 	// c.JSON(200, ErrorResponse(err))
// 	// }
// }
