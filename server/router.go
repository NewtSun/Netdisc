package server

import (
	"Netdisc/api"
	"Netdisc/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户注册
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
		}

		v1.POST("videos", api.CreateVideo)
		v1.GET("video/:id", api.ShowVideo)
		v1.GET("videos", api.ListVideo)
		v1.PUT("video/:id", api.UpdateVideo)
		v1.DELETE("video/:id", api.DeleteVideo)
	}
	f1 := r.Group("/api/f1")
	{
		f1.POST("ping", api.Ping)

		// 用户注册
		f1.POST("user/register", api.UserRegister)

		// 用户登录
		f1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		auth := f1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
		}

		f1.GET("file/upload", api.ReHf)
		f1.POST("file/upload", api.UploadFile)
		// f1.GET("file/upload/suc", api.UploadSucFile)
		f1.GET("file/meta/:fileSha1", api.GetFileMeta)
		f1.GET("file/query", api.QueryFile)
		f1.GET("file/download/:fileSha1", api.DownloadFile)
		f1.PUT("file/update/:fileSha1", api.UpdateFileMeta)
		f1.DELETE("file/delete/:fileSha1", api.DeleteFile)
	}
	return r
}
