package account

import (
	"banking-ledger/pkg/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	database.InitPostgres(&database.Config{PostgresURL: "postgres://user:password@localhost:5432/test_db?sslmode=disable"})

	acc, err := CreateAccount("Alice", 1000)
	assert.Nil(t, err)
	assert.NotNil(t, acc)
	assert.Equal(t, "Alice", acc.Owner)
	assert.Equal(t, 1000.0, acc.Balance)
}

func TestGetAccount(t *testing.T) {
	acc, _ := CreateAccount("Bob", 500)

	// Fetch account details
	retrievedAcc, err := GetAccount(acc.ID)
	assert.Nil(t, err)
	assert.Equal(t, "Bob", retrievedAcc.Owner)
	assert.Equal(t, 500.0, retrievedAcc.Balance)
}
