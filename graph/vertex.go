package dag

type Vertex interface {
	VertexID() interface{}
}

type vertices map[interface{}]*Vertex

type edges map[*Vertex]*uniqueList

type uniqueList struct {
	uniq map[*Vertex]int
	list []*Vertex
}

func (v *uniqueList) insert(at *Vertex) {
	_, ok := v.uniq[at]
	if ok {
		return
	}

	v.uniq[at] = len(v.list)
	v.list = append(v.list, at)
}
