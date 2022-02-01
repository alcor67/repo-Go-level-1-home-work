package main

import (
	"fmt"
	"os"

	"github.com/alcor67/repo-Go-level-1-home-work/project-sort-numbers/sortarray"
)

func main() {
	//ввод слайса пользователем
	fmt.Println("введите массив целых чисел:\n")
	arr := sortarray.GetSlice()
	//arr := []int64{-1, 0, 3, 4, 0, 1, 2, 5, 7, -1, 9}
	fmt.Println(arr)
	//arr := []int64{-1, 0, 3, 4, 0, 1, 2, 5, 7, -1}
	fmt.Println("исходный массив:\n", arr)
	fmt.Println("введите направление сортировки массива:\n",
		"'+' сортировка по увеличению значений элементов массива\n",
		"'-' сортировка по уменьшению значений элементов массива\n")
	var direction string
	fmt.Scanln(&direction)
	arrSort, err := sortarray.GetInsertSort(arr, direction)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	switch direction {
	case "+":
		fmt.Println("сортировка по увеличению значений элементов массива:\n", arrSort)
	case "-":
		fmt.Println("сортировка по уменьшению значений элементов массива:\n", arrSort)
	}

}

/*
//функция ввод слайса пользователем
func GetSlice() []int64 {
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
*/
