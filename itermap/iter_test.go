package itermap

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/kafkiansky/itertools"
	"github.com/stretchr/testify/assert"
)

func TestIterator(t *testing.T) {
	assert.Equal(
		t,
		map[string]int{"kafkiansky": 2},
		itertools.CollectMap(Iterator(map[string]int{"kafkiansky": 2})),
	)
}

func TestKeys(t *testing.T) {
	keys := itertools.CollectSlice(Keys(map[string]int{"kafkiansky": 0, "itermap": 1, "iter": 2}))
	sort.Strings(keys)

	assert.Equal(
		t,
		[]string{"iter", "itermap", "kafkiansky"},
		keys,
	)
}

func TestValues(t *testing.T) {
	values := itertools.CollectSlice(Values(map[string]int{"kafkiansky": 0, "itermap": 1, "iter": 2}))
	sort.Ints(values)

	assert.Equal(
		t,
		[]int{0, 1, 2},
		values,
	)
}

func TestEach(t *testing.T) {
	var n []string
	Each(Iterator(map[string]int{"kafkiansky": 1}), func(k string, v int) {
		n = append(n, fmt.Sprintf("%s:%d", k, v))
	})

	assert.Equal(t, []string{"kafkiansky:1"}, n)
}

func TestJoin(t *testing.T) {
	assert.True(
		t,
		reflect.DeepEqual(
			map[string]int{"kafkiansky": 1, "itertools": 2},
			itertools.CollectMap(
				Join(
					Iterator(map[string]int{"itertools": 2}),
					Iterator(map[string]int{"kafkiansky": 1}),
				),
			),
		),
	)
}
