package tree_test

import (
	"github.com/partyzanex/algolib/tree"
	"github.com/partyzanex/testutils"
	"testing"
)

func TestBTree_Insert(t *testing.T) {
	btree := tree.Balanced{}
	start, end := int64(1), int64(100)

	for i := start; i <= end; i++ {
		btree.Insert(&node{
			ID:   testutils.RandInt64(0, i),
			Data: testutils.RandomString(30),
		})
	}

	btree.Dump()
}
