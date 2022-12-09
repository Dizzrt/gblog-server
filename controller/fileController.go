package controller

import (
	"io"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadAvatarImg(ctx *gin.Context) {
	img, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "格式错误",
		})
		return
	}

	imgName := header.Filename
	ext := path.Ext(imgName)

	name := "avatar_" + time.Now().Format("20060102150405")
	newFileName := name + ext
	out, err := os.Create("static/images/avatars/" + newFileName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "server error",
		})
		return
	}

	defer out.Close()
	_, err = io.Copy(out, img)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": "500",
			"msg":  "server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"filePath": "/images/avatars/" + newFileName},
		"msg":  "ok",
	})
}
