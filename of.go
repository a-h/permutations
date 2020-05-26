package permutations

// OfStrings returns all permutations of the input a.
func OfStrings(a []string, f func(permutation []string) (stop bool)) {
	N(len(a), func(permutation []int) (stop bool) {
		op := make([]string, len(a))
		for i := 0; i < len(permutation); i++ {
			op[i] = a[permutation[i]]
		}
		return f(op)
	})
}

// N returns all permutations of n values.
func N(n int, f func(permutation []int) (stop bool)) {
	values := make([]int, n)
	for i := 0; i < n; i++ {
		values[i] = i
	}
	generate(n, values, f)
}

func generate(k int, A []int, f func(permutation []int) (stop bool)) {
	if k == 1 {
		op := make([]int, len(A))
		copy(op, A)
		if stop := f(op); stop {
			return
		}
		return
	}
	generate(k-1, A, f)
	for i := 0; i < k-1; i++ {
		if k%2 == 0 {
			A[i], A[k-1] = A[k-1], A[i]
		} else {
			A[0], A[k-1] = A[k-1], A[0]
		}
		generate(k-1, A, f)
	}
}

/*
https://en.wikipedia.org/wiki/Heap%27s_algorithm

procedure generate(k : integer, A : array of any):
    if k = 1 then
        output(A)
    else
        // Generate permutations with kth unaltered
        // Initially k == length(A)
        generate(k - 1, A)

        // Generate permutations for kth swapped with each k-1 initial
        for i := 0; i < k-1; i += 1 do
            // Swap choice dependent on parity of k (even or odd)
            if k is even then
                swap(A[i], A[k-1]) // zero-indexed, the kth is at k-1
            else
                swap(A[0], A[k-1])
            end if
            generate(k - 1, A)

        end for
    end if
*/
