package binary_test

import (
	"github.com/partyzanex/algolib/tree/binary"
	"github.com/partyzanex/testutils"
	"github.com/stretchr/testify/assert"
	"testing"
)

type node struct {
	ID   int64
	Data string
}

func (n *node) Value() int64 {
	return n.ID
}

func TestTree_Insert(t *testing.T) {
	tree := &binary.Tree{}
	start, end := int64(1), int64(100)
	exp := make(map[int64]*binary.Node)
	s := make([]*node, end)

	for i := start; i <= end; i++ {
		c := &node{
			ID:   testutils.RandInt64(start, end),
			Data: testutils.RandomString(20),
		}

		n := tree.Insert(c)
		exp[c.ID] = n
		s[i-1] = c
	}

	f := tree.Search(&node{ID: testutils.RandInt64(end+1, end+end)})
	assert.Equal(t, true, f == nil)

	i := testutils.RandInt64(start-1, end-1)
	n := tree.Search(&node{ID: s[i].ID}).Value.(*node)
	assert.Equal(t, true, n != nil)
	assert.Equal(t, exp[n.ID].Value.(*node), n)

	// file, _ := os.Create("tree.json")
	// defer file.Close()
	// d, _ := json.Marshal(tree)
	// ioutil.WriteFile("tree.json", d, 0777)
	// d, _ := json.NewEncoder().Encode(v)
}

func TestBinary_Delete(t *testing.T) {
	tree := &binary.Tree{}
	start, end := int64(1), int64(10)
	exp := make(map[int64]*binary.Node)
	s := make([]*node, end)

	for i := start; i <= end; i++ {
		v := testutils.RandInt64(start, end)
		c := &node{
			ID:   v,
			Data: testutils.RandomString(20),
		}
		n := tree.Insert(c)
		exp[v] = n
		s[i-1] = c
	}

	i := testutils.RandInt64(start-1, end-1)
	tree.Delete(s[i])

	n := tree.Search(&node{ID: s[i].ID})
	assert.Equal(t, true, n == nil)
}

func BenchmarkBinary_Insert(b *testing.B) {
	tree := &binary.Tree{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		n := &node{
			ID:   testutils.RandInt64(0, 999999),
			Data: testutils.RandomString(30),
		}
		b.StartTimer()

		_ = tree.Insert(n)
	}
}

func BenchmarkTree_Search(b *testing.B) {
	tree := &binary.Tree{}
	start, end := int64(1), int64(1000)
	exp := make(map[int64]*binary.Node)
	s := make([]*node, end)

	for i := start; i <= end; i++ {
		v := testutils.RandInt64(start, end)
		c := &node{
			ID:   v,
			Data: testutils.RandomString(20),
		}
		n := tree.Insert(c)
		exp[v] = n
		s[i-1] = c
	}

	i := testutils.RandInt64(start-1, end-1)
	v := s[i]

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = tree.Search(v)
	}
}

func BenchmarkBinary_Delete(b *testing.B) {
	tree := &binary.Tree{}
	start, end := int64(1), int64(100000)

	for i := start; i <= end; i++ {
		v := testutils.RandInt64(start, end)
		c := &node{
			ID:   v,
			Data: testutils.RandomString(20),
		}
		tree.Insert(c)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tree.Delete(&node{
			ID: testutils.RandInt64(0, 99999),
		})
	}
}
