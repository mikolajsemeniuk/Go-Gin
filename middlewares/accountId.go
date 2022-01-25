package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AccountId(context *gin.Context) {
	id, err := uuid.Parse(context.Param("accountId"))

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "id not valid"})
		return
	}

	context.Set("accountId", id)
}
