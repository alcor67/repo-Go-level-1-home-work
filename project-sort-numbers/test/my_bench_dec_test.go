package sortarray_test

import (
	"fmt"
	"testing"

	"github.com/alcor67/repo-Go-level-1-home-work/project-sort-numbers/sortarray"
)

func BenchmarkSampleMainDec(b *testing.B) {
	b.SetBytes(2)
	b.ReportAllocs()

	arr := []int64{-4894, 626, 6426, 7076, 7684, 8402, -2730, -7750, -4948, -4222, -9220, -8500, 4494, -3174, 6368, -5592, -6602, -9856, -5190}

	b.ResetTimer() //Теперь начальная настройка не будет учтена в результатах
	for i := 0; i < b.N; i++ {

		direction := "-"
		_, err := sortarray.GetInsertSort(arr, direction)
		if err != nil {
			b.Fatalf("arry %+v expected to return no error, but received: %s", arr, err)

		}

	}
	fmt.Printf("\n ===========my method============\n")
}
