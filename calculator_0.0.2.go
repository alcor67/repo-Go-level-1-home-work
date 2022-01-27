package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"strconv"
)

var op = flag.String("op", "", "operation")
var x = flag.Float64("x", 0, "first operand")
var y = flag.Float64("y", 0, "seconfd operand")
var result float64
var err error

func main() {
	flag.Parse()

	if *op == "" && *x == 0 && *y == 0 {

		// ввод двух значений и оператора
		fmt.Print("введите через пробел 2 значения и оператор\n",
			"{значение_1 оператор значение_2}\n")
		var numberString1, numberString2 string
		var operator string
		var number1, number2 float64

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
		x = &number1
		y = &number2
		op = &operator
		/*
			fmt.Println("значение 1: ", number1)
			fmt.Println("  оператор: ", operator)
			fmt.Println("значение 2: ", number2)
		*/
	}

	//fmt.Println(*op, *x, *y)

	result, err = calc(*op, *x, *y)

	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf(" результат: %v %s %v = %.6g \n", *x, *op, *y, result)
}

func calc(op string, x, y float64) (float64, error) {

	switch op {
	case "+":
		//применяю  чтобы избавиться от цифрового хвоста в значении результата
		//кстати, сам догадался
		return x + y, nil
	case "-":
		return x - y, nil
	case "*":
		return x * y, nil
	case "/":
		//todo: проверка деления на ноль
		if y == 0 {
			//сообщение об ошибке
			//fmt.Println("type of error: division by zero")
			return 0, errors.New("can not divide by zero")
		} else {
			return x / y, nil
		}

	case "%":
		fmt.Println("операция взятия остатка от деления")
		fmt.Print("данные и результат будут приведены к целочисленному типу\n",
			"путем отбрасывания дробной части\n")
		if y < 1 {
			//сообщение об ошибке
			//fmt.Println("type of error: integer divide by zero")
			return 0, errors.New("can not divide by zero")
		} else {
			return float64(int(x) % int(y)), nil
		}
	default:
		//сообщение об ошибке
		return 0, errors.New("operation not supported")

	}

}
