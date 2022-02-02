package sortarray_test

import (
	"sort"
	"testing"

	"math/rand"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/alcor67/repo-Go-level-1-home-work/project-sort-numbers/sortarray"
)

func TestTestifySort(t *testing.T) {

	arr := []int64{-1, 2, -3, 4}
	expected := []int64{-3, -1, 2, 4}

	direction := "+"
	received, _ := sortarray.GetInsertSort(arr, direction)

	assert.Equal(t, received, expected, "they should be equal")

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
		received, err := sortarray.GetInsertSort(c.Arr, c.Direction)
		if err != nil {
			t.Errorf("expected to return no error, but received: %s", err)
		}
		assert.Equal(t, received, c.Expected, "they should be equal")
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
		received, err := sortarray.GetInsertSort(c.Arr, c.Direction)
		if err != nil {
			t.Errorf("expected to return no error, but received: %s", err)
		}
		assert.NotEqual(t, received, c.Expected, "they should not be equal")
	}
}

func TestTestifySortGlobal(t *testing.T) {
	type call struct {
		Direction string
		Arr       []int64
		Expected  []int64
	}

	for i := 0; i < 100; i++ {
		randSlice, _ := getRandSlice()

		verifiedRandSliseIncr := make([]int64, len(randSlice))
		copy(verifiedRandSliseIncr, randSlice)

		verifiedRandSliseDecr := make([]int64, len(randSlice))
		copy(verifiedRandSliseDecr, randSlice)

		sort.Slice(verifiedRandSliseIncr, func(i, j int) bool { return verifiedRandSliseIncr[i] < verifiedRandSliseIncr[j] })

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
			received, err := sortarray.GetInsertSort(c.Arr, c.Direction)
			if err != nil {
				t.Errorf("expected to return no error, but received: %s", err)
			}
			assert.Equal(t, received, c.Expected, "they should be equal")
		}
	}
}

func getRandSlice() ([]int64, error) {
	var randNumber int64
	inputNums := []int64{}
	rand.Seed(getRandSeed())
	cycleLen := rand.Intn(10000)

	for i := 0; i < cycleLen; i++ {

		rand.Seed(getRandSeed())
		randNumberNew := rand.Int63()

		if randNumber == randNumberNew {
			continue
		}
		randNumber = randNumberNew
		randOutPut := (randNumber/1e15 - 5000) * 2
		inputNums = append(inputNums, randOutPut)
	}
	return inputNums, nil
}

func getRandSeed() int64 {
	var randSeed int64 = 0
	timeNano := time.Now().UnixNano()

	randSeedNew := int64(timeNano-timeNano/1e6*1e6) / 10

	if randSeed != randSeedNew {
		randSeed = randSeedNew
	}
	return randSeed
}
