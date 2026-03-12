package permutations

import "iter"

// Of returns an iterator over all permutations of the input slice.
// The yielded slice is reused between iterations; callers must copy it if needed.
func Of[T any](a []T) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		n := len(a)
		if n == 0 {
			return
		}
		values := make([]T, n)
		copy(values, a)
		generate(values, yield)
	}
}

// N returns an iterator over all permutations of n values (0 to n-1).
// The yielded slice is reused between iterations; callers must copy it if needed.
func N(n int) iter.Seq[[]int] {
	return func(yield func([]int) bool) {
		if n == 0 {
			return
		}
		values := make([]int, n)
		for i := range n {
			values[i] = i
		}
		generate(values, yield)
	}
}

// generate implements the iterative version of Heap's algorithm.
func generate[T any](A []T, yield func([]T) bool) {
	n := len(A)
	c := make([]int, n)

	if !yield(A) {
		return
	}

	i := 0
	for i < n {
		if c[i] < i {
			if i%2 == 0 {
				A[0], A[i] = A[i], A[0]
			} else {
				A[c[i]], A[i] = A[i], A[c[i]]
			}
			if !yield(A) {
				return
			}
			c[i]++
			i = 0
		} else {
			c[i] = 0
			i++
		}
	}
}
