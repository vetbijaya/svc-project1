package handlers

import (
	"encoding/json"
	"github.com/go-project1/models"
)

func encodeToJSON(account models.Account) ([]byte, error){
//marshal account details and return
return json.Marshal(account)
}

