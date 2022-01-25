package application

import (
	"github.com/gin-gonic/gin"
	"github.com/mikolajsemeniuk/Supreme-Go/controllers"
	docs "github.com/mikolajsemeniuk/Supreme-Go/docs"
	"github.com/mikolajsemeniuk/Supreme-Go/middlewares"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	BasePath = "/api/v1"
)

var (
	router = gin.Default()
)

func Listen() {
	docs.SwaggerInfo.BasePath = BasePath
	v1 := router.Group(BasePath)
	{
		accounts := v1.Group("/accounts")
		{
			accounts.GET("", controllers.GetAccounts)
			accounts.POST("", middlewares.AccountBody, controllers.AddAccount)
			accounts.GET(":accountId", middlewares.AccountId, controllers.GetAccount)
			accounts.DELETE(":accountId", middlewares.AccountId, controllers.RemoveAccount)
			accounts.PATCH(":accountId", middlewares.AccountId, middlewares.AccountBody, controllers.UpdateAccount)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run()
}
