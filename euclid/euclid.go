package euclid

// PlainGCD implements greatest common divisor
// slow algorithm
func PlainGCD(a, b float64) float64 {
	for {
		if a == b {
			break
		}

		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}

// ExtendedGCD implements fastest algorithm, but only int numbers
func ExtendedGCD(a, b int64) int64 {
	for {
		if a == 0 || b == 0 {
			break
		}

		if a > b {
			a %= b
		} else {
			b %= a
		}
	}

	if a == 0 {
		return b
	}

	return a
}
