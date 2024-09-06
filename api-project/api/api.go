package main

import (
	"encoding/json"
	"net/http"
)

// Coint Balance Params
type CoinBalanceParams struct {
	Username string
}

// Coin Balance Response
type CoinBalanceResponse struct {
	// Success Code ,Ussually 200
	Code int 

	// Account Balance
	Balance int
}

// ERRor response 
type ErrorResponse struct {

	// Error Code
	Code int

	// Error Message
	Message string
}


