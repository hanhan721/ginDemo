package controllers

import (
	"ginDemo/global"
	"ginDemo/models"
	"ginDemo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	hashPwd, err := utils.HashPassword(user.Password) //加密密码
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	user.Password = hashPwd

	token, err2 := utils.GenerateJWT(user.Username)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err2.Error()})
		return
	}

	err3 := global.Db.AutoMigrate(&user) //如果不存在会自动创建user表
	if err3 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err3.Error()})
		return
	}

	if err := global.Db.Create(&user).Error; err != nil { //插入数据
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func Login(ctx *gin.Context) {
	//定义入参结构体
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	var user models.User
	if err := global.Db.Where("username=?", input.Username).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "用户名有误!"})
		return
	}
	if !utils.CheckPassword(input.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "密码或用户名有误!"})
		return
	}
	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})

}
