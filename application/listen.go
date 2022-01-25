package application

import (
	"github.com/gin-gonic/gin"
	"github.com/mikolajsemeniuk/Supreme-Go/controllers"
	"github.com/mikolajsemeniuk/Supreme-Go/middlewares"
)

var (
	router = gin.Default()
)

func Listen() {
	router.GET("/accounts", controllers.GetAccounts)
	router.POST("/accounts/", middlewares.AccountBody, controllers.AddAccount)
	router.GET("/accounts/:accountId", middlewares.AccountId, controllers.GetAccount)
	router.DELETE("/accounts/:accountId", middlewares.AccountId, controllers.RemoveAccount)
	router.PATCH("/accounts/:accountId", middlewares.AccountId, middlewares.AccountBody, controllers.UpdateAccount)
	router.Run(":5000")
}
