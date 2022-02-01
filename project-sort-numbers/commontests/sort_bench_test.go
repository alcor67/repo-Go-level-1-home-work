package sortarray_test

import (
	"fmt"
	"sort"
	"testing"
)

func BenchmarkSampleSort(b *testing.B) {
	b.SetBytes(2)
	b.ReportAllocs()

	arr := []int64{-4894, 626, 6426, 7076, 7684, 8402, -2730, -7750, -4948, -4222, -9220, -8500, 4494, -3174, 6368, -5592, -6602, -9856, -5190}

	arrIncr := arr
	arrDecr := arr

	b.ResetTimer() //Теперь начальная настройка не будет учтена в результатах
	for i := 0; i < b.N; i++ {

		sort.Slice(arrIncr, func(i, j int) bool { return arrIncr[i] < arrIncr[j] })
		sort.Slice(arrDecr, func(i, j int) bool { return arrDecr[i] > arrDecr[j] })
	}
	fmt.Printf("\n===========sort method============\n")
}
