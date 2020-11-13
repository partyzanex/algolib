package tree

type Value interface {
	Value() int64
}

type Binary struct {
	left, right *Binary
	value       Value
}

func (b *Binary) Insert(v Value) *Binary {
	return insert(b, v)
}

func (b *Binary) Value() Value {
	if b == nil {
		return nil
	}

	return b.value
}

func (b *Binary) Get(id int64) *Binary {
	if b.value == nil {
		switch {
		case b.left != nil && b.right != nil:
			lv := b.left.value.Value()
			rv := b.right.value.Value()

			if lv == id {
				return b.left
			}

			if rv == id {
				return b.right
			}

			if rv-id < lv-id {
				return b.left.Get(id)
			}

			b.right.Get(id)
		case b.left != nil && b.right == nil:
			return b.left.Get(id)
		case b.left == nil && b.right != nil:
			return b.right.Get(id)
		default:
			return nil
		}
	}

	i := b.value.Value()

	if i == id {
		return b
	}

	if id < i {
		if b.left == nil {
			return nil
		}

		return b.left.Get(id)
	}

	if b.right == nil {
		return nil
	}

	return b.right.Get(id)
}

func (b *Binary) Delete(id int64) {
	if t := b.Get(id); t != nil {
		t.value = nil
	}
}

func (b *Binary) Slice() (slice []Value) {
	if b.left != nil {
		slice = append(slice, b.left.Slice()...)
	}

	if b.value != nil {
		slice = append(slice, b.value)
	}

	if b.right != nil {
		slice = append(slice, b.right.Slice()...)
	}

	return
}

func insert(b *Binary, v Value) *Binary {
	if b == nil {
		return &Binary{value: v}
	}

	if b.value == nil {
		b.value = v

		return b
	}

	i := v.Value()

	if i == b.value.Value() {
		b.value = v

		return b
	}

	if i < b.value.Value() {
		b.left = insert(b.left, v)

		return b
	}

	b.right = insert(b.right, v)

	return b
}
