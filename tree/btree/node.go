package btree

import "github.com/partyzanex/algolib/tree"

type Node struct {
	Value       tree.Comparable
	Left, Right *Node
	height      int
}

func (n *Node) Height() int {
	if n == nil {
		return 0
	}

	return n.height
}

func (n *Node) search(v tree.Comparable) *Node {
	if n == nil {
		return nil
	}

	x := n.Value.Value()
	y := v.Value()

	if y == x {
		return n
	}

	if y < x {
		if n.Left == nil {
			return nil
		}

		return n.Left.search(v)
	}

	if n.Right == nil {
		return nil
	}

	return n.Right.search(v)
}

func (n *Node) walk(f WalkFunc) {
	if n.Left != nil {
		n.Left.walk(f)

		return
	}

	if !f(n) {
		return
	}

	if n.Right != nil {
		n.Right.walk(f)

		return
	}
}

func rightRotate(y *Node) *Node {
	x := y.Left
	t := x.Right
	x.Right = y
	y.Left = t

	y.height = max(y.Left.Height(), y.Right.Height()) + 1
	x.height = max(x.Left.Height(), x.Right.Height()) + 1

	return x
}

func leftRotate(x *Node) *Node {
	y := x.Right
	t := y.Left
	y.Left = x
	x.Right = t

	x.height = max(x.Left.Height(), x.Right.Height()) + 1
	y.height = max(y.Left.Height(), y.Right.Height()) + 1

	return y
}

func getBalance(n *Node) int {
	if n == nil {
		return 0
	}

	return n.Left.Height() - n.Right.Height()
}

func insert(n *Node, v tree.Comparable) *Node {
	if n == nil {
		return &Node{Value: v}
	}

	x := n.Value.Value()
	y := v.Value()

	switch {
	case y < x:
		n.Left = insert(n.Left, v)

		break
	case y > x:
		n.Right = insert(n.Right, v)

		break
	default:
		return n
	}

	n.height = max(n.Left.Height(), n.Right.Height()) + 1
	balance := getBalance(n)

	if balance > 1 {
		x = n.Left.Value.Value()
		if y < x {
			return rightRotate(n)
		} else if y > x {
			n.Left = leftRotate(n.Left)
			return rightRotate(n)
		}
	}

	if balance < -1 {
		x = n.Right.Value.Value()
		if y > x {
			return leftRotate(n)
		} else if y < x {
			n.Right = rightRotate(n.Right)

			return leftRotate(n)
		}
	}

	return n
}

func remove(n *Node, v tree.Comparable) *Node {
	if n == nil {
		return n
	}

	x := n.Value.Value()
	y := v.Value()

	switch {
	case y < x:
		n.Left = remove(n.Left, v)

		break
	case y > x:
		n.Right = remove(n.Right, v)
	default:
		if n.Left == nil || n.Right == nil {
			var t *Node

			if t == n.Left {
				t = n.Right
			} else {
				t = n.Left
			}

			if t == nil {
				t = n
				n = nil
			} else {
				n = t
			}
		} else {
			t := nodeWithMinValue(n.Right)
			n.Value = t.Value
			n.Right = remove(n.Right, t.Value)
		}
	}

	if n == nil {
		return n
	}

	n.height = max(n.Left.Height(), n.Right.Height()) + 1
	balance := getBalance(n)

	if balance > 1 {
		if getBalance(n.Left) >= 0 {
			return rightRotate(n)
		} else {
			n.Left = leftRotate(n.Left)

			return rightRotate(n)
		}
	}

	if balance < -1 {
		if getBalance(n.Right) <= 0 {
			return leftRotate(n)
		} else {
			n.Right = rightRotate(n.Right)

			return leftRotate(n)
		}
	}

	return n
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func nodeWithMinValue(n *Node) *Node {
	c := n

	for c.Left != nil {
		c = c.Left
	}

	return c
}
