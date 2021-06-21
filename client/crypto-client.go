package client

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-projects/crypto-cli/model"
)

func FetchCrypto(fiat, crypto string) (string, error) {
	URL := "https://api.nomics.com/v1/currencies/ticker?key=f2138e28c176868fd55d93a5233e6498fd32ad7e&interval=1d&ids=" + crypto + "&convert=" + fiat

	resp, err := http.Get(URL)
	if err != nil {
		log.Fatalln("Something went Wrong,Try Again")
	}

	defer resp.Body.Close()

	var cResp model.CryptoResponse

	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		log.Fatal("OOPS..went Wrong")
	}
	return cResp.TextOutput(), nil
}
