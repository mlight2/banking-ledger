package transaction

import (
	"banking-ledger/pkg/database"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func Withdraw(accountID int64, amount float64) error {
	tx, err := database.DB.Beginx()
	if err != nil {
		return err
	}

	var balance float64
	query := fmt.Sprintf("SELECT balance FROM accounts WHERE id = %d FOR UPDATE", accountID)
	err = tx.Get(&balance, query)
	if err != nil {
		tx.Rollback()
		if err == sql.ErrNoRows {
			return errors.New("account not found")
		}
		return err
	}

	if balance < amount {
		tx.Rollback()
		return errors.New("insufficient funds")
	}

	query = fmt.Sprintf("UPDATE accounts SET balance = balance - %f WHERE id = %d", amount, accountID)
	_, err = tx.Exec(query)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	_, err = database.MongoDB.Collection("ledger").InsertOne(context.TODO(), map[string]interface{}{
		"account_id": accountID,
		"amount":     amount,
		"type":       "withdraw",
	})
	if err != nil {
		return errors.New("withdrawal recorded in SQL but failed in MongoDB")
	}

	return nil
}

func Deposit(accountID int64, amount float64) error {
	tx, err := database.DB.Beginx()
	if err != nil {
		return err
	}

	var balance float64
	query := fmt.Sprintf("SELECT balance FROM accounts WHERE id = %d FOR UPDATE", accountID)
	err = tx.Get(&balance, query)
	if err != nil {
		tx.Rollback()
		if err == sql.ErrNoRows {
			return errors.New("account not found")
		}
		return err
	}

	// Update balance
	query = fmt.Sprintf("UPDATE accounts SET balance = balance + %f WHERE id = %d", amount, accountID)
	_, err = tx.Exec(query)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	_, err = database.MongoDB.Collection("ledger").InsertOne(context.TODO(), map[string]interface{}{
		"account_id": accountID,
		"amount":     amount,
		"type":       "deposit",
	})
	if err != nil {
		return errors.New("deposit recorded in SQL but failed in MongoDB")
	}

	return nil
}
