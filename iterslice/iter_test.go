package iterslice

import (
	"errors"
	"strconv"
	"testing"

	"github.com/kafkiansky/itertools"
	"github.com/stretchr/testify/assert"
)

func TestIterator(t *testing.T) {
	assert.EqualValues(
		t,
		[]string{"kafkiansky", "iterslice", "iter"},
		itertools.CollectSlice(Iterator([]string{"kafkiansky", "iterslice", "iter"})),
	)
}

func TestRange(t *testing.T) {
	assert.EqualValues(
		t,
		[]int{0, 1, 2, 3, 4},
		itertools.CollectSlice(Range(5)),
	)
}

func TestBetween(t *testing.T) {
	assert.EqualValues(
		t,
		[]int{5, 6, 7, 8, 9},
		itertools.CollectSlice(Between(5, 10)),
	)
}

func TestFirst(t *testing.T) {
	v, ok := First(Iterator([]string{"kafkiansky", "iterslice", "iter"}))

	assert.True(t, ok)
	assert.Equal(t, "kafkiansky", v)

	v, ok = First(Iterator[string](nil))
	assert.False(t, ok)
	assert.Equal(t, "", v)
}

func TestLast(t *testing.T) {
	v, ok := Last(Iterator([]string{"kafkiansky", "iterslice", "iter"}))

	assert.True(t, ok)
	assert.Equal(t, "iter", v)

	v, ok = Last(Iterator[string](nil))
	assert.False(t, ok)
	assert.Equal(t, "", v)
}

func TestNth(t *testing.T) {
	v, ok := Nth(Iterator([]string{"kafkiansky", "iterslice", "iter"}), 2)

	assert.True(t, ok)
	assert.Equal(t, "iterslice", v)

	v, ok = Nth(Iterator[string](nil), 10)
	assert.False(t, ok)
	assert.Equal(t, "", v)
}

func TestReduce(t *testing.T) {
	for _, testCase := range []struct {
		name    string
		values  []int
		f       func(int, int) int
		acc     int
		initial int
	}{
		{
			name:   "addition",
			values: []int{2, 2, 1, 5},
			f:      func(i1, i2 int) int { return i1 + i2 },
			acc:    10,
		},
		{
			name:    "multiply",
			values:  []int{2, 2, 1, 5},
			f:       func(i1, i2 int) int { return i1 * i2 },
			acc:     20,
			initial: 1,
		},
		{
			name:    "subtraction",
			values:  []int{2, 2, 1, 5},
			f:       func(i1, i2 int) int { return i1 - i2 },
			acc:     10,
			initial: 20,
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(
				t,
				testCase.acc,
				Reduce(Iterator(testCase.values), testCase.f, testCase.initial),
			)
		})
	}
}

func TestMap(t *testing.T) {
	assert.EqualValues(
		t,
		[]string{"10", "20", "30"},
		itertools.CollectSlice(Map(Iterator([]int{10, 20, 30}), strconv.Itoa)),
	)
}

func TestTryMap(t *testing.T) {
	assert.EqualValues(
		t,
		map[int]error{
			10: nil,
			0: &strconv.NumError{
				Func: "Atoi",
				Num:  "invalid",
				Err:  errors.New("invalid syntax"),
			},
		},
		itertools.CollectMap(TryMap(Iterator([]string{"10", "invalid"}), strconv.Atoi)),
	)
}

func TestEach(t *testing.T) {
	var n []int
	Each(Range(5), func(a int) {
		n = append(n, a)
	})

	assert.EqualValues(t, []int{0, 1, 2, 3, 4}, n)
}

func TestSplit(t *testing.T) {
	assert.EqualValues(
		t,
		[]string{"kafkiansky", "iterslice", "iter"},
		itertools.CollectSlice(Split("kafkiansky,iterslice,iter", ",")),
	)
}

func TestChars(t *testing.T) {
	assert.EqualValues(
		t,
		[]string{"т", "е", "с", "т"},
		itertools.CollectSlice(Chars("тест")),
	)
}

func TestConsumeChannel(t *testing.T) {
	c := make(chan int, 5)
	for i := range 5 {
		c <- i
	}
	close(c)

	assert.EqualValues(
		t,
		[]int{0, 1, 2, 3, 4},
		itertools.CollectSlice(ConsumeChannel(c)),
	)
}

func TestPartition(t *testing.T) {
	even, odd := Partition(
		Between(1, 4),
		func(v int) bool {
			return v%2 == 0
		},
	)

	assert.EqualValues(t, []int{2}, itertools.CollectSlice(even))
	assert.EqualValues(t, []int{1, 3}, itertools.CollectSlice(odd))
}

func TestPosition(t *testing.T) {
	v, ok := Position(Range(5), 2)
	assert.True(t, ok)
	assert.Equal(t, 2, v)
}

func TestMax(t *testing.T) {
	assert.Equal(
		t,
		11,
		Max(Iterator([]int{2, 0, 11, 7, 9})),
	)
}

func TestMin(t *testing.T) {
	assert.Equal(
		t,
		-11,
		Min(Iterator([]int{2, 0, -11, 7, 9})),
	)
}

func TestJoin(t *testing.T) {
	assert.EqualValues(
		t,
		[]int{1, 2, 3, 8, 9, 10},
		itertools.CollectSlice(Join(Iterator([]int{1, 2, 3}), Iterator([]int{8, 9, 10}))),
	)
}
