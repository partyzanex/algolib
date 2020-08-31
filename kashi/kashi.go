package kashi

// algorithm Al Kashi
func Pow(x float64, n uint64) float64 {
	y := float64(1)
	z := x

	for {
		if n%2 != 0 {
			y *= z
		}

		if n == 0 {
			break
		}

		z *= z

		n /= 2
	}

	return y
}
