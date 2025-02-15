package transaction

import (
	"banking-ledger/internal/account"
	"banking-ledger/pkg/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeposit(t *testing.T) {
	database.InitPostgres(&database.Config{PostgresURL: "postgres://user:password@localhost:5432/test_db?sslmode=disable"})

	acc, _ := account.CreateAccount("Charlie", 1000)

	err := Deposit(acc.ID, 500)
	assert.Nil(t, err)

	updatedAcc, _ := account.GetAccount(acc.ID)
	assert.Equal(t, 1500.0, updatedAcc.Balance)
}

func TestWithdraw(t *testing.T) {
	acc, _ := account.CreateAccount("David", 1000)

	err := Withdraw(acc.ID, 500)
	assert.Nil(t, err)

	updatedAcc, _ := account.GetAccount(acc.ID)
	assert.Equal(t, 500.0, updatedAcc.Balance)
}
