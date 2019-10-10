package service

import (
	"errors"
	"github.com/go-project1/models"
)

// DeleteAccount is used to delete the account
func DeleteAccount(id string) (string, error) {

	var ac []models.Account
	var isAccountDeleted bool

	// take all accounts and range over the loop
	for _, acc := range Accounts {
		// Check if account id is matching or not
		// If account id is not matching, store the account details back to same object
		if acc.ID != id {
			ac = append(ac, acc)
		} else {
			// if account id is matching, mark the isAccountDeleted flag as true
			isAccountDeleted = true
		}
	}
	Accounts = ac

	// if account deleted, then send response to the user, saying account deleted
	if isAccountDeleted {
		return "Account deleted successfully", nil
	}
	return "", errors.New("Account not found")
}

