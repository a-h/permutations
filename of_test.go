package permutations

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestN(t *testing.T) {
	var tests = []struct {
		n        int
		expected [][]int
	}{
		{
			n: 1,
			expected: [][]int{
				{0},
			},
		},
		{
			n: 2,
			expected: [][]int{
				{0, 1},
				{1, 0},
			},
		},
		{
			n: 3,
			expected: [][]int{
				{0, 1, 2},
				{1, 0, 2},
				{2, 0, 1},
				{0, 2, 1},
				{1, 2, 0},
				{2, 1, 0},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprint(tt.n), func(t *testing.T) {
			var actual [][]int
			N(tt.n, func(permutation []int) (stop bool) {
				actual = append(actual, permutation)
				return false
			})
			if diff := cmp.Diff(actual, tt.expected); diff != "" {
				t.Errorf("(%d): expected %v, actual %v", tt.n, tt.expected, actual)
				t.Errorf(diff)
			}
		})
	}
}

func TestStrings(t *testing.T) {
	var tests = []struct {
		input    []string
		expected [][]string
	}{
		{
			input: []string{"a"},
			expected: [][]string{
				{"a"},
			},
		},
		{
			input: []string{"a", "b"},
			expected: [][]string{
				{"a", "b"},
				{"b", "a"},
			},
		},
		{
			input: []string{"a", "b", "c"},
			expected: [][]string{
				{"a", "b", "c"},
				{"b", "a", "c"},
				{"c", "a", "b"},
				{"a", "c", "b"},
				{"b", "c", "a"},
				{"c", "b", "a"},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			var actual [][]string
			OfStrings(tt.input, func(permutation []string) (stop bool) {
				actual = append(actual, permutation)
				return false
			})
			if diff := cmp.Diff(actual, tt.expected); diff != "" {
				t.Errorf("(%v): expected %v, actual %v", tt.input, tt.expected, actual)
				t.Errorf(diff)
			}
		})
	}
}
