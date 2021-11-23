package binary

import "github.com/partyzanex/algolib/tree"

type Node struct {
	Value       tree.Comparable
	Left, Right *Node
}

func (n *Node) Insert(v tree.Comparable) *Node {
	i := v.Value()
	j := n.Value.Value()

	if i == j {
		return n
	}

	if i < j {
		if n.Left == nil {
			n.Left = &Node{Value: v}

			return n.Left
		}

		return n.Left.Insert(v)
	}

	if n.Right == nil {
		n.Right = &Node{Value: v}

		return n.Right
	}

	return n.Right.Insert(v)
}

func (n *Node) Slice() []tree.Comparable {
	var slice []tree.Comparable

	if n.Left != nil {
		slice = append(slice, n.Left.Slice()...)
	}

	slice = append(slice, n.Value)

	if n.Right != nil {
		slice = append(slice, n.Right.Slice()...)
	}

	return slice
}

type WalkFunc func(n *Node) bool

func (n *Node) Walk(f WalkFunc) {
	if n.Left != nil {
		n.Left.Walk(f)

		return
	}

	if !f(n) {
		return
	}

	if n.Right != nil {
		n.Right.Walk(f)

		return
	}
}

func (n *Node) Search(v tree.Comparable) *Node {
	if n == nil {
		return nil
	}

	i := v.Value()
	j := n.Value.Value()

	if i == j {
		return n
	}

	if i < j {
		if n.Left == nil {
			return nil
		}

		return n.Left.Search(v)
	}

	if n.Right == nil {
		return nil
	}

	return n.Right.Search(v)
}
