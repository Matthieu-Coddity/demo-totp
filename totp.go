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
		Length:    6,
	}
}

func generateQRCode() {

	keyUri := account.KeyUri("debbie.harry@coddity.com", "demoT-OTP")
	uri := keyUri.String()
	qrCode, err := keyUri.QRCode()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("************URI**************\n%s\n*****************************\n\n\n************QRCode**************\n%s\n*****************************", uri, qrCode)
}

func run() {
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

func verify(code string) {
	access, err := account.Validate(code)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if access {
		fmt.Println("login in!")
	} else {
		fmt.Println("access denied")
	}
}

func main() {
	account = newOTP()
	if len(os.Args) == 1 || len(os.Args) > 3 {
		fmt.Println("Usage : generate / run / verify <CODE_2_VERIFY>")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "generate":
		generateQRCode()
	case "run":
		run()
	case "verify":
		if len(os.Args) != 3 {
			fmt.Println("need code to verify")
			os.Exit(1)
		}
		verify(os.Args[2])
	default:
		fmt.Println("Usage : generate / run / verify <CODE_2_VERIFY>")
		os.Exit(1)
	}
}
