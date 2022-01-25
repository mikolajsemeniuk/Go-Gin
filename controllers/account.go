package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/mikolajsemeniuk/Supreme-Go/entities"
	"github.com/mikolajsemeniuk/Supreme-Go/inputs"
	"github.com/mikolajsemeniuk/Supreme-Go/services"
)

func GetAccounts(context *gin.Context) {
	accounts := services.GetAccounts()

	context.JSON(http.StatusOK, gin.H{"message": accounts})
}

func GetAccount(context *gin.Context) {
	id := context.MustGet("accountId").(uuid.UUID)
	account := services.GetAccount(id)

	if account.Id == uuid.Nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": account})
}

func AddAccount(context *gin.Context) {
	account := entities.Account{}
	input := context.MustGet("accountInput").(inputs.Account)

	copier.Copy(&account, &input)

	if err := services.AddAccount(&account); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "record created"})
}

func UpdateAccount(context *gin.Context) {
	id := context.MustGet("accountId").(uuid.UUID)
	input := context.MustGet("accountInput").(inputs.Account)
	account := services.GetAccount(id)

	if account.Id == uuid.Nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		return
	}

	copier.Copy(&account, &input)

	if err := services.UpdateAccount(id, account); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
}

func RemoveAccount(context *gin.Context) {
	id := context.MustGet("accountId").(uuid.UUID)
	account := services.GetAccount(id)

	if account.Id == uuid.Nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		return
	}

	if err := services.RemoveAccount(account); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
}
