package stat

import "sort"

func Min(s ...float64) float64 {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	return s[0]
}

func Max(s ...float64) float64 {
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})

	return s[0]
}
