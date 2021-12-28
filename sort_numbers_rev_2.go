package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//ввод слайса пользователем
	fmt.Println("введите массив целых чисел:\n")
	arr := getSlice()
	fmt.Println("исходный массив:\n", arr)
	fmt.Println("введите направление сортировки массива:\n",
		"'+' сортировка по увеличению значений элементов массива\n",
		"'-' сортировка по уменьшению значений элементов массива\n")
	var direction string
	fmt.Scanln(&direction)
	switch direction {
	case "+":
		getInsertSortIncrease(arr)
		fmt.Println("сортировка по увеличению значений элементов массива:\n", arr)
	case "-":
		getInsertSortDecrease(arr)
		fmt.Println("сортировка по уменьшению значений элементов массива:\n", arr)
	default:
		fmt.Println("type of error: unknown operation")
	}
}

//функция ввод слайса пользователем
func getSlice() []int64 {
	inputNums := []int64{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			break
		}
		inputNums = append(inputNums, num)
	}
	return inputNums
}

//функция сортировки по увеличению значений элементов массива
func getInsertSortIncrease(arr []int64) {
	for i := range arr {
		for j := i + 1; j > 0 && j < len(arr); j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

//функция сортировки по уменьшению значений элементов массива
func getInsertSortDecrease(arr []int64) {

	for i := range arr {
		for j := i + 1; j > 0 && j < len(arr); j-- {
			if arr[j] > arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}
