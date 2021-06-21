package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/go-projects/crypto-cli/client"
)

func main() {
	fiatCurrency := flag.String("fiat", "INR", "The name of the fiat currency")
	cryptoCurrency := flag.String("crypto", "BTC", "The name of the cryptoCurrency")

	flag.Parse()

	crypto, err := client.FetchCrypto(*fiatCurrency, *cryptoCurrency)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(crypto)
}
