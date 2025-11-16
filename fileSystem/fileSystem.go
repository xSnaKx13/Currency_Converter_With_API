package filesystem

import (
	"fmt"
	"os"
)

func WriteFile(fileName string, content []byte) {
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

func ReadFile(fileName string) []byte {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return data
}
