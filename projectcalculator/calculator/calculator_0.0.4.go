package calculator

import (
	"errors"
	"fmt"
	"math"
)

func Calc2(op string, x, y float64) (float64, error) {
	return x + y, nil
}

func Calc(op string, x, y float64) (float64, error) {

	switch op {
	case "+":
		return x + y, nil
		//return x + y + 1, fmt.Errorf("plus operation is broken")

	case "-":
		return x - y, nil
	case "*":
		return x * y, nil

	case "**":
		result := math.Pow(x, y)
		//fmt.Println(result)
		if math.IsNaN(result) {
			return result, errors.New("result is NaN")
		}

		//fmt.Println(result)
		if math.IsInf(result, 0) {
			return result, errors.New("result is Inf")
		}

		return result, nil

	case "/":

		//проверка деления на ноль
		if y == 0 {
			//return 0, errors.New("can not divide by zero")
			return 0, fmt.Errorf("can not divide by zero")
		} else {
			return x / y, nil
		}

	case "%":
		/*
			fmt.Println("операция взятия остатка от деления")
			fmt.Print("данные и результат будут приведены к целочисленному типу\n",
				"путем отбрасывания дробной части\n")
		*/
		if y < 1 {
			return 0, errors.New("can not divide by zero")
		} else {
			return float64(int(x) % int(y)), nil
		}
	default:
		return 0, errors.New("operation not supported")
	}
}
