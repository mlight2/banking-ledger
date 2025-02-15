package ledger

import (
	"banking-ledger/pkg/database"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Transaction struct
type Transaction struct {
	AccountID int64     `bson:"account_id"`
	Amount    float64   `bson:"amount"`
	Type      string    `bson:"type"`
	Timestamp time.Time `bson:"timestamp"`
}

func SaveTransaction(accountID int64, amount float64, txnType string) error {
	collection := database.MongoDB.Collection("ledger")
	txn := Transaction{
		AccountID: accountID,
		Amount:    amount,
		Type:      txnType,
		Timestamp: time.Now(),
	}

	_, err := collection.InsertOne(context.TODO(), txn)
	if err != nil {
		log.Printf("[ERROR] Failed to save transaction (AccountID: %d, Amount: %.2f, Type: %s): %v", accountID, amount, txnType, err)
		return err
	}

	log.Printf("[INFO] Transaction saved successfully (AccountID: %d, Amount: %.2f, Type: %s)", accountID, amount, txnType)
	return nil
}

func GetTransactionHistory(accountID int64) ([]Transaction, error) {
	collection := database.MongoDB.Collection("ledger")

	filter := bson.M{"account_id": accountID}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Printf("[ERROR] Failed to fetch transaction history (AccountID: %d): %v", accountID, err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var transactions []Transaction
	for cursor.Next(context.TODO()) {
		var txn Transaction
		if err := cursor.Decode(&txn); err != nil {
			log.Printf("[ERROR] Failed to decode transaction record (AccountID: %d): %v", accountID, err)
			return nil, err
		}
		transactions = append(transactions, txn)
	}

	log.Printf("[INFO] Retrieved %d transactions for AccountID: %d", len(transactions), accountID)
	return transactions, nil
}
