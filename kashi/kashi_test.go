package kashi_test

import (
	"github.com/partyzanex/algolib/kashi"
	"math"
	"testing"
)

func TestPow(t *testing.T) {
	x := float64(3)
	n := uint64(10)
	a := math.Pow(x, float64(n))
	b := kashi.Pow(x, n)

	if a != b {
		t.Fatal("wrong result")
	}
}

func BenchmarkMathPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = math.Pow(3, 17)
	}
}

func BenchmarkPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = kashi.Pow(3, 17)
	}
}
