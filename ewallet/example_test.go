package ewallet_test

import (
	"fmt"
	"log"

	"github.com/mantishK/xendit-go"
	"github.com/mantishK/xendit-go/ewallet"
)

func ExampleCreatePayment() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := ewallet.CreatePaymentParams{
		ExternalID:  "dana-ewallet",
		Amount:      20000,
		Phone:       "08123123123",
		EWalletType: xendit.EWalletTypeDANA,
		CallbackURL: "mystore.com/callback",
		RedirectURL: "mystore.com/redirect",
	}

	resp, err := ewallet.CreatePayment(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created payment: %+v\n", resp)
}

func ExampleGetPaymentStatus() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := ewallet.GetPaymentStatusParams{
		ExternalID:  "data-ewallet",
		EWalletType: xendit.EWalletTypeDANA,
	}

	resp, err := ewallet.GetPaymentStatus(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("retrieved payment: %+v\n", resp)
}
