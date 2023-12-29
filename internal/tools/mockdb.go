package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"javi": {
		AuthToken: "123ABC",
		Username:  "javi",
	},
	"guillermo": {
		AuthToken: "456DEF",
		Username:  "guillermo",
	},
	"ines": {
		AuthToken: "789GHI",
		Username:  "ines",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"javi": {
		Coins:    100,
		Username: "javi",
	},
	"guillermo": {
		Coins:    200,
		Username: "guillermo",
	},
	"ines": {
		Coins:    300,
		Username: "ines",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	clientData := LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	clientData := CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
