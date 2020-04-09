package serializer

import "Netdisc/model"

// FileMeta 视频序列化器
type FileMeta struct {
	// ID        uint   `json:"id"`
	// Title     string `json:"title"`
	// Info      string `json:"info"`
	// CreatedAt int64  `json:"created_at"`
	ID       uint   `json:"id"`
	FileSha1 string `json:"file_sha1"`
	FileName string `json:"file_name"`
	FileSize int64  `json:"file_size"`
	Location string `json:"location"`
	UploadAt string `json:"upload_at"`
}

// BuildFileMeta 序列化视频
func BuildFileMeta(item model.FileMeta) FileMeta {
	return FileMeta{
		ID:       item.ID,
		FileSha1: item.FileSha1,
		FileSize: item.FileSize,
		Location: item.Location,
		UploadAt: item.UploadAt,
	}
}

// BuildVideos 序列化视频列表
// func BuildVideos(items []model.Video) (videos []Video) {
// 	for _, item := range items {
// 		video := BuildVideo(item)
// 		videos = append(videos, video)
// 	}
// 	return videos
// }
