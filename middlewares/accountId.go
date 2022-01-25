package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AccountId(context *gin.Context) {
	id, err := uuid.Parse(context.Param("accountId"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "uuid not valid"})
		context.Abort()
		return
	}
	context.Set("accountId", id)
}
