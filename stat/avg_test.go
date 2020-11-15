package stat_test

import (
	"github.com/partyzanex/algolib/stat"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAvg(t *testing.T) {
	exp := float64(1+2+3+10+15+20) / 6
	got := stat.Avg(1, 2, 3, 10, 15, 20)
	assert.Equal(t, exp, got)
}

func TestMedian(t *testing.T) {
	exp := float64(7)
	got := stat.Median(1, 2, 9, 3, 5, 7, 8, 9, 15, 4, 1, 3, 5, 50, 5, 100, 110, 60, 90)
	assert.Equal(t, exp, got)
}

func TestVariance(t *testing.T) {
	got := stat.Variance(1, 3, 5, 10, 20, 30, 50, 100, 200, 300, 500)
	t.Log(got)
}
