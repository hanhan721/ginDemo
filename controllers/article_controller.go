package controllers

import (
	"encoding/json"
	"errors"
	"ginDemo/global"
	"ginDemo/models"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
)

func CreateArticle(ctx *gin.Context) {
	var article models.Article
	if err := ctx.ShouldBindJSON(&article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := global.Db.AutoMigrate(&article); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := global.Db.Create(&article).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, article)
}

func GetArticles(ctx *gin.Context) {
	var articles []models.Article
	if err := global.Db.Find(&articles).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, articles)
}

func GetArticleById(ctx *gin.Context) {
	id := ctx.Param("id") //获取路径中的参数,类似springboot的 /article/getById/{id}
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id 参数缺失"})
		return
	}
	var article models.Article
	likeKey := "article:id:" + id
	bytes, err := global.Redis.Get(likeKey).Bytes()
	if errors.Is(err, redis.Nil) {
		//查出为空
		if err := global.Db.Where("id=?", id).First(&article).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		serialized, err := json.Marshal(article) //序列化结构体
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		//存入redis
		global.Redis.Set(likeKey, serialized, 0)
	} else if err != nil {
		//redis 异常
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		if err := json.Unmarshal(bytes, &article); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	ctx.JSON(http.StatusOK, article)
}
