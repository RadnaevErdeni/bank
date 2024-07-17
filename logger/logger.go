package logger

import (
	"log"
	"time"
)

func LogOperation(operation string, accountID int, amount float64) {
	log.Printf("[%s] Account ID: %d, Amount: %.2f, Time: %s\n", operation, accountID, amount, time.Now().Format(time.RFC3339))
}
