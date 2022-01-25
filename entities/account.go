package entities

import "github.com/google/uuid"

type Account struct {
	Entity
	Code  string `json:"code"`
	Price uint   `json:"price"`
}

type AccountRepository interface {
	GetAccounts() []Account
	GetAccount(accountId uuid.UUID) *Account
	AddAccount(account *Account) error
}
