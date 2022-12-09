package controller

import (
	"gblog-server/common"
	"gblog-server/model"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ArticleController struct {
	DB *gorm.DB
}

type IArticleController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Show(ctx *gin.Context)
	List(ctx *gin.Context)
}

func (this ArticleController) Create(ctx *gin.Context) {
	var articleRequest model.CreateArticleRequest

	if err := ctx.ShouldBindJSON(&articleRequest); err != nil {
		common.Fail(ctx, nil, "data error")
		return
	}

	user, _ := ctx.Get("user")
	article := model.Article{
		ID:        uuid.NewV4(),
		UserId:    user.(model.User).ID,
		Title:     articleRequest.Title,
		Content:   articleRequest.Content,
		HeadImage: articleRequest.HeadImage,
	}

	if err := this.DB.Create(&article).Error; err != nil {
		common.Fail(ctx, nil, "发布失败")
		return
	}
	common.Success(ctx, gin.H{"id": article.ID}, "发布成功")
}

func (this ArticleController) Update(ctx *gin.Context) {
	var articleRequest model.CreateArticleRequest

	if err := ctx.ShouldBindJSON(&articleRequest); err != nil {
		common.Fail(ctx, nil, "data error")
		return
	}

	articleID := ctx.Params.ByName("id")
	var article model.Article
	// FIXME not found
	if this.DB.Where("id = ?", articleID).First(&article) == nil {
		common.Fail(ctx, nil, "文章不存在")
		return
	}

	user, _ := ctx.Get("user")
	userId := user.(model.User).ID

	if userId != article.UserId {
		common.Fail(ctx, nil, "权限不足")
		return
	}

	if err := this.DB.Model(&article).Update("content", articleRequest.Content).Error; err != nil {
		common.Fail(ctx, nil, "修改失败")
		return
	}
	common.Success(ctx, nil, "修改成功")
}

func (this ArticleController) Delete(ctx *gin.Context) {
	articleID := ctx.Params.ByName("id")

	var article model.Article
	// FIXME not found
	if this.DB.Where("id = ?", articleID).First(&article) == nil {
		common.Fail(ctx, nil, "文章不存在")
		return
	}

	user, _ := ctx.Get("user")
	userID := user.(model.User).ID
	if userID != article.UserId {
		common.Fail(ctx, nil, "权限不足")
		return
	}

	if err := this.DB.Delete(&article).Error; err != nil {
		common.Fail(ctx, nil, "删除失败")
		return
	}
	common.Success(ctx, nil, "删除成功")
}

func (this ArticleController) Show(ctx *gin.Context) {
	articleID := ctx.Params.ByName("id")

	var article model.Article
	// FIXME not found
	if this.DB.Where("id = ?", articleID).First(&article) == nil {
		common.Fail(ctx, nil, "文章不存在")
		return
	}
	common.Success(ctx, gin.H{"artical": article}, "success")
}

func (this ArticleController) List(ctx *gin.Context) {
	// TODO
}

func NewArticleController() IArticleController {
	db := common.GetDB()
	db.AutoMigrate(model.Article{})
	return ArticleController{DB: db}
}
