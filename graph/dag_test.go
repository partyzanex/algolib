package dag

import (
	"fmt"
	"runtime"
	"testing"
)

func memUsage(m1, m2 *runtime.MemStats) {
	fmt.Println("Alloc:", m2.Alloc-m1.Alloc,
		"TotalAlloc:", m2.TotalAlloc-m1.TotalAlloc,
		"HeapAlloc:", m2.HeapAlloc-m1.HeapAlloc,
		"HeapObjects:", m2.HeapObjects,
		"Sys:", m2.Sys,
	)
}

type vertexInt int

func (v vertexInt) VertexID() interface{} {
	return int(v)
}

func TestDAG_Set(t *testing.T) {
	var m1, m2 runtime.MemStats

	d := New()

	runtime.ReadMemStats(&m1)

	for i := 0; i < 1000; i++ {
		_ = d.Set(vertexInt(i-1), vertexInt(i))
	}

	runtime.ReadMemStats(&m2)

	memUsage(&m1, &m2)
}

func TestDAG_Get(t *testing.T) {
	d := New()

	var (
		v1, v2, v3, v4 vertexInt = 1, 2, 3, 4
		v5, v6, v7, v8 vertexInt = 5, 6, 7, 8

		m1, m2 runtime.MemStats
	)

	runtime.ReadMemStats(&m1)

	noError(t, d.Set(v1, v2, v3, v4))
	noError(t, d.Set(v2, v3, v4))
	noError(t, d.Set(v2, v3, v5))
	noError(t, d.Set(v2, v6, v7))
	noError(t, d.Set(v8, v5, v6, v7))

	runtime.ReadMemStats(&m2)

	memUsage(&m1, &m2)

	runtime.ReadMemStats(&m1)
	items := d.Get(v2)
	runtime.ReadMemStats(&m2)

	memUsage(&m1, &m2)

	if len(items) != 5 {
		t.Error("wrong length")
	}
}

func noError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Error(err)
	}
}

func BenchmarkDAG_Set(b *testing.B) {
	d := New()

	for i := 0; i < b.N; i++ {
		_ = d.Set(vertexInt(i-1), vertexInt(i))
	}
}

func BenchmarkDAG_Get(b *testing.B) {
	d := New()

	var (
		v1, v2, v3, v4 vertexInt = 1, 2, 3, 4
		v5, v6, v7, v8 vertexInt = 5, 6, 7, 8
	)

	noError(b, d.Set(v1, v2, v3, v4))
	noError(b, d.Set(v2, v3, v4))
	noError(b, d.Set(v2, v3, v5))
	noError(b, d.Set(v2, v6, v7))
	noError(b, d.Set(v8, v5, v6, v7))

	for i := 0; i < b.N; i++ {
		_ = d.Get(v2)
	}
}
