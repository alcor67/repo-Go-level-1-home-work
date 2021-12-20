package main

import (
	"fmt"
	"math"
)

func main() {

	var circleArea, circleDiam, circleLength float64
	fmt.Print("Введите площадь круга:")
	fmt.Scan(&circleArea)
	//применяю Println
	fmt.Println("Площадь круга:", circleArea)
	circleDiam = math.Sqrt(circleArea/math.Pi) * 2
	//Применяю Print
	fmt.Print("Диаметр окружности: ", circleDiam, " \n")
	circleLength = circleDiam * math.Pi
	//Применяю Printf %f форматирует число в формате с плавающей точкой
	fmt.Printf("Длина окружности: %f\n", circleLength)
}
