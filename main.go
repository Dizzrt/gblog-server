package main

import (
	"gblog-server/common"
	"gblog-server/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	common.InitDB()

	r := gin.Default()
	r.StaticFS("/images", http.Dir("./static/images"))
	router.Routers(r)
	panic(r.Run())
}
