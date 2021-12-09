package btree_test

import (
	"testing"

	"github.com/partyzanex/algolib/tree/btree"
	"github.com/partyzanex/testutils"
	"github.com/stretchr/testify/assert"
)

type node struct {
	ID   int64
	Data string
}

func (n *node) Value() int64 {
	return n.ID
}

func TestBalanced_Insert(t *testing.T) {
	tree := &btree.Tree{}
	start, end := int64(1), int64(100)
	exp := make(map[int64]*node)
	s := make([]*node, end)

	for i := start; i <= end; i++ {
		c := &node{
			ID:   i,
			Data: testutils.RandomString(20),
		}

		tree.Insert(c)
		exp[c.ID] = c
		s[i-1] = c
	}

	f := tree.Search(&node{ID: testutils.RandInt64(end+1, end+end)})
	assert.Equal(t, true, f == nil)

	i := testutils.RandInt64(start-1, end-1)
	n := tree.Search(&node{ID: s[i].ID}).Value.(*node)
	assert.Equal(t, true, n != nil)
	assert.Equal(t, exp[n.ID], n)
}

func TestBalanced_Delete(t *testing.T) {
	tree := &btree.Tree{}
	start, end := int64(1), int64(100)
	exp := make(map[int64]*node)
	s := make([]*node, end)

	for i := start; i <= end; i++ {
		c := &node{
			ID:   i,
			Data: testutils.RandomString(20),
		}

		tree.Insert(c)
		exp[c.ID] = c
		s[i-1] = c
	}

	i := testutils.RandInt64(start-1, end-1)
	tree.Delete(s[i])

	n := tree.Search(&node{ID: s[i].ID})
	assert.Equal(t, true, n == nil)
}

func BenchmarkBalanced_Insert(b *testing.B) {
	tree := &btree.Tree{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()

		n := &node{
			ID:   testutils.RandInt64(0, 999999),
			Data: testutils.RandomString(30),
		}

		b.StartTimer()

		tree.Insert(n)
	}
}

func BenchmarkBalanced_Search(b *testing.B) {
	tree := &btree.Tree{}
	start, end := int64(1), int64(1000)
	s := make([]*node, end)

	for i := start; i <= end; i++ {
		v := testutils.RandInt64(start, end)
		c := &node{
			ID:   v,
			Data: testutils.RandomString(20),
		}

		tree.Insert(c)
		s[i-1] = c
	}

	i := testutils.RandInt64(start-1, end-1)
	v := s[i]

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = tree.Search(v)
	}
}
