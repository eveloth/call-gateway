package models

import (
	"net/url"

	money "github.com/Rhymond/go-money"
)

type User struct {
	Id           int
	Email        *email
	PasswordHash []byte
	Role         string
	ApiKeys      []ApiKey
	Webhook      url.URL
	Balance      money.Amount
	Transactions []Transaction
}
