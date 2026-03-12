## permutations

Go implementation of [Heap's algorithm](https://en.wikipedia.org/wiki/Heap%27s_algorithm) using Go 1.23+ iterators.

### Usage

```go
import "github.com/a-h/permutations"

// Permutations of any slice type
for p := range permutations.Of([]string{"a", "b", "c"}) {
	fmt.Println(p)
}

// Permutations of indices 0 to n-1
for p := range permutations.N(3) {
	fmt.Println(p)
}
```

Output:

```
[a b c]
[b a c]
[c a b]
[a c b]
[b c a]
[c b a]
```

### Requirements

Go 1.23 or later (uses `iter.Seq` and range-over-func).
