package service

import (
	"github.com/go-project1/models"
)

// UpdateAccount is used to update the account details
func UpdateAccount(account models.Account) {

	var ac []models.Account

	// Read all accounts stored in accounts object
	for _, acc := range Accounts {

// Check for the account to be updated
if acc.ID == account.ID {

	// apply name changes
	if acc.Name != "" {
		acc.Name = account.Name
	}
	// apply email changes
	if acc.Email != "" {
		acc.Email = account.Email
	}
	//apply phone number changes
	if acc.PhoneNumber != "" {
		acc.PhoneNumber = account.PhoneNumber
	}
}

// store the updated account detail
ac = append(ac, acc)
}
//update account map
Accounts = ac
}