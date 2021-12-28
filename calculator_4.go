package main

import (
	"fmt"

	"strconv"
)

func main() {

	// ввод двух значений и оператора
	fmt.Print("введите через пробел 2 значения и оператор\n",
		"{значение_1 оператор значение_2}\n")
	var numberString1, numberString2 string
	var operator string
	var number1, number2 float64
	var result float32
	fmt.Scanln(&numberString1, &operator, &numberString2)
	//конвертация строки в число метод strconv.ParseFloat(string, 64)
	number1, err := strconv.ParseFloat(numberString1, 64)
	//обработка ошибки
	if err != nil {
		//сообщение об ошибке
		fmt.Println(err.Error())
		return
	}

	//конвертация строки в число метод strconv.ParseFloat(string, 64)
	number2, err1 := strconv.ParseFloat(numberString2, 64)
	//обработка ошибки
	if err1 != nil {
		//сообщение об ошибке
		fmt.Println(err.Error())
		return
	}

	fmt.Println("значение 1: ", number1)
	fmt.Println("  оператор: ", operator)
	fmt.Println("значение 2: ", number2)
	switch operator {
	case "+":
		//применяю float32 чтобы избавиться от цифрового хвоста в значении результата
		//кстати, сам догадался
		result = float32(number1 + number2)
	case "-":
		result = float32(number1 - number2)
	case "*":
		result = float32(number1 * number2)
	case "/":
		//todo: проверка деления на ноль
		if number2 == 0 {
			//сообщение об ошибке
			fmt.Println("type of error: division by zero")
			return
		} else {
			result = float32(number1 / number2)
		}

	case "%":
		fmt.Println("операция взятия остатка от деления")
		fmt.Print("данные и результат будут приведены к целочисленному типу\n",
			"путем отбрасывания дробной части\n")
		if number2 < 1 {
			//сообщение об ошибке
			fmt.Println("type of error: integer divide by zero")
			return
		} else {
			result = float32(int(number1) % int(number2))
		}
	default:
		//сообщение об ошибке
		fmt.Println("type of error: unknown operation")
	}
	fmt.Println(" результат: ", result)
}
