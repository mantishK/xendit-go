package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mantishK/xendit-go"
	"github.com/mantishK/xendit-go/ewallet"
)

func ewalletTest() {
	createPaymentData := ewallet.CreatePaymentParams{
		ExternalID:  "dana-" + time.Now().String(),
		Amount:      20000,
		Phone:       "08123123123",
		EWalletType: xendit.EWalletTypeDANA,
		CallbackURL: "mystore.com/callback",
		RedirectURL: "mystore.com/redirect",
	}

	resp, err := ewallet.CreatePayment(&createPaymentData)
	if err != nil {
		log.Panic(err)
	}
	getPaymentStatusData := ewallet.GetPaymentStatusParams{
		ExternalID:  resp.ExternalID,
		EWalletType: resp.EWalletType,
	}

	_, err = ewallet.GetPaymentStatus(&getPaymentStatusData)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("EWallet integration tests done!")
}
