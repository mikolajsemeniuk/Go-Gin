package application

import (
	"github.com/gin-gonic/gin"
	"github.com/mikolajsemeniuk/Supreme-Go/controllers"
	docs "github.com/mikolajsemeniuk/Supreme-Go/docs"
	"github.com/mikolajsemeniuk/Supreme-Go/middlewares"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	router = gin.Default()
)

func Listen() {
	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := router.Group("/api/v1")
	{
		accounts := v1.Group("/accounts")
		{
			accounts.GET("", controllers.GetAccounts)
			accounts.GET(":accountId", middlewares.AccountId, controllers.GetAccount)
		}
	}
	router.POST("/accounts/", middlewares.AccountBody, controllers.AddAccount)
	router.DELETE("/accounts/:accountId", middlewares.AccountId, controllers.RemoveAccount)
	router.PATCH("/accounts/:accountId", middlewares.AccountId, middlewares.AccountBody, controllers.UpdateAccount)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run()
}
