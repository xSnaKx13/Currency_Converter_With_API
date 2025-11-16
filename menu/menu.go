package menu

import (
	"Currency_Converter/convert"
	filesystem "Currency_Converter/fileSystem"
	getexchangerates "Currency_Converter/getExchangeRates"
	printallrates "Currency_Converter/printAllRates"
	promptdata "Currency_Converter/promptData"
	"fmt"
)

func Menu() {
	prompt, err := promptdata.PromptData("--Конвертер валют--\n",
		"1 - Показать курс валют онлайн\n",
		"2 - Выгрузить курс валют и записать в файл\n",
		"3 - Конвертировать\n",

		"0 - Выход",
	)
	if err != nil {
		promptdata.PrintErr(err)
		return
	}
	switch prompt {
	case "1":
		printallrates.PrintAllRates()
	case "2":
		retes, err := getexchangerates.GetExchangeRates()
		if err != nil {
			promptdata.PrintErr(err)
		}
		filesystem.WriteInFile(retes)
	case "3":
		result := convert.ConvertCurrency()
		fmt.Printf("Результат: %.2f\n", result)
	case "0":
		return
	default:
		promptdata.PrintErr("Неверный ввод")
	}
}
