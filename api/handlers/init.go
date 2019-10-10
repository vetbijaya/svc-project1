package handlers

import (
	"net/http"

	"github.com/spf13/viper"

	"github.com/gorilla/mux"

	"github.com/go-project1/config"
	domain "github.com/go-project1/models"
)

// acc is used to hold the account details
var acc []domain.Account

// Init hold all the exposed api's by this service
func Init() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/create", createHandler).Methods(http.MethodPost)
	r.HandleFunc("/get/{id}", getAccountDetailsHandler).Methods(http.MethodGet)
	r.HandleFunc("/update/{id}", updateAccountHandler).Methods(http.MethodPut)
	r.HandleFunc("/delete/{id}", deleteAccountHandler).Methods(http.MethodDelete)

	http.ListenAndServe(viper.GetString(config.PortNumber), r)
	return r
}
