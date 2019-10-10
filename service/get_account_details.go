package service

import (
	"github.com/go-project1/models"
)

//GetAccounts is used to get the account details
func GetAccounts(id string) models.Account {
	// take all accounts and iterate over the for loop
	for i := range Accounts {
		// Check if id passed by the user is matching with any record
		if id == Accounts[i].ID {
			// If account id matched, send the details to user
			return Accounts[i]
		}
	}
	return models.Account{}
}
