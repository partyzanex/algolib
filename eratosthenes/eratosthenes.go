package eratosthenes

import "math"

func Primes(max int64) []int64 {
	n := int(math.Abs(float64(max)))

	if n > 9 {
		n /= 2
	}

	primes := make([]int64, n)
	j := 0

	min := int64(3)
	primes[j] = 1
	j++
	primes[j] = 2
	j++

	for i := min; i <= max; i += 2 {
		if i > 5 && i%5 == 0 {
			continue
		}

		if i > 3 && i%3 == 0 {
			continue
		}

		primes[j] = i
		j++
	}

	return primes[0:j]
}
