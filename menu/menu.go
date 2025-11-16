package menu

import (
	"Currency_Converter/convert"
	printallrates "Currency_Converter/printAllRates"
	promptdata "Currency_Converter/promptData"
	"fmt"
)

func Menu() {
	prompt, err := promptdata.PromptData("--Конвертер валют--\n",
		"1 - Показать курс валют\n",
		"2 - Выгрузить курс валют в файл\n",
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
		// сделать общую функцию в отдельном пакете под запись в файл
	case "3":
		result := convert.ConvertCurrency()
		fmt.Printf("Результат: %.2f\n", result)
	case "0":
		return
	default:
		promptdata.PrintErr("Неверный ввод")
	}
}
