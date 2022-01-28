package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mikolajsemeniuk/Supreme-Go/data"
	"github.com/mikolajsemeniuk/Supreme-Go/entities"
)

func GetAccounts(channel chan []entities.Account) {
	accounts := []entities.Account{}
	data.Context.Find(&accounts)
	channel <- accounts
}

func GetAccount(accountId uuid.UUID, channel chan entities.Account) {
	account := entities.Account{}
	data.Context.Take(&account, accountId)
	channel <- account
}

func AddAccount(account *entities.Account, channel chan error) {
	result := data.Context.Create(&account)
	if result.RowsAffected == 0 {
		channel <- errors.New("error has occured")
	}
	channel <- nil
}

func RemoveAccount(account *entities.Account, channel chan error) {
	result := data.Context.Delete(&account)
	if result.RowsAffected == 0 {
		channel <- errors.New("error has occured")
	}
	channel <- nil
}

func UpdateAccount(accountId uuid.UUID, account *entities.Account, channel chan error) {
	result := data.Context.Save(&account)
	if result.RowsAffected == 0 {
		channel <- errors.New("error has occured")
	}
	channel <- nil
}
