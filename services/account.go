package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mikolajsemeniuk/Supreme-Go/data"
	"github.com/mikolajsemeniuk/Supreme-Go/entities"
	"gorm.io/gorm"
)

var (
	Account IAccountService = &accountService{
		context: data.Context,
	}
)

type accountService struct {
	context *gorm.DB
}

type IAccountService interface {
	All(channel chan []entities.Account)
	SingleById(accountId uuid.UUID, channel chan entities.Account)
	Add(account *entities.Account, channel chan error)
	Remove(account *entities.Account, channel chan error)
	Update(accountId uuid.UUID, account *entities.Account, channel chan error)
}

func (a *accountService) All(channel chan []entities.Account) {
	accounts := []entities.Account{}
	a.context.Find(&accounts)
	channel <- accounts
}

func (a *accountService) SingleById(accountId uuid.UUID, channel chan entities.Account) {
	account := entities.Account{}
	a.context.Take(&account, accountId)
	channel <- account
}

func (a *accountService) Add(account *entities.Account, channel chan error) {
	result := a.context.Create(&account)
	if result.RowsAffected == 0 {
		channel <- errors.New("error has occured")
	}
	channel <- nil
}

func (a *accountService) Remove(account *entities.Account, channel chan error) {
	result := a.context.Delete(&account)
	if result.RowsAffected == 0 {
		channel <- errors.New("error has occured")
	}
	channel <- nil
}

func (a *accountService) Update(accountId uuid.UUID, account *entities.Account, channel chan error) {
	result := a.context.Save(&account)
	if result.RowsAffected == 0 {
		channel <- errors.New("error has occured")
	}
	channel <- nil
}
