package router

import (
	"ginDemo/controllers"
	"ginDemo/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}
	api := r.Group("/api")
	api.GET("/exchangeRates", controllers.GetExchangeRates)
	api.Use(middlewares.AuthMiddleware())
	{
		api.POST("/exchangeRate/create", controllers.CreateExchangeRate)
		api.GET("/exchangeRate/getById", controllers.GetExchangeRateById)
	}
	api.GET("/articles", controllers.GetArticles)
	{
		api.POST("/article/create", controllers.CreateArticle)
		api.GET("/article/getById/:id", controllers.GetArticleById)
	}
	return r
}