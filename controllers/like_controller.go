package controllers

import (
	"errors"
	"ginDemo/global"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
)

func LikeArticle(ctx *gin.Context) {
	articleId := ctx.Param("id")

	likeKey := "article:id:" + articleId + ":likes"
	if err := global.Redis.Incr(likeKey).Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "点赞成功"})
}

func GetArticleLikes(ctx *gin.Context) {
	articleId := ctx.Param("id")

	likeKey := "article:id:" + articleId + ":likes"
	likes, err := global.Redis.Get(likeKey).Result()
	if errors.Is(err, redis.Nil) {
		likes = "0"
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"likes": likes})
}
