package main

import "fmt"

/*
функция может принимать неопределенное количество параметров одного типа
*/
func main() {
	add(1, 2, 3)       // sum = 6
	add(1, 2, 3, 4)    // sum = 10
	add(5, 6, 7, 2, 3) // sum = 23
	/*
		Если мы хотим передать срез,
		то надо указать после аргумента-массива многоточие:
	*/
	add([]int{1, 2, 3}...)
	add([]int{1, 2, 3, 4}...)
	var nums = []int{5, 6, 7, 2, 3}
	add(nums...)

}

func add(numbers ...int) {
	var sum = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println("sum = ", sum)
}
