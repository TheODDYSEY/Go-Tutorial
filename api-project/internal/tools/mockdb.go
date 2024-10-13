package tools

import (
    "time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
    "alex": {
        AuthToken: "1234",
        Username:  "alex",
    },
    "john": {
        AuthToken: "5678",
        Username:  "john",
    },
    "mary": {
        AuthToken: "91011",
        Username:  "mary",
    },
    "jane": {
        AuthToken: "121314",
        Username:  "jane",
    },
    "peter": {
        AuthToken: "151617",
        Username:  "peter",
    },
    "lisa": {
        AuthToken: "181920",
        Username:  "lisa",
    },
}

var mockCoinDetails = map[string]CoinDetails{
    "alex": {
        Coins:    100,
        Username: "alex",
    },
    "john": {
        Coins:    200,
        Username: "john",
    },
    "mary": {
        Coins:    300,
        Username: "mary",
    },
    "jane": {
        Coins:    400,
        Username: "jane",
    },
    "peter": {
        Coins:    500,
        Username: "peter",
    },
    "lisa": {
        Coins:    600,
        Username: "lisa",
    },
}

func (db *mockDB) GetUserLoginDetails(username string) *LoginDetails {
    // simulate DB Call
    time.Sleep(1 * time.Second)

    clientData, ok := mockLoginDetails[username]
    if !ok {
        return nil
    }
    return &clientData
}

func (db *mockDB) GetUserCoins(username string) *CoinDetails {
    // simulate DB Call
    time.Sleep(1 * time.Second)

    clientData, ok := mockCoinDetails[username]
    if !ok {
        return nil
    }
    return &clientData
}

func (db *mockDB) SetupDatabase() error {
    return nil
}