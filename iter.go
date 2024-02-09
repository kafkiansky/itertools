package itertools

import "iter"

// Collect collects values from iter.Seq[T] to []T.
func CollectSlice[T any](seq iter.Seq[T]) []T {
	var s []T

	for v := range seq {
		s = append(s, v)
	}

	return s
}

// Collect2 collects values from iter.Seq2[K, V] to map[K]V.
func CollectMap[K comparable, V any](seq iter.Seq2[K, V]) map[K]V {
	m := make(map[K]V)

	for k, v := range seq {
		m[k] = v
	}

	return m
}
