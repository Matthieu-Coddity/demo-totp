package main

import (
	"encoding/base32"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gosuri/uilive"
	"github.com/jltorresm/otpgo"
	"github.com/jltorresm/otpgo/config"
)

var account *otpgo.TOTP

func newOTP() *otpgo.TOTP {
	secretString := []byte("a dirty litte secret")
	secret := base32.StdEncoding.EncodeToString(secretString)

	return &otpgo.TOTP{
		Key:       secret,
		Period:    30,
		Delay:     2,
		Algorithm: config.HmacSHA1,
		Length:    0,
	}
}

func generateQRCode() {

	keyUri := account.KeyUri("votre compte", "Etabli")
	uri := keyUri.String()
	qrCode, err := keyUri.QRCode()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("************URI**************\n%s\n*****************************\n\n\n************QRCode**************\n%s\n*****************************", uri, qrCode)
}

func main() {
	account = newOTP()
	if len(os.Args) != 1 {
		generateQRCode()
		os.Exit(1)
	}

	writer := uilive.New()
	writer.Start()
	for {
		token, err := account.Generate()
		if err != nil {
			log.Fatalf(err.Error())
		}
		fmt.Fprintf(writer, "code : %s\n", token)
		time.Sleep(500 * time.Millisecond)
	}
	writer.Stop()
}
