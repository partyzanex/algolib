package euclid_test

import (
	"github.com/partyzanex/algolib/euclid"
	"testing"
)

var nums = []struct {
	A, B float64
}{
	{1234567890, 12},
	{1223345, 185},
	{5654, 798746541},
	{56254, 798746},
	{640, 480},
}

func TestPlainGCD(t *testing.T) {
	for _, num := range nums {
		gcd := euclid.PlainGCD(num.A, num.B)
		t.Log(gcd)
	}
}

func BenchmarkPlainGCD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = euclid.PlainGCD(1234567890, 12)
	}
}

func TestExtendedGCD(t *testing.T) {
	for _, num := range nums {
		gcd := euclid.ExtendedGCD(int64(num.A), int64(num.B))
		t.Log(gcd)
	}
}

func BenchmarkExtendedGCD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = euclid.ExtendedGCD(1234567890, 12)
	}
}
