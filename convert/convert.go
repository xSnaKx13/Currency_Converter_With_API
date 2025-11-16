package convert

import (
	getexchangerates "Currency_Converter/getExchangeRates"
	promptdata "Currency_Converter/promptData"
	"fmt"
)

func ConvertCurrency() float64 {
	rate, err := getexchangerates.GetExchangeRates()
	if err != nil {
		promptdata.PrintErr(err)
	}
	to, err := promptdata.PromptData("Введите нужную валюту: ")
	if err != nil {
		promptdata.PrintErr(err)
	}
	fmt.Println("Введите сумму: ")
	var amount float64
	fmt.Scan(&amount)
	userCurrency := rate.Rates[to]
	return amount * userCurrency

}
