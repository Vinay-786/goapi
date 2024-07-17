package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"Vinay": {
		AuthToken: "23123",
		Username:  "Vinay",
	},
	"Aanjney": {
		AuthToken: "53231",
		Username:  "Aanjney",
	},
	"Himanshu": {
		AuthToken: "54235",
		Username:  "Himanshu",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"Vinay": {
		Coins:    299,
		Username: "Vinay",
	},
	"Aanjney": {
		Coins:    239,
		Username: "Aanjney",
	},
	"Himanshu": {
		Coins:    290,
		Username: "Himanshu",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	//DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	//DB call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
