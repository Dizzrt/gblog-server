package router

import (
	"gblog-server/controller"
	"gblog-server/middleware"

	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) *gin.Engine {
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.GET("/user", middleware.AuthMiddleware(), controller.GetInfo)

	return r
}
