package btree

import "github.com/partyzanex/algolib/tree"

type Tree struct {
	Root   *Node
	height height
}

func (t *Tree) Insert(v tree.Comparable) {
	t.Root = insert(t.Root, v)
}

func (t *Tree) Delete(v tree.Comparable) {
	t.Root = remove(t.Root, v)
}

func (t *Tree) CheckBalance() bool {
	return checkBalance(t.Root, &t.height)
}

func (t *Tree) Search(v tree.Comparable) *Node {
	return t.Root.search(v)
}

type WalkFunc func(n *Node) bool

func (t *Tree) Walk(f WalkFunc) {
	t.Root.walk(f)
}

type height struct {
	value int
}

func checkBalance(t *Node, h *height) bool {
	if t == nil {
		h.value = 0

		return true
	}

	leftHeight, rightHeight := &height{}, &height{}
	l := checkBalance(t.Left, leftHeight)
	r := checkBalance(t.Right, rightHeight)
	lh, rh := leftHeight.value, rightHeight.value

	if lh > rh {
		h.value = lh
	} else {
		h.value = rh
	}

	h.value++

	if (lh-rh >= 2) || (rh-lh >= 2) {
		return false
	}

	return l && r
}
