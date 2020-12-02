package tree_test

import (
	"github.com/partyzanex/algolib/tree"
	"github.com/partyzanex/testutils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBalanced_Insert(t *testing.T) {
	btree := &tree.Balanced{}
	start, end := int64(1), int64(100)
	exp := make(map[int64]*node)
	s := make([]*node, end)

	for i := start; i <= end; i++ {
		c := &node{
			ID:   i,
			Data: testutils.RandomString(20),
		}

		btree.Insert(c)
		exp[c.ID] = c
		s[i-1] = c
	}

	f := btree.Search(&node{ID: testutils.RandInt64(end+1, end+end)})
	assert.Equal(t, true, f == nil)

	i := testutils.RandInt64(start-1, end-1)
	n := btree.Search(&node{ID: s[i].ID}).Value.(*node)
	assert.Equal(t, true, n != nil)
	assert.Equal(t, exp[n.ID], n)

	// file, _ := os.Create("tree.json")
	// defer file.Close()
	// d, _ := json.Marshal(btree.Root)
	// ioutil.WriteFile("tree.json", d, 0777)
	// d, _ := json.NewEncoder(file).Encode(btree)
}

func TestBalanced_Delete(t *testing.T) {
	btree := &tree.Balanced{}
	start, end := int64(1), int64(100)
	exp := make(map[int64]*node)
	s := make([]*node, end)

	for i := start; i <= end; i++ {
		c := &node{
			ID:   i,
			Data: testutils.RandomString(20),
		}

		btree.Insert(c)
		exp[c.ID] = c
		s[i-1] = c
	}

	i := testutils.RandInt64(start-1, end-1)
	btree.Delete(s[i])

	n := btree.Search(&node{ID: s[i].ID})
	assert.Equal(t, true, n == nil)
}

func BenchmarkBalanced_Insert(b *testing.B) {
	btree := &tree.Balanced{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		n := &node{
			ID:   testutils.RandInt64(0, 999999),
			Data: testutils.RandomString(30),
		}
		b.StartTimer()

		btree.Insert(n)
	}
}

func BenchmarkBalanced_Search(b *testing.B) {
	btree := &tree.Balanced{}
	start, end := int64(1), int64(1000)
	s := make([]*node, end)

	for i := start; i <= end; i++ {
		v := testutils.RandInt64(start, end)
		c := &node{
			ID:   v,
			Data: testutils.RandomString(20),
		}
		btree.Insert(c)
		s[i-1] = c
	}

	i := testutils.RandInt64(start-1, end-1)
	v := s[i]

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = btree.Search(v)
	}
}
