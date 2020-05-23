package permutations

// Strings.
func Strings(a []string) chan []string {
	c := make(chan []string)
	go func() {
		for indices := range N(len(a)) {
			op := make([]string, len(a))
			for i := 0; i < len(a); i++ {
				op[i] = a[indices[i]]
			}
			c <- op
		}
		close(c)
	}()
	return c
}

// N values.
func N(n int) chan []int {
	c := make(chan []int)
	values := make([]int, n)
	for i := 0; i < n; i++ {
		values[i] = i
	}
	go func() {
		generate(n, values, c)
		close(c)
	}()
	return c
}

func generate(k int, A []int, output chan []int) {
	if k == 1 {
		op := make([]int, len(A))
		copy(op, A)
		output <- op
		return
	}
	generate(k-1, A, output)
	for i := 0; i < k-1; i++ {
		if k%2 == 0 {
			A[i], A[k-1] = A[k-1], A[i]
		} else {
			A[0], A[k-1] = A[k-1], A[0]
		}
		generate(k-1, A, output)
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
