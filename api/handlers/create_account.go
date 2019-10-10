package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	domain "github.com/go-project1/models"
	"github.com/go-project1/service"
)

// create handler receives the api request to create an acount
func createHandler(w http.ResponseWriter, r *http.Request) {

	// read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	// Decode the body and store the details in variable account
	var account domain.Account

	// Unmarshal the body into variable a
	err = json.Unmarshal(body, &account)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte ("Bad request"))
		return
	}
	// send details to service from API
	account.ID = service.CreateAccount(account)

	// encoding the data to JSON
	encodedData, err := encodeToJSON(account)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// account created successfully
	w.WriteHeader(http.StatusCreated)
	w.Write(encodedData)
}