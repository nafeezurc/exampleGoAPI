package tools

import (
	log "github.com/sirupsen/logrus"
)

// database collections

// contains authtoken for validating request
type LoginDetails struct {
	AuthToken string
	Username  string
}

// contains coin balance
type CoinDetails struct {
	Coins    int64
	Username string
}

// define a few methods required for our api
// using an interface because we can swap between our databases really easily
type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

// function that returns above interface
func NewDatabase() (*DatabaseInterface, error) {
	// create a database variable and set it to a mockDB struct, which we'll create to set up our interface
	var database DatabaseInterface = &mockDB{}

	// use the setup database method and then do standard error checks
	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}
