## Itertools for Go using iterators. 

## Contents
- [Installation](#installation)
- [Usage](#usage)
  - [iterslice](#iterslice)
    - [Iterator](#iterslice-iterator)
    - [Filter](#filter)
    - [Map](#map)
    - [TryMap](#trymap)
    - [Each](#each)
    - [First](#first)
    - [Last](#last)
    - [Nth](#nth)
    - [Range](#range)
    - [Between](#between)
    - [Reduce](#reduce)
    - [Split](#split)
    - [Chars](#chars)
    - [ConsumeChannel](#consumechannel)
    - [Partition](#partition)
    - [Position](#position)
    - [Max](#max)
    - [Min](#min)
    - [Join](#join)
  - [itermap](#itermap)
    - [Iterator](#itermap-iterator)
    - [Keys](#keys)
    - [Values](#values)
    - [Each](#itermap-each)
    - [Join](#itermap-join)
  - [CollectSlice](#collectslice)
  - [CollectMap](#collectmap)
- [Testing](#testing)
- [License](#license)

## Installation

```bash
$ go get github.com/kafkiansky/itertools
```

## Usage

The `itertools` package consists from two subpackages: `iterslice` and `itermap` for working with slices and maps as iterators separately.

### iterslice

#### Iterator

Creates an iterator from slice.

```go
package main

import (
	"log"

    "github.com/kafkiansky/itertools/iterslice"
)

func main() {
    for v := range iterslice.Iterator([]int{1, 2, 3}) {
        log.Println(v)
    }
}
```

#### Filter

Filters an `iter.Seq[T]`.

```go
package main

import (
	"log"

	"github.com/kafkiansky/itertools/iterslice"
)

func main() {
	for v := range iterslice.Filter(
		iterslice.Iterator([]int{1, 2, 3}),
		func(v int) bool {
			return v%2 == 0
		},
	) {
		log.Println(v)
	}
}
```

#### Map

Apply `F` to an `iter.Seq[T]` returning new `iter.Seq[E]`.

```go
package main

import (
	"log"
	"strconv"

	"github.com/kafkiansky/itertools/iterslice"
)

func main() {
	for v := range iterslice.Map(
		iterslice.Iterator([]int{1, 2, 3}),
		strconv.Itoa,
	) {
		log.Println(v)
	}
}
```

#### TryMap

Works as `Map` but with errors.

```go
package main

import (
	"log"
	"strconv"

	"github.com/kafkiansky/itertools/iterslice"
)

func main() {
	for v, err := range iterslice.TryMap(
		iterslice.Iterator([]string{"1", "2", "invalid"}),
		strconv.Atoi,
	) {
		if err != nil {
			log.Println(err)
		} else {
			log.Println(v)
		}
	}
}
```

#### Each

Iterates over an `iter.Seq[T]` and applies `F`.

```go
package main

import (
	"log"

	"github.com/kafkiansky/itertools/iterslice"
)

func main() {
	iterslice.Each(
		iterslice.Iterator([]int{1, 2, 3}),
		func(v int) {
			log.Println(v)
		},
	)
}
```

#### First

Yields the first element from an `iter.Seq[T]`.

```go
package main

import (
	"log"

	"github.com/kafkiansky/itertools/iterslice"
)

func main() {
	log.Println(
		iterslice.First(
			iterslice.Iterator([]int{6, 7, 8}),
		),
	) // 6, true
}
```

#### Last

Yield the last element from an `iter.Seq[T]`.

```go
package main

import (
	"log"

	"github.com/kafkiansky/itertools/iterslice"
)

func main() {
	log.Println(
		iterslice.Last(
			iterslice.Iterator([]int{6, 7, 8}),
		),
	) // 8, true
}
```

#### Nth

Yields the `N` element from `iter.Seq[T]`. Index starts from 1.

```go
package main

import (
	"log"

	"github.com/kafkiansky/itertools/iterslice"
)

func main() {
	log.Println(
		iterslice.Nth(
			iterslice.Iterator([]int{6, 7, 8}),
			2,
		),
	) // 7, true
}
```

#### Range

Creates an iterator from sequence.

```go
package main

import (
	"log"

	"github.com/kafkiansky/itertools/iterslice"
)

func main() {
	for v := range iterslice.Range(5) {
		log.Println(v)
	}
}
```

#### Between

Creates an iterator starting from left value to rigth exclusive.

```go
package main

import (
	"log"

	"github.com/kafkiansky/itertools/iterslice"
)

func main() {
	for v := range iterslice.Between(5, 10) {
		log.Println(v)
	}
}
```

#### Reduce

Reduces an `iter.Seq[T]` to a single value using an `F`. Optional accepts an initial value.

```go
package main

import (
	"log"

	"github.com/kafkiansky/itertools/iterslice"
)

func main() {
	log.Println(
		iterslice.Reduce(
			iterslice.Iterator([]int{1, 2, 3}),
			func(a, b int) int { return a + b },
		),
	) // 6
}
```

#### Split

Splits string using separator and yield chunks as an `iter.Seq[string]`.

```go
package main

import (
	"log"

	"github.com/kafkiansky/itertools/iterslice"
)

func main() {
	for v := range iterslice.Split("a,b,c", ",") {
		log.Println(v)
	}
}
```

#### Chars

Yield string chars as runes.

```go
package main

import (
	"log"

	"github.com/kafkiansky/itertools/iterslice"
)

func main() {
	for v := range iterslice.Chars("итератор") {
		log.Println(v)
	}
}
```

#### ConsumeChannel

Consume values from an `chan T` and yields values as `iter.Seq[T]` until channel closes.

```go
package main

import (
	"log"

	"github.com/kafkiansky/itertools/iterslice"
)

func main() {
	c := make(chan int, 3)
	for v := range 3 {
		c <- v
	}
	close(c)

	for v := range iterslice.ConsumeChannel(c) {
		log.Println(v)
	}
}
```

#### Partition

Creates two iterators from one using `F`.

```go
package main

import (
	"log"

	"github.com/kafkiansky/itertools/iterslice"
)

func main() {
	even, odd := iterslice.Partition(
		iterslice.Iterator([]int{1, 2, 3}),
		func(v int) bool {
			return v%2 == 0
		},
	)

	for v := range even {
		log.Printf("even: %d", v)
	}

	for v := range odd {
		log.Printf("odd: %d", v)
	}
}
```

#### Position

Search index of value in an iterator.

```go
package main

import (
	"log"

	"github.com/kafkiansky/itertools/iterslice"
)

func main() {
	log.Println(
		iterslice.Position(
			iterslice.Iterator([]int{1, 2, 3}),
			2,
		),
	) // 1 (index of value 2)
}
```

#### Max

Search the max value in an `iter.Seq[constraints.Order]`.

```go
package main

import (
	"log"

	"github.com/kafkiansky/itertools/iterslice"
)

func main() {
	log.Println(
		iterslice.Max(
			iterslice.Iterator([]int{10, 2, 3}),
		),
	)
}
```

#### Min

Search the min value in an `iter.Seq[constraints.Ordered]`.

```go
package main

import (
	"log"

	"github.com/kafkiansky/itertools/iterslice"
)

func main() {
	log.Println(
		iterslice.Min(
			iterslice.Iterator([]int{-10, -2, -3}),
		),
	) // -10
}
```

#### Join

Joins n iterators to the single one.

```go
package main

import (
	"log"

	"github.com/kafkiansky/itertools/iterslice"
)

func main() {
	for v := range iterslice.Join(
		iterslice.Between(0, 3),
		iterslice.Between(6, 10),
	) {
		log.Println(v)
	}
}
```

### itermap

#### Iterator

Creates an iterator `iter.Seq2[K, V]` from `map[K]V`.

```go
package main

import (
	"log"

	"github.com/kafkiansky/itertools/itermap"
)

func main() {
	for k, v := range itermap.Iterator(map[string]int{"x": 1, "y": 2}) {
		log.Printf("key: %s, value: %d", k, v)
	}
}
```

#### Keys

Creates an iterator `iter.Seq[K]` from map keys.

```go
package main

import (
	"log"

	"github.com/kafkiansky/itertools/itermap"
)

func main() {
	for k := range itermap.Keys(
		map[string]int{"x": 1, "y": 2},
	) {
		log.Printf("key: %s", k)
	}
}
```

#### Values

Creates an iterator `iter.Seq[V]` from map values.

```go
package main

import (
	"log"

	"github.com/kafkiansky/itertools/itermap"
)

func main() {
	for v := range itermap.Values(
		map[string]int{"x": 1, "y": 2},
	) {
		log.Printf("value: %d", v)
	}
}
```

#### Each

Iterates over an `iter.Seq2[K, V]` an applies `F` for each key-value pair.

```go
package main

import (
	"log"

	"github.com/kafkiansky/itertools/itermap"
)

func main() {
	itermap.Each(
		itermap.Iterator(map[string]int{"x": 1, "y": 2}),
		func(k string, v int) {
			log.Printf("key: %s, value: %d", k, v)
		},
	)
}
```

#### Join

Joins n iterators `iter.Seq[K, V]` to the single one.

```go
package main

import (
	"log"

	"github.com/kafkiansky/itertools/itermap"
)

func main() {
	for k, v := range itermap.Join(
		itermap.Iterator(map[string]int{"x": 1, "y": 2}),
		itermap.Iterator(map[string]int{"z": 3}),
	) {
		log.Printf("key: %s, value: %d", k, v)
	}
}
```

#### CollectSlice

Collects an `iter.Seq[T]` to the slice `[]T`.

```go
package main

import (
	"log"

	"github.com/kafkiansky/itertools"
	"github.com/kafkiansky/itertools/iterslice"
)

func main() {
	log.Println(
		itertools.CollectSlice(
			iterslice.Between(5, 10),
		),
	) // [5, 6, 7, 8, 9]
}
```

#### CollectMap

Collects an `iter.Seq2[K, V]` to the map `map[K]V`.

```go
package main

import (
	"log"

	"github.com/kafkiansky/itertools"
	"github.com/kafkiansky/itertools/itermap"
)

func main() {
	log.Println(
		itertools.CollectMap(
			itermap.Iterator(map[string]int{"x": 1, "y": 2}),
		),
	) // map[x:1 y:2]
}
```

## Testing

``` bash
$ GOEXPERIMENT=rangefunc go test ./...
```  

## License

The MIT License (MIT). See [License File](LICENSE) for more information.