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
	getInsertSort(arr, direction)
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

//универсальная функция сортировки по значениям элементов массива
func getInsertSort(arr []int64, direction string) {

	for i := range arr {
		for j := i + 1; j > 0 && j < len(arr); j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	switch direction {
	case "+":
		fmt.Println("сортировка по увеличению значений элементов массива:\n", arr)
	case "-":
		for i := 0; i <= (len(arr)-1)/2; i++ {
			arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
		}
		fmt.Println("сортировка по уменьшению значений элементов массива:\n", arr)
	default:
		fmt.Println("type of error: unknown operation")
	}

}
