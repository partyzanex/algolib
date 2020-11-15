package stat_test

import (
	"github.com/partyzanex/algolib/stat"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMin(t *testing.T) {
	exp := -12.5364
	got := stat.Min(1, 2, 3, 4, exp, 34, -5, 8)
	assert.Equal(t, exp, got)
}

func TestMax(t *testing.T) {
	exp := 132.5364
	got := stat.Max(1, 2, 3, 4, exp, 34, -5, 8)
	assert.Equal(t, exp, got)
}
