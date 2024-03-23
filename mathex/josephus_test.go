package mathex

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJosephusFlavius(t *testing.T) {
	testCases := []*struct {
		n, k     int
		expected int
	}{
		{
			n:        6,
			k:        2,
			expected: 5,
		},
		{
			n:        5,
			k:        2,
			expected: 3,
		},
		{
			n:        8,
			k:        8,
			expected: 4,
		},
		{
			n:        3,
			k:        9,
			expected: 2,
		},
		{
			n:        4,
			k:        3,
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d %d", tc.n, tc.k), func(t *testing.T) {
			actual := JosephusFlavius(tc.n, tc.k)
			if actual != tc.expected {
				t.Errorf("JosephusFlavius(%d, %d) = %d, expected %d", tc.n, tc.k, actual, tc.expected)
			}
		})
	}
}

func TestJosephusFlaviusOrder(t *testing.T) {
	testCases := []*struct {
		n, k     int
		expected []int
	}{
		{
			n:        6,
			k:        2,
			expected: []int{2, 4, 6, 3, 1, 5},
		},
		{
			n:        5,
			k:        2,
			expected: []int{2, 4, 1, 5, 3},
		},
		{
			n:        8,
			k:        8,
			expected: []int{8, 1, 3, 6, 5, 2, 7, 4},
		},
		{
			n:        3,
			k:        9,
			expected: []int{3, 1, 2},
		},
		{
			n:        4,
			k:        3,
			expected: []int{3, 2, 4, 1},
		},
		{
			n:        4,
			k:        3,
			expected: []int{3, 2, 4, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d %d", tc.n, tc.k), func(t *testing.T) {
			actual := JosephusFlaviusOrder(tc.n, tc.k)
			assert.ElementsMatch(t, tc.expected, actual)
		})
	}
}

func BenchmarkJosephusFlavius(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = JosephusFlavius(i%10+2, i%10+1)
	}
}

func BenchmarkJosephusFlaviusOrder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = JosephusFlaviusOrder(i%10+2, i%10+1)
	}
}
