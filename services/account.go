package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mikolajsemeniuk/Supreme-Go/data"
	"github.com/mikolajsemeniuk/Supreme-Go/entities"
)

func GetAccounts() []entities.Account {
	accounts := []entities.Account{}
	data.Context.Find(&accounts)
	return accounts
}

func GetAccount(accountId uuid.UUID) *entities.Account {
	account := entities.Account{}
	data.Context.Take(&account, accountId)
	return &account
}

func AddAccount(account *entities.Account) error {
	result := data.Context.Create(&account)
	if result.RowsAffected == 0 {
		return errors.New("error has occured")
	}
	return nil
}

func RemoveAccount(account *entities.Account) error {
	result := data.Context.Delete(&account)
	if result.RowsAffected == 0 {
		return errors.New("error has occured")
	}
	return nil
}

func UpdateAccount(accountId uuid.UUID, account *entities.Account) error {
	result := data.Context.Save(&account)
	if result.RowsAffected == 0 {
		return errors.New("error has occured")
	}
	return nil
}
