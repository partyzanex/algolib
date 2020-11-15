package tree

type Comparable interface {
	Value() int64
}

type Node struct {
	Value       Comparable
	Left, Right *Node
}

func (n *Node) Insert(v Comparable) *Node {
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

func (n *Node) Slice() []Comparable {
	var slice []Comparable

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

func (n *Node) Search(v Comparable) *Node {
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

type Binary struct {
	Root *Node
}

func (t *Binary) Insert(v Comparable) *Node {
	if t.Root == nil {
		t.Root = &Node{Value: v}
		return t.Root
	}

	return t.Root.Insert(v)
}

func (t *Binary) Slice() []Comparable {
	return t.Root.Slice()
}

func (t *Binary) Delete(v Comparable) {
	var (
		parent  *Node
		current = t.Root
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
						t.Root.Insert(n.Value)
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
				t.Root = nil
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

func (t *Binary) Walk(f WalkFunc) {
	t.Root.Walk(f)
}

func (t *Binary) Search(v Comparable) *Node {
	return t.Root.Search(v)
}
