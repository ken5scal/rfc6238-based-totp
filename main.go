package main

import (
	"fmt"
	"os"
	totp "totp/totp-client"
)

var encodedSecret = "CUBUZ3LGDQETIKD7" // example secret, already reset

func main() {
	if otp, err := totp.NewCredential(encodedSecret).GetTOTP(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	} else {
		fmt.Println(otp)
	}
}
