package promptdata

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PromptData(prompt ...any) (string, error) {
	fmt.Println(prompt...)
	data, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", err
	}
	data = strings.TrimSpace(data)
	return data, nil
}

func PrintErr(value any) {
	switch t := value.(type) {
	case string:
		fmt.Println(t)
	case int:
		fmt.Printf("Код ошибки: %d", t)
	case error:
		fmt.Println(t.Error())
	default:
		fmt.Printf("Неизвестный тип ошибки: %d", t)
	}
}
