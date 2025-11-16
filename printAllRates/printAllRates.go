package printallrates

import (
	getexchangerates "Currency_Converter/getExchangeRates"
	promptdata "Currency_Converter/promptData"
	"fmt"
)

func PrintAllRates() {
	rates, err := getexchangerates.GetExchangeRates()
	if err != nil {
		promptdata.PrintErr(err)
		return
	}
	fmt.Println(rates)
}
