# Transaction Client Module

This module provides two core functionalities for handling blockchain transactions:

1. **BroadcastTransaction**: Send a transaction to the blockchain network.
2. **CheckTransactionStatus**: Check the status of a previously broadcasted transaction.
3. **MonitorTransaction**: Continuously monitor until the status is marked as complete.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [BroadcastTransaction](#broadcasttransaction)
- [CheckTransactionStatus](#checktransactionstatus)
- [MonitorTransaction] (#monitortransaction)
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

_BroadcastTransaction_
