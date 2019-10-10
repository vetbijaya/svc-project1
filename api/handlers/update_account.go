package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	domain "github.com/go-project1/models"
	"github.com/go-project1/service"
)

func updateAccountHandler(w http.ResponseWriter, r *http.Request) {
	// capture the id passed via URI
	params := mux.Vars(r)

	//read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}
	var account domain.Account

	//Unmarshal the body into variable a
	err = json.Unmarshal(body, &account)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}
	// Store the account details into accounts object
	account.ID = params["id"]

	// Send request to service layer
	service.UpdateAccount(account)

	w.Write([]byte("Account updated successfully"))
}
