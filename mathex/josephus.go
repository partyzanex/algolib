package mathex

func JosephusFlavius(n, k int) int {
	if n == 1 {
		return 1
	}

	return (JosephusFlavius(n-1, k)+k-1)%n + 1
}

func JosephusFlaviusOrder(n, k int) []int {
	p := make([]int, n)
	for i := range p {
		p[i] = i + 1
	}

	i, j := 0, 0
	r := make([]int, n)

	for len(p) > 0 {
		i = (i + k - 1) % len(p)
		r[j] = p[i]
		j++
		p = append(p[:i], p[i+1:]...)
	}

	return r
}
