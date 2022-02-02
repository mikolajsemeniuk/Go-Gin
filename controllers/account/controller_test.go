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

type serviceMock struct{}

func (serviceMock) All(channel chan []entities.Account) {
	channel <- []entities.Account{}
}
func (serviceMock) SingleById(accountId uuid.UUID, channel chan entities.Account)             {}
func (serviceMock) Add(account *entities.Account, channel chan error)                         {}
func (serviceMock) Remove(account *entities.Account, channel chan error)                      {}
func (serviceMock) Update(accountId uuid.UUID, account *entities.Account, channel chan error) {}

func Test_should_return_accounts_with_status_code_200(t *testing.T) {
	// Arrange
	expected := []entities.Account{}
	actual := []entities.Account{}

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
