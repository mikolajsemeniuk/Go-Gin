package application

import (
	"github.com/gin-gonic/gin"
	"github.com/mikolajsemeniuk/Supreme-Go/controllers/account"
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
			accounts.GET("", account.Account.All)
			accounts.POST("", middlewares.AccountBody, account.Account.Add)
			accounts.GET(":accountId", middlewares.AccountId, account.Account.SingleById)
			accounts.DELETE(":accountId", middlewares.AccountId, account.Account.Remove)
			accounts.PATCH(":accountId", middlewares.AccountId, middlewares.AccountBody, account.Account.Update)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run()
}
