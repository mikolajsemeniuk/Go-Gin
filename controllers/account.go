package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/mikolajsemeniuk/Supreme-Go/entities"
	"github.com/mikolajsemeniuk/Supreme-Go/inputs"
	"github.com/mikolajsemeniuk/Supreme-Go/payloads"
	"github.com/mikolajsemeniuk/Supreme-Go/services"
)

// @BasePath /api/v1
// @Summary get all accounts
// @Schemes
// @Description get all accounts
// @Tags accounts
// @Accept json
// @Produce json
// @Success 200 {array} entities.Account
// @Router /accounts [get]
func GetAccounts(context *gin.Context) {
	channel := make(chan []entities.Account)
	go services.GetAccounts(channel)

	context.JSON(http.StatusOK, payloads.Ok{Data: <-channel})
}

// @BasePath /api/v1
// @Summary add account
// @Schemes
// @Description add account
// @Tags accounts
// @Accept json
// @Produce json
// @Success 200 {object} entities.Account
// @Failure 400 {object} payloads.BadRequest
// @Router /accounts [post]
func AddAccount(context *gin.Context) {
	account := entities.Account{}
	input := context.MustGet("accountInput").(inputs.Account)
	channel := make(chan error)

	copier.Copy(&account, &input)

	go services.AddAccount(&account, channel)
	if err := <-channel; err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, payloads.BadRequest{Message: err.Error()})
		return
	}

	// location := url.URL{Path: fmt.Sprintf("/accounts/%s", account.Id)}
	// context.Redirect(http.StatusFound, location.RequestURI())
	context.JSON(http.StatusOK, payloads.Ok{Data: account})
}

// @BasePath /api/v1
// @Summary get account by id
// @Schemes
// @Description get account by id
// @Tags accounts
// @Accept json
// @Produce json
// @Param accountId path string true "Account ID"
// @Success 200 {object} entities.Account
// @Failure 400 {object} payloads.BadRequest
// @Failure 404 {object} payloads.NotFound
// @Router /accounts/{accountId} [get]
func GetAccount(context *gin.Context) {
	accountId := context.MustGet("accountId").(uuid.UUID)
	channel := make(chan entities.Account)

	go services.GetAccount(accountId, channel)
	account := <-channel

	if account.Id == uuid.Nil {
		context.AbortWithStatusJSON(http.StatusNotFound, payloads.NotFound{Message: "Account not found"})
		return
	}

	context.JSON(http.StatusOK, payloads.Ok{Data: account})
}

// @BasePath /api/v1
// @Summary remove account
// @Schemes
// @Description remove account
// @Tags accounts
// @Accept json
// @Produce json
// @Param accountId path string true "Account ID"
// @Success 200 {object} entities.Account
// @Failure 400 {object} payloads.BadRequest
// @Failure 404 {object} payloads.NotFound
// @Router /accounts/{accountId} [delete]
func RemoveAccount(context *gin.Context) {
	accountId := context.MustGet("accountId").(uuid.UUID)
	accountChannel := make(chan entities.Account)
	errorChannel := make(chan error)

	go services.GetAccount(accountId, accountChannel)
	account := <-accountChannel

	if account.Id == uuid.Nil {
		context.AbortWithStatusJSON(http.StatusNotFound, payloads.NotFound{Message: "Account not found"})
		return
	}

	go services.RemoveAccount(&account, errorChannel)
	if err := <-errorChannel; err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, payloads.BadRequest{Message: err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, payloads.NoContent{Message: "Account removed"})
}

// @BasePath /api/v1
// @Summary update account
// @Schemes
// @Description update account
// @Tags accounts
// @Accept json
// @Produce json
// @Param accountId path string true "Account ID"
// @Success 200 {object} entities.Account
// @Failure 400 {object} payloads.BadRequest
// @Failure 404 {object} payloads.NotFound
// @Router /accounts/{accountId} [patch]
func UpdateAccount(context *gin.Context) {
	accountId := context.MustGet("accountId").(uuid.UUID)
	input := context.MustGet("accountInput").(inputs.Account)
	accountChannel := make(chan entities.Account)
	errorChannel := make(chan error)

	go services.GetAccount(accountId, accountChannel)
	account := <-accountChannel

	if account.Id == uuid.Nil {
		context.AbortWithStatusJSON(http.StatusNotFound, payloads.NotFound{Message: "Account not found"})
		return
	}

	copier.Copy(&account, &input)

	go services.UpdateAccount(accountId, &account, errorChannel)
	if err := <-errorChannel; err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, payloads.BadRequest{Message: err.Error()})
		return
	}

	context.JSON(http.StatusOK, payloads.Ok{Data: account})
}
