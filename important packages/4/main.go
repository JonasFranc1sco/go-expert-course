package main

import (
	"encoding/json"
	"os"
)

type Account struct { // tags
	Num     int `json:"n"`
	Balance int `json:"b"`
}

func main() {
	account := Account{Num: 1, Balance: 200}
	res, err := json.Marshal(account)
	if err != nil {
		println(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		println(err)
	}

	pureJson := []byte(`{"n":2,"b":300}`)
	var accountX Account

	json.Unmarshal(pureJson, &accountX)
	println(accountX.Balance)
}
