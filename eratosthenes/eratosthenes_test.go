package eratosthenes_test

import (
	"github.com/partyzanex/algolib/eratosthenes"
	"testing"
)

func TestPrimes(t *testing.T) {
	r := eratosthenes.Primes(20)
	t.Log(r)

	r = eratosthenes.Primes(200)
	t.Log(r)
}

func BenchmarkPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = eratosthenes.Primes(100000000)
	}
}
