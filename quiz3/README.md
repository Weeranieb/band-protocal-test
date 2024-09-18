# Transaction Client Module

This module provides two core functionalities for handling blockchain transactions:

1. **BroadcastTransaction**: Send a transaction to the blockchain network.
2. **CheckTransactionStatus**: Check the status of a previously broadcasted transaction.
<!-- 3. **MonitorTransaction**: Continuously monitor until the status is marked as complete. -->

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [BroadcastTransaction](#broadcasttransaction)
- [CheckTransactionStatus](#checktransactionstatus)
- [Transaction Status Handling](#transaction-status-handling)
- [Example Script](#example-script)

## Installation

To use this module, you will need Go 1.16+ installed. You can install the required dependencies by running:

```bash
go get github.com/weeranieb/band-protocal-test/quiz3

```

## Usage

**Initialization**

To begin, you need to initialize the `Client` by providing the HTTP client and the base URL for broadcasting and checking transaction statuses.

```go
import (
    "net/http"
    "github.com/weeranieb/band-protocal-test/quiz3/models"
    "github.com/weeranieb/band-protocal-test/quiz3/client"
)

func main() {
    httpClient := &http.Client{}
    c := client.Client{
        HTTPClient: httpClient,
        Broadcast:  "https://api.blockchain.com/v1/broadcast",
        Status:     "https://api.blockchain.com/v1/status",
    }
}
```

### BroadcastTransaction

The `BroadcastTransaction` method allows you to send a transaction to the blockchain. It accepts a `BroadcastTransactionRequest` object and returns a `BroadcastTransactionResponse` with the transaction hash.

```go
func (c *Client) BroadcastTransaction(req *models.BroadcastTransactionRequest) (*models.BroadcastTransactionResponse, error)
```

### Request Structure

```go
type BroadcastTransactionRequest struct {
	Symbol    string `json:"symbol"`
	Price     uint64 `json:"price"`
	Timestamp uint64 `json:"timestamp"`
}
```

### Response Structure

```go
type BroadcastTransactionResponse struct {
	TXHash string `json:"tx_hash"`
}
```

---

## CheckTransactionStatus

The `CheckTransactionStatus` method allows you to check the status of a previously broadcasted transaction using its hash.

```go
func (c *Client) CheckTransactionStatus(txHash string) (*models.TransactionStatusResponse, error)
```

### Response Structure

```go
type MonitorTransactionResponse struct {
	TXStatus string `json:"tx_status"`
}
```

---

### Transaction Status Handling

Handling transaction status is crucial for ensuring that your transaction has been successfully processed on the blockchain. The following strategies can be employed for each status:

- Does Not Exist:

  - The client should verify if the transaction was broadcasted successfully.
  - If the transaction does not exist, a new broadcast should be created to generate a new hash and reattempt the transaction.

- Pending:

  - The transaction has been broadcasted but is not yet confirmed.
  - It’s recommended to poll the transaction status periodically using CheckTransactionStatus.

- Confirmed:

  - The transaction has been successfully mined and included in a block.
  - At this point, the transaction can be considered complete.

- Failed:
  - The transaction did not succeed.
  - In case of failure, the application should log the error and take corrective action.

## Example Script

Here is an example script that demonstrates how to broadcast a transaction and check its status with interval time 10 seconds:

```go
package main

import (
    "log"
    "net/http"
    "time"
    "your-module-path/models"
    "your-module-path/client"
)

func main() {
    // Initialize the client
    httpClient := &http.Client{}
    c := client.Client{
        HTTPClient: httpClient,
        Broadcast:  "https://api.blockchain.com/v1/broadcast",
        Status:     "https://api.blockchain.com/v1/status",
    }

    // Broadcast a transaction
    broadcastReq := &models.BroadcastTransactionRequest{
        // Fill in the request details
    }

    resp, err := c.BroadcastTransaction(broadcastReq)
    if err != nil {
        log.Fatalf("Failed to broadcast transaction: %v", err)
    }

    log.Printf("Transaction broadcasted with hash: %s\n", resp.TXHash)

    // Poll for transaction status
    for {
        statusResp, err := c.CheckTransactionStatus(resp.TXHash)
        if err != nil {
            log.Fatalf("Failed to check transaction status: %v", err)
        }

        log.Printf("Transaction status: %s\n", statusResp.Status)

        if statusResp.Status == "confirmed" {
            log.Println("Transaction confirmed!")
            break
        } else if statusResp.Status == "failed" {
            log.Println("Transaction failed!")
            break
        }

        // Wait before checking the status again
        time.Sleep(10 * time.Second)
    }
}
```
