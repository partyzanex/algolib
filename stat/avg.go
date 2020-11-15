package stat

import "sort"

func Avg(s ...float64) float64 {
	var sum float64

	for _, n := range s {
		sum += n
	}

	return sum / float64(len(s))
}

func Median(s ...float64) float64 {
	n := len(s)

	if n == 0 {
		return 0
	}

	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	if n%2 == 1 {
		return s[(n-1)/2]
	}

	return Avg(s[n/2], s[n/2-1])
}

func Variance(s ...float64) float64 {
	avg := Avg(s...)
	sum := float64(0)

	for _, v := range s {
		sum += (v - avg) * (v - avg)
	}

	return sum / float64(len(s))
}
