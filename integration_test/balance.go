package main

import (
	"fmt"
	"log"

	"github.com/mantishK/xendit-go/balance"
)

func balanceTest() {
	getData := balance.GetParams{
		AccountType: "CASH",
	}
	_, err := balance.Get(&getData)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Balance integration tests done!")
}
