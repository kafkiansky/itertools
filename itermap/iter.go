package itermap

import "iter"

// Iterator creates iter.Seq2[K, V] from any map[K]V.
func Iterator[M ~map[K]V, K comparable, V any](m M) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m {
			if !yield(k, v) {
				return
			}
		}
	}
}

// Keys yields map[K]V keys as iter.Seq[K].
func Keys[M ~map[K]V, K comparable, V any](m M) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k := range m {
			if !yield(k) {
				return
			}
		}
	}
}

// Values yields map[K]V values as iter.Seq[V].
func Values[M ~map[K]V, K comparable, V any](m M) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range m {
			if !yield(v) {
				return
			}
		}
	}
}

// Each iterates over iter.Seq2[K, V] and applies F to each key and value.
func Each[K comparable, V any, F ~func(K, V)](seq iter.Seq2[K, V], f F) {
	for k, v := range seq {
		f(k, v)
	}
}

// Join N iter.Seq2[K, T] to one iterator.
func Join[K comparable, V any](iters ...iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for iterIdx := range iters {
			for k, v := range iters[iterIdx] {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}
