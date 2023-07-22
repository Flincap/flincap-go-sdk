# Flincap Go SDK

The Flincap library provides access to the Flincap API from Go.

## Installation
To install see the lines below.

```sh
go get -u github.com/flincap/flincap-go-sdk
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/your-username/flincap-go-sdk"
)

func main() {
	token := "YOUR_AUTH_TOKEN"
	client := flincap.NewFlincapClient(token)

	// Get rate
	rate, err := client.GetRate("USDT", "NGN")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Rate:", rate)
	}

	// Get exchange data
	exchangeData, err := client.GetExchange()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Exchange Data:", exchangeData)
	}

	// Create transaction
	transactionData := map[string]interface{}{
		"selectedCrypt": "USDT",
		"selectedFiat":  "NGN",
		"email":         "test@example.com",
		"bankName":      "Test Bank",
		"bankCode":      "12345",
		"accNum":        "9876543210",
		"accName":       "John Doe",
		"amountFiat":    "5000",
		"amountCrypt":   "100",
		"rate":          "50",
	}

	err = client.CreateTransaction(transactionData)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Transaction created successfully.")
	}

	// Get transaction by ID
	transactionID := "YOUR_TRANSACTION_ID"
	transaction, err := client.GetTransaction(transactionID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Transaction:", transaction)
	}

	// Get transaction history
	transactionType := "DEPOSIT"
	selectedFiat := "CRYPTO"
	history, err := client.GetTransactionHistory(transactionType, selectedFiat)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Transaction History:", history)
	}
}

```

## Beta status

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning the package version to a specific version in your package.json file. This way, you can install the same version each time without breaking changes unless you are intentionally looking for the latest version.

## Contributing

While we value open-source contributions to this SDK, this library is generated programmatically. Additions made directly to this library would have to be moved over to our generation code, otherwise they would be overwritten upon the next generated release. Feel free to open a PR as a proof of concept, but know that we will not be able to merge it as-is. We suggest [opening an issue](https://github.com/flincap/flincap-go-sdk) first to discuss with us!

On the other hand, contributions to the README are always very welcome!