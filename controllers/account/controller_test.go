package account

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mikolajsemeniuk/Supreme-Go/entities"
	"github.com/mikolajsemeniuk/Supreme-Go/services"
)

var allMock func(channel chan []entities.Account)
var singleById func(accountId uuid.UUID, channel chan entities.Account)
var add func(account *entities.Account, channel chan error)
var remove func(account *entities.Account, channel chan error)
var update func(accountId uuid.UUID, account *entities.Account, channel chan error)

type serviceMock struct{}

func (serviceMock) All(channel chan []entities.Account) {
	allMock(channel)
}
func (serviceMock) SingleById(accountId uuid.UUID, channel chan entities.Account) {
	singleById(accountId, channel)
}
func (serviceMock) Add(account *entities.Account, channel chan error) {
	add(account, channel)
}
func (serviceMock) Remove(account *entities.Account, channel chan error) {
	remove(account, channel)
}
func (serviceMock) Update(accountId uuid.UUID, account *entities.Account, channel chan error) {
	update(accountId, account, channel)
}

func Test_should_return_accounts_with_status_code_200(t *testing.T) {
	// Arrange
	expected := []entities.Account{}
	var actual []entities.Account

	allMock = func(channel chan []entities.Account) {
		channel <- expected
	}
	services.Account = serviceMock{}
	accountController := &account{}

	router := gin.Default()
	router.GET("api/v1/accounts", accountController.All)

	// Act
	request, _ := http.NewRequest(http.MethodGet, "/api/v1/accounts", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	json.Unmarshal(recorder.Body.Bytes(), &actual)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error(fmt.Sprintf("Status code should be 200 but was %d", recorder.Code))
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Error(fmt.Sprintf("Structs are not equal actual %v, expected %v", actual, expected))
	}
}
