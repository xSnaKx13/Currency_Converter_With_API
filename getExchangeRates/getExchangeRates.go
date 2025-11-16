package getexchangerates

import (
	promptdata "Currency_Converter/promptData"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ExchangeRates struct {
	Result string             `json:"result"`
	Base   string             `json:"base_code"`
	Date   string             `json:"time_last_update_utc"`
	Rates  map[string]float64 `json:"rates"`
}

func GetExchangeRates() (ExchangeRates, error) {
	baseCurrency, err := promptdata.PromptData("Введите базовую валюту: ")
	if err != nil {
		// добавить текст ошибки
		return ExchangeRates{}, err
	}
	url := fmt.Sprintf("https://open.er-api.com/v6/latest/%s", baseCurrency)
	resp, err := http.Get(url)
	if err != nil {
		// добавить текст ошибки
		return ExchangeRates{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// добавить текст ошибки
		return ExchangeRates{}, err
	}
	var rates ExchangeRates
	err = json.Unmarshal(body, &rates)
	if err != nil {
		// добавить текст ошибки
		return ExchangeRates{}, err
	}
	return rates, nil
}
