package controller

import (
	"gblog-server/common"
	"gblog-server/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	db := common.GetDB()

	var requestUser model.User
	ctx.Bind(&requestUser)

	userName := requestUser.UserName
	password := requestUser.Password
	email := requestUser.Email

	var user model.User
	db.Where("email = ?", email).First(&user)
	if user.ID != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 422,
			"msg":  "该邮箱已被注册",
		})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	newUser := model.User{
		UserName:  userName,
		Password:  string(hashedPassword),
		Email:     email,
		Avatar:    "/images/avatar.png",
		Collects:  model.Array{},
		Following: model.Array{},
		Fans:      0,
	}

	db.Create(&newUser)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}

func Login(ctx *gin.Context) {
	db := common.GetDB()

	var requestUser model.User
	ctx.Bind(&requestUser)

	email := requestUser.Email
	password := requestUser.Password

	var user model.User
	db.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 422,
			"msg":  "用户不存在",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 422,
			"msg":  "密码错误",
		})
		return
	}

	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "登陆成功",
	})

}

func GetInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"id": user.(model.User).ID},
		"msg":  "get info successfully!",
	})
}
