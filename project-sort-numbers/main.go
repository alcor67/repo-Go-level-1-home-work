package main

import (
	"fmt"
	"os"

	"github.com/alcor67/repo-Go-level-1-home-work/project-sort-numbers/sortarray"
)

func main() {
	
	fmt.Println("введите массив целых чисел:\n")
	arr := sortarray.GetSlice()
	
	fmt.Println(arr)
	
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
