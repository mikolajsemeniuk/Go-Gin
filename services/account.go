package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mikolajsemeniuk/Supreme-Go/data"
	"github.com/mikolajsemeniuk/Supreme-Go/entities"
)

var (
	Account IAccountService = &AccountService{}
)

type AccountService struct{}

type IAccountService interface {
	GetAccounts(channel chan []entities.Account)
	GetAccount(accountId uuid.UUID, channel chan entities.Account)
	AddAccount(account *entities.Account, channel chan error)
	RemoveAccount(account *entities.Account, channel chan error)
	UpdateAccount(accountId uuid.UUID, account *entities.Account, channel chan error)
}

func (*AccountService) GetAccounts(channel chan []entities.Account) {
	accounts := []entities.Account{}
	data.Context.Find(&accounts)
	channel <- accounts
}

func (*AccountService) GetAccount(accountId uuid.UUID, channel chan entities.Account) {
	account := entities.Account{}
	data.Context.Take(&account, accountId)
	channel <- account
}

func (*AccountService) AddAccount(account *entities.Account, channel chan error) {
	result := data.Context.Create(&account)
	if result.RowsAffected == 0 {
		channel <- errors.New("error has occured")
	}
	channel <- nil
}

func (*AccountService) RemoveAccount(account *entities.Account, channel chan error) {
	result := data.Context.Delete(&account)
	if result.RowsAffected == 0 {
		channel <- errors.New("error has occured")
	}
	channel <- nil
}

func (*AccountService) UpdateAccount(accountId uuid.UUID, account *entities.Account, channel chan error) {
	result := data.Context.Save(&account)
	if result.RowsAffected == 0 {
		channel <- errors.New("error has occured")
	}
	channel <- nil
}
