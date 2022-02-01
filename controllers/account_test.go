package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/mikolajsemeniuk/Supreme-Go/entities"
	"github.com/mikolajsemeniuk/Supreme-Go/mocks/services"
)

func Test_should_return_accounts_with_status_code_200(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	channel := make(chan []entities.Account)
	// accounts := []entities.Account{}

	mockService := services.NewMockIAccountService(controller)
	mockService.EXPECT().All(channel).Return()

	accountController := &account{
		service: mockService,
	}

	router := gin.Default()
	router.GET("api/v1/accounts", accountController.All)

	request, _ := http.NewRequest(http.MethodGet, "/api/v1/accounts", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Error(fmt.Sprintf("Status code should be 200 but was %d", recorder.Code))
	}
}
