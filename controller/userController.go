package controller

import (
	"gblog-server/common"
	"gblog-server/model"
	"net/http"
	"strconv"

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

func GetUserInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"id":     user.(model.User).ID,
			"avatar": user.(model.User).Avatar,
		},
		"msg": "get user info successfully!",
	})
}

func GetBriefInfo(ctx *gin.Context) {
	db := common.GetDB()

	userID := ctx.Params.ByName("id")
	user, _ := ctx.Get("user")

	var curUser model.User
	if userID == strconv.Itoa(int(user.(model.User).ID)) {
		curUser = user.(model.User)
	} else {
		db.Where("id = ?", userID).First(&curUser)
		if curUser.ID == 0 {
			common.Fail(ctx, nil, "用户不存在")
			return
		}
	}

	common.Success(ctx, gin.H{
		"id":      curUser.ID,
		"name":    curUser.UserName,
		"avatar":  curUser.Avatar,
		"loginID": user.(model.User).ID,
	}, "success")
}

func GetUserArticleList(ctx *gin.Context) {
	db := common.GetDB()

	user, _ := ctx.Get("user")
	var articles []model.ArticleInfo
	db.Table("articles").Select("id, title, LEFT(content,80) As content, head_image, created_at").Where("user_id = ?", user.(model.User).ID).Order("created_at desc").Find(&articles)

	common.Success(ctx, gin.H{
		"id":       user.(model.User).ID,
		"name":     user.(model.User).UserName,
		"avatar":   user.(model.User).Avatar,
		"articles": articles,
	}, "success")
}

func ModifyAvatar(ctx *gin.Context) {
	db := common.GetDB()
	user, _ := ctx.Get("user")

	var requestUser model.User
	ctx.Bind(&requestUser)
	avatar := requestUser.Avatar

	var curUser model.User
	db.Where("id = ?", user.(model.User).ID).First(&curUser)

	if err := db.Model(&curUser).Update("avatar", avatar).Error; err != nil {
		common.Fail(ctx, nil, "更新失败")
		return
	}
	common.Success(ctx, nil, "更新成功")
}

func ModifyName(ctx *gin.Context) {
	db := common.GetDB()

	user, _ := ctx.Get("user")
	var requestUser model.User
	ctx.Bind(&requestUser)
	userName := requestUser.UserName

	var curUser model.User
	db.Where("id = ?", user.(model.User).ID).First(&curUser)
	if err := db.Model(&curUser).Update("user_name", userName).Error; err != nil {
		common.Fail(ctx, nil, "更新失败")
		return
	}
	common.Success(ctx, nil, "更新成功")
}
