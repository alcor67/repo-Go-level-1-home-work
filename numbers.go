package main

import (
	"fmt"

	"strconv"
)

//функция валидации данных
func isValid(numberString string) bool {
	//конвертация строки в число метод strconv.Atoi()
	number, err := strconv.Atoi(numberString)
	if err != nil {
		fmt.Println(" input is not digit")
		return false
	} else if number < 100 {
		fmt.Println("number < 100")
		return false
	} else if number > 999 {
		fmt.Println("number > 999")
		return false
	} else {
		//fmt.Println("100 =< number <= 999")
		return true
	}
}

func main() {

	//Ввод положительного 3-х значного числа
	var numberString string
	fmt.Println("Введите положительное 3-х значное число: ")
	fmt.Scanln(&numberString)
	//валидация данных
	if isValid(numberString) {
		fmt.Println("    число сотен:", string(numberString[0]), "\n",
			"число десятков:", string(numberString[1]), "\n",
			"  число единиц:", string(numberString[2]), "\n")
	}
}
