package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikolajsemeniuk/Supreme-Go/inputs"
)

func AccountBody(context *gin.Context) {
	var input inputs.Account

	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		context.Abort()
		return
	}
	context.Set("accountInput", input)
}
