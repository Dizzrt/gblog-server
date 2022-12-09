package router

import (
	"gblog-server/controller"
	"gblog-server/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.POST("/upload/avatar", controller.UploadAvatarImg)

	articleRoutes := r.Group("/article")
	articleController := controller.NewArticleController()
	articleRoutes.POST("", middleware.AuthMiddleware(), articleController.Create)
	articleRoutes.PUT(":id", middleware.AuthMiddleware(), articleController.Update)
	articleRoutes.DELETE(":id", middleware.AuthMiddleware(), articleController.Delete)
	articleRoutes.GET(":id", articleController.Show)
	articleRoutes.POST("list", articleController.List)

	userRoutes := r.Group("/user")
	userRoutes.Use(middleware.AuthMiddleware())
	userRoutes.GET("", controller.GetUserInfo)
	userRoutes.GET("briefInfo/:id", controller.GetBriefInfo)
	userRoutes.GET("articles/:id", controller.GetUserArticleList)
	userRoutes.PUT("avatar/:id", controller.ModifyAvatar)
	userRoutes.PUT("name/:id", controller.ModifyName)

	return r
}
