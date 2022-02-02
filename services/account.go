package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mikolajsemeniuk/Supreme-Go/data"
	"github.com/mikolajsemeniuk/Supreme-Go/entities"
)

var (
	Account IAccountService = &accountService{}
)

type accountService struct{}

type IAccountService interface {
	All(channel chan []entities.Account)
	SingleById(accountId uuid.UUID, channel chan entities.Account)
	Add(account *entities.Account, channel chan error)
	Remove(account *entities.Account, channel chan error)
	Update(accountId uuid.UUID, account *entities.Account, channel chan error)
}

func (a *accountService) All(channel chan []entities.Account) {
	accounts := []entities.Account{}
	data.Context.Find(&accounts)
	channel <- accounts
}

func (a *accountService) SingleById(accountId uuid.UUID, channel chan entities.Account) {
	account := entities.Account{}
	data.Context.Take(&account, accountId)
	channel <- account
}

func (a *accountService) Add(account *entities.Account, channel chan error) {
	result := data.Context.Create(&account)
	if result.RowsAffected == 0 {
		channel <- errors.New("error has occured")
	}
	channel <- nil
}

func (a *accountService) Remove(account *entities.Account, channel chan error) {
	result := data.Context.Delete(&account)
	if result.RowsAffected == 0 {
		channel <- errors.New("error has occured")
	}
	channel <- nil
}

func (a *accountService) Update(accountId uuid.UUID, account *entities.Account, channel chan error) {
	result := data.Context.Save(&account)
	if result.RowsAffected == 0 {
		channel <- errors.New("error has occured")
	}
	channel <- nil
}
