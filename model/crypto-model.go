package model

import "fmt"

type CryptoResponse []struct {
	Name              string `json:"name"`
	Price             string `json:"price"`
	Rank              string `json:"rank"`
	High              string `json:"high"`
	CirculatingSupply string `json:"circulating_supply"`
}

func (c CryptoResponse) TextOutput() string {
	return fmt.Sprintf(
		"Name: %s\nPrice: %s\nRank: %s\nHigh: %s\nCirculating Supply: %s\n",
		c[0].Name, c[0].Price, c[0].Rank, c[0].High, c[0].CirculatingSupply,
	)
}
