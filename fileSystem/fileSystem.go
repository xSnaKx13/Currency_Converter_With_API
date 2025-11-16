package filesystem

import (
	promptdata "Currency_Converter/promptData"
	tobyte "Currency_Converter/toByte"
	"fmt"
	"os"
)

func writeFile(fileName string, content []byte) {
	data, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	_, err = data.Write(content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Запись успешна!")
	defer data.Close()
}

func readFile(fileName string) []byte {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return data
}

func deleteFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		panic(err)
	}
}

func renameFile(oldName string, newName string) {
	err := os.Rename(oldName, newName)
	if err != nil {
		panic(err)
	}
}

func WriteInFile(content any) {
	fileName, err := promptdata.PromptData("Введите название файла: ")
	if err != nil {
		promptdata.PrintErr(err)
	}
	data, err := tobyte.ToByte(content)
	if err != nil {
		promptdata.PrintErr(err)
	}
	writeFile(fileName, data)
}
