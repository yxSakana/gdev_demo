package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yxSakana/gdev_demo/internal/dao"
	"github.com/yxSakana/gdev_demo/internal/logic/middleware"

	"github.com/yxSakana/gdev_demo/internal/controller/image"
	"github.com/yxSakana/gdev_demo/internal/controller/novel"
	"github.com/yxSakana/gdev_demo/internal/controller/user"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(dao.DBMiddleware())

	v1 := r.Group("/api/v1")

	userG := v1.Group("/user")
	{
		userG.POST("/login", user.Login)
		userG.POST("/register", user.Register)
	}

	v1.Use(middleware.Auth)
	novelG := v1.Group("/novel")
	{
		novelG.POST("/create", novel.CreateNovel)
		novelG.POST("/upload_chapter", novel.UploadChapter)
		novelG.GET("/:novel_id", novel.DetailNovel)
		novelG.GET("/query", novel.Query)
		novelG.POST("/update/:novel_id", novel.UpdateNovel)
		novelG.POST("/delete/:novel_id", novel.DelNovel)
	}
	imageG := v1.Group("/image")
	{
		imageG.GET("/:collection_id", image.DetailImgCollection)
		imageG.POST("/create", image.Create)
		imageG.POST("/create_image", image.UploadImage)
		imageG.POST("/create_images", image.UploadImages)
		imageG.POST("/update/:collection_id", image.UpdateImageCollection)
		imageG.POST("/delete/:collection_id", image.DelImageCollection)
	}

	return r
}
