package queue

import (
	"banking-ledger/pkg/config"
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

// TransactionMessage represents the structure of the Kafka message
type TransactionMessage struct {
	AccountID int64   `json:"account_id"`
	Amount    float64 `json:"amount"`
	Type      string  `json:"type"`
}

// Kafka writer instance
var kafkaWriter *kafka.Writer

// Initialize Kafka producer
func InitQueue(cfg *config.Config) {
	kafkaWriter = &kafka.Writer{
		Addr:  kafka.TCP(cfg.KafkaBrokers),
		Topic: "transactions",
	}

	log.Println("[INFO] Kafka producer initialized")
}

// PublishTransaction sends a structured transaction message to Kafka
func PublishTransaction(accountID int64, amount float64, txnType string) error {
	// Create a transaction message struct
	txnMessage := TransactionMessage{
		AccountID: accountID,
		Amount:    amount,
		Type:      txnType,
	}

	// Convert struct to JSON
	messageBytes, err := json.Marshal(txnMessage)
	if err != nil {
		log.Printf("[ERROR] Failed to serialize transaction message: %v", err)
		return err
	}

	// Publish to Kafka
	err = kafkaWriter.WriteMessages(context.Background(),
		kafka.Message{
			Value: messageBytes,
		},
	)
	if err != nil {
		log.Printf("[ERROR] Failed to publish transaction: %v", err)
		return err
	}

	log.Printf("[INFO] Transaction published to Kafka: %s", string(messageBytes))
	return nil
}

func CloseQueue() {
	if kafkaWriter != nil {
		kafkaWriter.Close()
		log.Println("[INFO] Kafka producer closed")
	}
}
