package client_test

import (
	"fmt"
	"log"

	"github.com/mantishK/xendit-go/client"
	"github.com/mantishK/xendit-go/invoice"
)

// This example shows the usage of Xendit's API client
func Example() {
	api := client.New("examplesecretkey")

	data := invoice.GetParams{
		ID: "invoice-id",
	}

	resp, err := api.Invoice.Get(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("retrieved invoice: %+v\n", resp)
}
