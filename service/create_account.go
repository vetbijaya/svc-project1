package service

import (
	"math/rand"
	"time"
	"github.com/go-project1/models"
)
const (
	charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var (
	// Accounts holds all the application account details
	Accounts []models.Account

	seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
)
//CreateAccount is used to create the account
func CreateAccount(account models.Account) string {
	// generate the id for an account, with length of 15 digit
account.ID = genString(15)
// store the values
Accounts = append(Accounts, account)
// Send account id in the response, if account stored successfully
return account.ID
}
// genString will generate the alphanumeric id for an account
func genString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}