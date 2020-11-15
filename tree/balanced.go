package tree

import (
	"fmt"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

type BNode struct {
	Value       Comparable
	Left, Right *BNode

	height, bal int
}

func (n *BNode) Height() int {
	if n == nil {
		return 0
	}

	return max(n.Right.Height(), n.Left.Height())
}

func (n *BNode) Balance() int {
	return n.Right.Height() - n.Left.Height()
}

func (n *BNode) Insert(v Comparable) int {
	i := v.Value()
	j := n.Value.Value()

	switch {
	case i == j:
		n.Value = v
		return 0
	case i < j:
		if n.Left == nil {
			n.Left = &BNode{
				Value:  v,
				height: 1,
			}

			break
		}

		if balance := n.Left.Insert(v); balance < -1 || balance > 1 {
			n.balance(n.Left)
		}
	case i > j:
		if n.Right == nil {
			n.Right = &BNode{
				Value:  v,
				height: 1,
			}

			break
		}

		if balance := n.Right.Insert(v); balance < -1 || balance > 1 {
			n.balance(n.Right)
		}
	}

	return n.Balance()
}

func (n *BNode) rotateLeft(c *BNode) {
	r := c.Right
	c.Right = r.Left
	c.bal -= 1
	r.Left = c

	if c == n.Left {
		n.Left = c
		return
	}

	n.Right = r
}

func (n *BNode) rotateRight(c *BNode) {
	l := c.Left
	c.Left = l.Right
	l.Right = c
	c.bal = 0
	n.bal = 0

	if c == n.Left {
		n.Left = l
		return
	}

	n.Right = l
}

func (n *BNode) rotateRightLeft(c *BNode) {
	c.Right.Left.bal = 1
	c.rotateRight(c.Right)
	c.Right.bal = 1
	n.rotateLeft(c)
}

func (n *BNode) rotateLeftRight(c *BNode) {
	c.Left.Right.bal = -1
	c.rotateLeft(c.Left)
	c.Left.bal = -1
	n.rotateRight(c)
}

func (n *BNode) balance(c *BNode) {
	switch {
	case c.bal == -2 && c.Left.bal == -1:
		n.rotateRight(c)
	case c.bal == 2 && c.Right.bal == 1:
		n.rotateLeft(c)
	case c.bal == -2 && c.Left.bal == 1:
		n.rotateLeftRight(c)
	case c.bal == 2 && c.Right.bal == -1:
		n.rotateRightLeft(c)
	}
}

func (n *BNode) Search(v Comparable) *BNode {
	if n == nil {
		return nil
	}

	i := v.Value()
	j := n.Value.Value()

	switch {
	case i == j:
		return n
	case i < j:
		return n.Left.Search(v)
	default:
		return n.Right.Search(v)
	}
}

func (n *BNode) Walk(f WalkBFunc) {
	if n.Left != nil && !f(n.Left) {
		return
	}

	if !f(n) {
		return
	}

	if n.Right != nil && !f(n.Right) {
		return
	}
}

func (n *BNode) Dump(i int, lr string) {
	if n == nil {
		return
	}

	ident := ""

	if i > 0 {
		ident = strings.Repeat(" ", (i-1)*2) + "+" + lr + "--"
	}

	fmt.Printf("%s%d[%d]\n", ident, n.Value.Value(), n.bal)
	n.Left.Dump(i+1, "L")
	n.Right.Dump(i+1, "R")
}

type Balanced struct {
	Root *BNode
}

func (t *Balanced) Insert(v Comparable) {
	if t.Root == nil {
		t.Root = &BNode{
			Value: v,
		}

		return
	}

	if balance := t.Root.Insert(v); balance < -1 || balance > 1 {
		t.balance()
	}
}

func (t *Balanced) balance() {
	if t == nil || t.Root == nil {
		return
	}

	fake := &BNode{Left: t.Root}
	fake.balance(t.Root)

	t.Root = fake.Left
}

func (t *Balanced) Search(v Comparable) *BNode {
	if t.Root == nil {
		return nil
	}

	return t.Root.Search(v)
}

type WalkBFunc func(n *BNode) bool

func (t *Balanced) Walk(f WalkBFunc) {
	if t.Root == nil {
		return
	}

	t.Root.Walk(f)
}

func (t *Balanced) Dump() {
	if t.Root == nil {
		return
	}

	t.Root.Dump(0, "")
}
