package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"

	"github.com/go-project1/service"
)

func getAccountDetailsHandler(w http.ResponseWriter, r *http.Request) {
	// capture the id passed via URI
	params := mux.Vars(r)

	log.Printf("Started processing get account details, with params %v", params)

	// Send request to service layer
	account := service.GetAccounts(params["id"])

	// encode data
	accJSON, err := json.Marshal(&account)
	if err != nil {
		// if there is any error, write the
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}
	// write the response
	w.WriteHeader(http.StatusOK)
	w.Write(accJSON)
}