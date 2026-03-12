package permutations

import "testing"

func BenchmarkN8(b *testing.B) {
	for range b.N {
		for range N(8) {
		}
	}
}

func BenchmarkN10(b *testing.B) {
	for range b.N {
		for range N(10) {
		}
	}
}
