package binary

import "github.com/partyzanex/algolib/tree"

type Tree struct {
	root *Node
}

func (t *Tree) Insert(v tree.Comparable) *Node {
	if t.root == nil {
		t.root = &Node{Value: v}

		return t.root
	}

	return t.root.Insert(v)
}

func (t *Tree) Slice() []tree.Comparable {
	return t.root.Slice()
}

func (t *Tree) Delete(v tree.Comparable) {
	var (
		parent  *Node
		current = t.root
		i       = v.Value()
	)

	for current != nil {
		j := current.Value.Value()

		switch {
		case i < j:
			parent = current
			current = current.Left
		case i > j:
			parent = current
			current = current.Right
		case i == j:
			if current.Left != nil {
				right := current.Right
				*current = *current.Left

				if right != nil {
					right.Walk(func(n *Node) bool {
						t.root.Insert(n.Value)

						return true
					})
				}

				return
			}

			if current.Right != nil {
				*current = *current.Right

				return
			}

			if parent == nil {
				t.root = nil

				return
			}

			if parent.Left == current {
				parent.Left = nil
			} else {
				parent.Right = nil
			}

			return
		}
	}
}

func (t *Tree) Walk(f WalkFunc) {
	t.root.Walk(f)
}

func (t *Tree) Search(v tree.Comparable) *Node {
	return t.root.Search(v)
}
