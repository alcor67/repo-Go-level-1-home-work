package sortarray

import "errors"

//универсальная функция сортировки по значениям элементов массива
func GetInsertSort(arr []int64, direction string) ([]int64, error) {

	for i := range arr {
		for j := i + 1; j > 0 && j < len(arr); j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	switch direction {

	case "+":
		return arr, nil

	case "-":
		for i := 0; i <= (len(arr)-1)/2; i++ {
			arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
		}
		return arr, nil

	default:
		return arr, errors.New("неизвестная операция сортировки массива: " + direction)
	}

}
