package tree_test

import (
	"github.com/partyzanex/algolib/tree"
	"github.com/partyzanex/testutils"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

type node struct {
	ID   int64
	Data string
}

func (n *node) Value() int64 {
	return n.ID
}

func TestBinary_Insert(t *testing.T) {
	btree := &tree.Binary{}
	start, end := int64(1), int64(1000)
	exp := make(map[int64]*node)

	for i := start; i <= end; i++ {
		n := &node{
			ID:   i,
			Data: testutils.RandomString(10),
		}

		btree.Insert(n)
		exp[n.ID] = n
	}

	f := btree.Get(testutils.RandInt64(end+1, end+end))
	assert.Equal(t, true, f == nil)

	n := btree.Get(testutils.RandInt64(start, end))
	assert.Equal(t, true, n != nil)

	got, ok := n.Value().(*node)
	assert.Equal(t, true, ok)
	assert.Equal(t, true, got != nil)

	e, ok := exp[got.ID]
	assert.Equal(t, true, ok)
	assert.Equal(t, got, e)
}

func TestBinary_Get(t *testing.T) {
	btree := &tree.Binary{}
	exp := &node{
		ID:   testutils.RandInt64(1, 2000),
		Data: testutils.RandomString(50),
	}

	btree.Insert(exp)

	for i := 0; i < 1000; i++ {
		btree.Insert(&node{
			ID:   testutils.RandInt64(2000, 3000),
			Data: testutils.RandomString(10),
		})
	}

	got := btree.Get(exp.ID).Value()
	assert.Equal(t, exp, got)
}

func TestBinary_Delete(t *testing.T) {
	btree := &tree.Binary{}
	n := &node{
		ID:   testutils.RandInt64(10, 1000),
		Data: testutils.RandomString(50),
	}
	exp := &node{
		ID:   n.ID - 1,
		Data: testutils.RandomString(50),
	}

	btree.Insert(n)
	btree.Insert(exp)

	for i := 0; i < 1000; i++ {
		btree.Insert(&node{
			ID:   testutils.RandInt64(2000, 3000),
			Data: testutils.RandomString(10),
		})
	}

	btree.Delete(n.ID)

	got := btree.Get(exp.ID).Value()
	assert.Equal(t, exp, got)
}

func TestBinary_Slice(t *testing.T) {
	btree := &tree.Binary{}
	start, end := int64(0), int64(100)
	exp := make(map[int64]*node)

	for i := start; i <= end; i += 10 {
		n := &node{
			ID:   i,
			Data: testutils.RandomString(10),
		}

		btree.Insert(n)
		exp[n.ID] = n
	}

	n := &node{
		ID:   start + 31,
		Data: testutils.RandomString(20),
	}

	btree.Insert(n)
	exp[n.ID] = n

	slice := btree.Slice()

	for _, value := range slice {
		got, ok := value.(*node)
		assert.Equal(t, true, ok)
		assert.Equal(t, true, got != nil)

		e, ok := exp[got.ID]
		assert.Equal(t, true, ok)
		assert.Equal(t, got, e)
	}
}

func BenchmarkCreateStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := node{
			ID: int64(i),
		}
		_ = n
	}
}

func BenchmarkCopyStruct(b *testing.B) {
	n := node{ID: -1}
	for i := 0; i < b.N; i++ {
		m := n
		m.ID = int64(i)
		_ = m
	}
}

func BenchmarkCopyStructFromPtr(b *testing.B) {
	n := &node{}
	for i := 0; i < b.N; i++ {
		m := *n
		m.ID = int64(i)
		_ = m
	}
}

func BenchmarkBinary_Insert(b *testing.B) {
	btree := &tree.Binary{}
	str := testutils.RandomString(10)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = btree.Insert(&node{
			ID:   int64(i),
			Data: str,
		})
	}
}

func BenchmarkBinary_Get(b *testing.B) {
	btree := &tree.Binary{}
	start, end := int64(1), int64(1000)

	for i := start; i <= end; i++ {
		n := &node{
			ID:   i,
			Data: testutils.RandomString(10),
		}

		btree.Insert(n)
	}

	id := int64(500)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = btree.Get(id)
	}
}

func BenchmarkBinary_Delete(b *testing.B) {
	btree := &tree.Binary{}
	start, end := int64(1), int64(1000)

	for i := start; i <= end; i++ {
		n := &node{
			ID:   i,
			Data: testutils.RandomString(10),
		}

		btree.Insert(n)
	}

	id := int64(500)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		btree.Delete(id)
	}
}

func BenchmarkBinary_Slice(b *testing.B) {
	btree := &tree.Binary{}
	start, end := int64(1), int64(1000)

	for i := start; i <= end; i++ {
		n := &node{
			ID:   i,
			Data: testutils.RandomString(10),
		}

		btree.Insert(n)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = btree.Slice()
	}
}

func BenchmarkBinary_Slice2(b *testing.B) {
	start, end := int64(1), int64(1000)
	v := make(map[int64]*node)

	for i := start; i <= end; i++ {
		v[i] = &node{
			ID:   i,
			Data: testutils.RandomString(10),
		}
	}

	s := make([]*node, len(v))
	i := 0

	for _, n := range v {
		s[i] = n
		i++
	}

	for i := 0; i < b.N; i++ {
		sort.Slice(s, func(i, j int) bool {
			return s[i].ID < s[j].ID
		})
	}
}
