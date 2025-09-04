package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/nafeezurc/exampleGoAPI/api"
	"github.com/nafeezurc/exampleGoAPI/internal/tools"
	log "github.com/sirupsen/logrus"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	// we assume the call already ran through the authorization middleware, so we just need to grab the username from the parameters
	// the most "go-like" way to do this is to decode our parameters into our CoinBalanceParams struct
	var params = api.CoinBalanceParams{}

	// can use gorilla/schema package for decoder
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	// this line will grab the parameters from the url and set them to the parameters in the struct
	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// instantiate a database
	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	// call get user coins method
	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// set value to response struct
	var response = api.CoinBalanceResponse{
		Balance: (*tokenDetails).Coins,
		Code:    http.StatusOK,
	}

	// write it to response writer
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
