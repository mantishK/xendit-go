package balance_test

import (
	"fmt"
	"log"

	"github.com/mantishK/xendit-go"
	"github.com/mantishK/xendit-go/balance"
)

func ExampleGet() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := balance.GetParams{
		AccountType: "CASH",
	}

	resp, err := balance.Get(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("balance: %+v\n", resp)
}
