package router

import (
	"gblog-server/controller"
	"gblog-server/middleware"

	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.GET("/user", middleware.AuthMiddleware(), controller.GetUserInfo)
	r.POST("/upload/avatar", controller.UploadAvatarImg)
	return r
}

func CollectRoutes(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.POST("/upload/avatar", controller.UploadAvatarImg)
	r.GET("/user", middleware.AuthMiddleware(), controller.GetUserInfo)

	articleRoutes := r.Group("/article")
	articleController := controller.NewArticleController()
	articleRoutes.POST("", middleware.AuthMiddleware(), articleController.Create)
	articleRoutes.PUT(":id", middleware.AuthMiddleware(), articleController.Update)
	articleRoutes.DELETE(":id", middleware.AuthMiddleware(), articleController.Delete)
	articleRoutes.GET(":id", articleController.Show)
	articleRoutes.POST("list", articleController.List)

	return r
}
