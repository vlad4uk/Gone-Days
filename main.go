package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/vlad4uk/Gone-Days/internal/days"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите дату в формате DD.MM.YYYY: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	dateStr := strings.TrimSpace(input)

	daysPassed, err := days.CountDaysFrom(dateStr)

	if err != nil {
		fmt.Printf("Ошибка обработки даты: %v\n", err)
		return
	}

	fmt.Printf("Прошло %d дней с %s\n", daysPassed, dateStr)
}
