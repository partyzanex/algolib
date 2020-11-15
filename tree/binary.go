package tree

type Comparable interface {
	Value() int64
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
