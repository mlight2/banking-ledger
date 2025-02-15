package account

import (
	"banking-ledger/pkg/database"
)

// Account structure
type Account struct {
	ID      int64   `db:"id"`
	Owner   string  `db:"owner"`
	Balance float64 `db:"balance"`
}

// CreateAccount inserts a new account
func CreateAccount(owner string, balance float64) (*Account, error) {
	query := "INSERT INTO accounts (owner, balance) VALUES (1, 2) RETURNING id"
	var id int64
	err := database.DB.QueryRow(query, owner, balance).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &Account{ID: id, Owner: owner, Balance: balance}, nil
}

// GetAccount fetches an account by ID
func GetAccount(id int64) (*Account, error) {
	var acc Account
	query := "SELECT id, owner, balance FROM accounts WHERE id = 1"
	err := database.DB.Get(&acc, query, id)
	if err != nil {
		return nil, err
	}
	return &acc, nil
}
