package sortarray_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/alcor67/repo-Go-level-1-home-work/project-sort-numbers/sortarray"
)

func TestSortSmoke(t *testing.T) {
	arr := []int64{-1, 2, -3, 4}
	direction := "+"
	Expected, err := sortarray.GetInsertSort(arr, direction)
	if err != nil {
		//t.Errorf("arry {-1, 2, -3, 4} expected to return no error, but received: %s", err)
		t.Errorf("arry %+v expected to return no error, but received: %s", arr, err)

	}
	expected := []int64{-3, -1, 2, 4}
	for i := 0; i < len(arr); i++ {
		if Expected[i] != expected[i] {
			t.Errorf("результат %+v не совпалает с проверкой %+v", Expected[i], expected[i])
		}
	}
}

func ExampleSortArry() {
	arr := []int64{-1, 2, -3, 4}
	direction := "+"
	Expected, err := sortarray.GetInsertSort(arr, direction)
	fmt.Println(Expected, err)
	// Output: [-3 -1 2 4] <nil>
}

func TestSortTableOk(t *testing.T) {
	type call struct {
		Direction string
		Arr       []int64
		Expected  []int64
	}
	calls := []call{
		{
			Direction: "+",
			Arr:       []int64{-1, 2, -3, 4},
			Expected:  []int64{-3, -1, 2, 4},
		},
		{
			Direction: "-",
			Arr:       []int64{-1, 2, -3, 4},
			Expected:  []int64{4, 2, -1, -3},
		},
	}

	for _, c := range calls {
		expected, err := sortarray.GetInsertSort(c.Arr, c.Direction)
		if err != nil {
			t.Errorf("expected to return no error, but received: %s", err)
		}
		for i := 0; i < len(c.Arr); i++ {
			if c.Expected[i] != expected[i] {
				t.Errorf("результат %+v не совпалает с проверкой %+v", c.Expected[i], expected[i])
				break
			}
		}
	}
}

func TestSortTableFailure(t *testing.T) {
	type call struct {
		Direction string
		Arr       []int64
		//Expected  []int64
	}
	calls := []call{
		//Not supported operation
		{
			Direction: "a",
			Arr:       []int64{-1, 2, -3, 4},
		},
		{
			Direction: "*",
			Arr:       []int64{-1, 2, -3, 4},
		},
		{
			Direction: "%",
			Arr:       []int64{-1, 2, -3, 4},
		},
	}

	for _, c1 := range calls {
		//fmt.Println(c1)
		_, err := sortarray.GetInsertSort(c1.Arr, c1.Direction)
		//fmt.Println(err)
		if err == nil {
			t.Errorf("expected to return error, but received no error")
			//break
		}
	}
}

func TestSortGlobal(t *testing.T) {
	type call struct {
		Direction string
		Arr       []int64
		Expected  []int64
	}
	//todo поменять количество итераций
	for i := 0; i < 11; i++ {
		randSlice, _ := sortarray.GetRandSlice()
		//todo randSlice
		//fmt.Println("len randSlice:",len(randSlice))

		verifiedRandSliseIncr := make([]int64, len(randSlice))
		copy(verifiedRandSliseIncr, randSlice)
		//todo verifiedRandSliseIncr
		//	fmt.Println("verifiedRandSliseIncr")
		//	fmt.Println(verifiedRandSliseIncr)

		verifiedRandSliseDecr := make([]int64, len(randSlice))
		copy(verifiedRandSliseDecr, randSlice)
		//fmt.Println("Sort verifiedRandSliseDecr")
		//fmt.Println(verifiedRandSliseDecr)

		//todo verifiedRandSliseIncr
		//fmt.Println("verifiedRandSliseIncr")
		//fmt.Println(verifiedRandSliseIncr)
		sort.Slice(verifiedRandSliseIncr, func(i, j int) bool { return verifiedRandSliseIncr[i] < verifiedRandSliseIncr[j] })
		//fmt.Println("Sort verifiedRandSliseIncr")
		//fmt.Println(verifiedRandSliseIncr)

		sort.Slice(verifiedRandSliseDecr, func(i, j int) bool { return verifiedRandSliseDecr[i] > verifiedRandSliseDecr[j] })
		//fmt.Println(verifiedRandSliseDecr)

		calls := []call{

			{
				Direction: "+",
				Arr:       randSlice,
				Expected:  verifiedRandSliseIncr,
			},
			{
				Direction: "-",
				Arr:       randSlice,
				Expected:  verifiedRandSliseDecr,
			},
		}

		for _, c := range calls {
			expected, err := sortarray.GetInsertSort(c.Arr, c.Direction)
			if err != nil {
				t.Errorf("expected to return no error, but received: %s", err)
			}
			for i := 0; i < len(c.Arr); i++ {
				if c.Expected[i] != expected[i] {
					//fmt.Println(expected)
					//fmt.Println(c.Expected)

					t.Errorf("результат %+v[%v] не совпалает с проверкой %+v[%v]", c.Expected[i], i, expected[i], i)
					break
				}
			}
		}
	}
}
