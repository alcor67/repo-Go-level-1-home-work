package sortarray_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alcor67/repo-Go-level-1-home-work/project-sort-numbers/sortarray"
)

func TestTestifySort(t *testing.T) {

	arr := []int64{-1, 2, -3, 4}
	arrTest := []int64{-3, -1, 2, 4}

	direction := "+"
	Expected, err := sortarray.GetInsertSort(arr, direction)
	// assert equality Testify
	assert.Equal(t, Expected, arrTest, "they should be equal")
	if err != nil {
		t.Errorf("expected to return no error, but received: %s", err)
	}
}

func TestTestifySortTableOk(t *testing.T) {
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
		// assert equality Testify
		assert.Equal(t, expected, c.Expected, "they should be equal")
	}
}

func TestTestifySortTableFailure(t *testing.T) {
	type call struct {
		Direction string
		Arr       []int64
		Expected  []int64
	}
	calls := []call{
		{
			Direction: "+",
			Arr:       []int64{-1, 2, -3, 4},
			Expected:  []int64{-3, -1, 2, 0},
		},
		{
			Direction: "-",
			Arr:       []int64{-1, 2, 4, -3},
			Expected:  []int64{4, 2, -1, -0},
		},
	}

	for _, c := range calls {
		expected, err := sortarray.GetInsertSort(c.Arr, c.Direction)
		if err != nil {
			t.Errorf("expected to return no error, but received: %s", err)
		}
		// assert equality Testify
		assert.NotEqual(t, expected, c.Expected, "they should not be equal")
	}
}

func TestTestifySortGlobal(t *testing.T) {
	type call struct {
		Direction string
		Arr       []int64
		Expected  []int64
	}

	for i := 0; i < 1; i++ {
		randSlice, _ := sortarray.GetRandSlice()

		//копируем слайсы
		verifiedRandSliseIncr := make([]int64, len(randSlice))
		copy(verifiedRandSliseIncr, randSlice)

		verifiedRandSliseDecr := make([]int64, len(randSlice))
		copy(verifiedRandSliseDecr, randSlice)
		//если поломаем тест
		/*
			fmt.Println("verifiedRandSliseDecr[0]", verifiedRandSliseDecr[0])
			verifiedRandSliseDecr[0] += 1
			fmt.Println("verifiedRandSliseDecr[0]", verifiedRandSliseDecr[0])
		*/
		sort.Slice(verifiedRandSliseIncr, func(i, j int) bool { return verifiedRandSliseIncr[i] < verifiedRandSliseIncr[j] })
		//если поломаем тест
		/*
			fmt.Println("verifiedRandSliseIncr[0]", verifiedRandSliseIncr[0])
			verifiedRandSliseIncr[0] += 1
			fmt.Println("verifiedRandSliseIncr[0]", verifiedRandSliseIncr[0])
		*/
		sort.Slice(verifiedRandSliseDecr, func(i, j int) bool { return verifiedRandSliseDecr[i] > verifiedRandSliseDecr[j] })

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
			// assert equality Testify
			assert.Equal(t, expected, c.Expected, "they should be equal")
		}
	}
}
