package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikolajsemeniuk/Supreme-Go/inputs"
	"github.com/shyandsy/ShyGinErrors"
)

func AccountBody(context *gin.Context) {
	var input inputs.Account
	errors := ShyGinErrors.NewShyGinErrors(inputs.ErrorMessages)

	if err := context.BindJSON(&input); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errors.ListAllErrors(input, err)})
		return
	}

	context.Set("accountInput", input)
}
