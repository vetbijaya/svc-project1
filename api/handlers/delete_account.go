package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/go-project1/service"
)

func deleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	// capture parameters passed via URL
	params := mux.Vars(r)

	resp, err := service.DeleteAccount(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([] byte("Account not found"))
		return
	}
w.WriteHeader(http.StatusNotFound)
w.Write([]byte(resp))
}