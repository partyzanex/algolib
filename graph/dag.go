package dag

type DAG struct {
	vertices vertices
	edges    edges
}

func New() *DAG {
	return &DAG{
		vertices: make(vertices),
		edges:    make(edges),
	}
}

func (d *DAG) Set(from Vertex, to ...Vertex) error {
	var (
		fromID = from.VertexID()

		fromPtr *Vertex
		node    *uniqueList
		ok      bool
	)

	fromPtr, ok = d.vertices[fromID]
	if !ok {
		fromPtr = &from
		d.vertices[fromID] = fromPtr
	}

	node, ok = d.edges[fromPtr]
	if !ok {
		node = &uniqueList{
			uniq: make(map[*Vertex]int),
			list: make([]*Vertex, 0),
		}
		d.edges[fromPtr] = node
	}

	for i := range to {
		var (
			toID     = to[i].VertexID()
			toPtr    *Vertex
			toExists bool
		)

		toPtr, toExists = d.vertices[toID]
		if !toExists {
			toPtr = &to[i]
			d.vertices[toID] = toPtr
		}

		node.insert(toPtr)
	}

	return nil
}

func (d *DAG) Get(from Vertex) []Vertex {
	fid := from.VertexID()

	f, ok := d.vertices[fid]
	if !ok {
		return nil
	}

	e, ok := d.edges[f]
	if !ok {
		return nil
	}

	res := make([]Vertex, len(e.list))

	for i, v := range e.list {
		res[i] = *v
	}

	return res
}
