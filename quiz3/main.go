package main

import (
	"bandProtocol/quiz3/client"
	"bandProtocol/quiz3/models"
	"fmt"
	"time"
)

func main() {
	// create a new client
	c := client.NewClient()

	// create a new BroadcastTransactionRequest
	req := &models.BroadcastTransactionRequest{
		Symbol:    "BTC",
		Price:     10000,
		Timestamp: 1599999999,
	}

	// broadcast the transaction
	resp, err := c.BroadcastTransaction(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	// print the TXHash
	fmt.Println("Transaction hash is ", resp.TXHash)

	// monitor the transaction
	c.MonitorTransaction(resp.TXHash, 10*time.Second)
}
