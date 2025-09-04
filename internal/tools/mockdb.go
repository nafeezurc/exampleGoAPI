package tools

import (
	"time"
)

// create a mockDB type
type mockDB struct{}

// login details data
var mockLoginDetails = map[string]LoginDetails{
	"naf": {
		AuthToken: "123ABC",
		Username:  "naf",
	},
	"jason": {
		AuthToken: "456DEF",
		Username:  "jason",
	},
	"marie": {
		AuthToken: "789GHI",
		Username:  "marie",
	},
}

// coin details data
var mockCoinDetails = map[string]CoinDetails{
	"naf": {
		Coins:    100,
		Username: "naf",
	},
	"jason": {
		Coins:    200,
		Username: "jason",
	},
	"marie": {
		Coins:    300,
		Username: "marie",
	},
}

// in order for this mockDB struct to conform to our database interface, need to create a GetUserLoginDetails, GetCoinDetails, and SetupDatabase from database.go
func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// simulate DB call
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
