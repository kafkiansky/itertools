package iterslice

import (
	"iter"
	"strings"

	"golang.org/x/exp/constraints"
)

// Iterator creates iter.Seq[T] from any []T.
func Iterator[T any](s []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := range s {
			if !yield(s[i]) {
				return
			}
		}
	}
}

// Filter filters iter.Seq[T] with F and returns iter.Seq[T] containing only those values for which F returns true.
func Filter[T any, F ~func(T) bool](seq iter.Seq[T], f F) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if f(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Map maps iter.Seq[T] with F and return iter.Seq[E].
func Map[T any, E any, F ~func(T) E](seq iter.Seq[T], f F) iter.Seq[E] {
	return func(yield func(E) bool) {
		for v := range seq {
			if !yield(f(v)) {
				return
			}
		}
	}
}

// TryMap tries to convert T to E using F and returns iter.Seq2[T, error].
func TryMap[T any, E any, F ~func(T) (E, error)](seq iter.Seq[T], f F) iter.Seq2[E, error] {
	return func(yield func(E, error) bool) {
		for v := range seq {
			if !yield(f(v)) {
				return
			}
		}
	}
}

// Each iterates over iter.Seq[T] and applies F to each element.
func Each[T any, F ~func(T)](seq iter.Seq[T], f F) {
	for v := range seq {
		f(v)
	}
}

// First gets the first element from iter.Seq[T] and returns (T, bool).
func First[T any](seq iter.Seq[T]) (T, bool) {
	return Nth(seq, 1)
}

// Last gets the last element from iter.Seq[T] and returns (T, bool).
func Last[T any](seq iter.Seq[T]) (T, bool) {
	var l T
	var found bool
	for v := range seq {
		l = v
		found = true
	}

	return l, found
}

// Nth gets the N element from iter.Seq[T] and returns (T, bool). N starts from 1.
func Nth[T any](seq iter.Seq[T], n uint) (T, bool) {
	i := uint(1)
	for v := range seq {
		if i == n {
			return v, true
		}

		i++
	}

	return *new(T), false
}

// Range iterates over an integer and yield values as iter.Seq[int].
func Range(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for v := range n {
			if !yield(v) {
				return
			}
		}
	}
}

// Between iterates between l and r (exclusive) and yield values as iter.Seq[int].
func Between(l, r int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := l; i < r; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

// Reduce reduce iter.Seq[T] with A and returns single value. Optional accepts initial.
func Reduce[T any, A ~func(T, T) T](seq iter.Seq[T], a A, initial ...T) T {
	var acc T
	if len(initial) > 0 {
		acc = initial[0]
	}

	for v := range seq {
		acc = a(acc, v)
	}

	return acc
}

// Split splits string using separator and yield chunks as iter.Seq[string].
func Split(s string, sep string) iter.Seq[string] {
	return func(yield func(string) bool) {
		for _, v := range strings.Split(s, sep) {
			if !yield(v) {
				return
			}
		}
	}
}

// Chars iterate over string and returns runes as iter.Seq[string].
func Chars(s string) iter.Seq[string] {
	return func(yield func(string) bool) {
		for _, v := range s {
			if !yield(string(v)) {
				return
			}
		}
	}
}

// ConsumeChannel reads from channel until it closes and yield values as iter.Seq[T].
func ConsumeChannel[T any](c <-chan T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range c {
			if !yield(v) {
				return
			}
		}
	}
}

// Partition consumes an iter.Seq[T], creating two iterators from it.
func Partition[T any, F ~func(T) bool](seq iter.Seq[T], f F) (iter.Seq[T], iter.Seq[T]) {
	var l []T
	var r []T

	for v := range seq {
		if f(v) {
			l = append(l, v)
		} else {
			r = append(r, v)
		}
	}

	return Iterator(l), Iterator(r)
}

// Position searches for an element in an iter.Seq[T], returning its index.
func Position[T comparable](seq iter.Seq[T], i T) (int, bool) {
	var idx int
	for v := range seq {
		if v == i {
			return idx, true
		}

		idx++
	}

	return 0, false
}

// Max finds maximum element in an iter.Seq[T].
func Max[T constraints.Ordered](seq iter.Seq[T]) T {
	var max T
	var init bool

	for v := range seq {
		if !init {
			max = v
			init = true
		} else if v > max {
			max = v
		}
	}

	return max
}

// Min finds minimum element in an iter.Seq[T].
func Min[T constraints.Ordered](seq iter.Seq[T]) T {
	var min T
	var init bool

	for v := range seq {
		if !init {
			min = v
			init = true
		} else if v < min {
			min = v
		}
	}

	return min
}

// Join N iter.Seq[T] to one iterator.
func Join[T any](iters ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for iterIdx := range iters {
			for v := range iters[iterIdx] {
				if !yield(v) {
					return
				}
			}
		}
	}
}
