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
		FileName: item.FileName,
		FileSize: item.FileSize,
		Location: item.Location,
		UploadAt: item.UploadAt,
	}
}

// BuildFileMetas 序列化视频列表
func BuildFileMetas(items []model.FileMeta) (filemetas []FileMeta) {
	for _, item := range items {
		filemeta := BuildFileMeta(item)
		filemetas = append(filemetas, filemeta)
	}
	return filemetas
}
