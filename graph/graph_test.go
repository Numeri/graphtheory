package graph_test

import (
	"graphtheory/graph"
	"testing"
)

func TestEqual(t *testing.T) {
	a := graph.Edge{0, 1}
	b := graph.Edge{0, 1}
	c := graph.Edge{1, 0}
	d := graph.Edge{1, 2}

	if !a.Equal(b) || !b.Equal(a) || !a.Equal(c) || !c.Equal(a) {
		t.Errorf("Edge Equal: returned false for equivalent edges")
	}

	if d.Equal(a) || a.Equal(d) || d.Equal(c) || c.Equal(d) {
		t.Errorf("Edge Equal: returned true for non-equivalent edges")
	}
}

func TestNewGraph(t *testing.T) {
	size := uint8(3)
	g := graph.NewGraph(size)
	r, c := g.Adj.Dims()

	if uint8(r) != size || uint8(c) != size {
		t.Errorf("Adjecency Matrix created with wrong dimensions: expected (%d,%d) got (%d,%d)", size, size, r, c)
	}
	if g.Size != size {
		t.Errorf("Graph Size set incorrectly: expected %d, got %d", size, g.Size)
	}
}

func TestString(t *testing.T) {
	g := graph.NewGraph(3)
	got := g.String()
	want := "" +
		"⎡0  0  0⎤\n" +
		"⎢0  0  0⎥\n" +
		"⎣0  0  0⎦"

	if got != want {
		t.Errorf("Graph to String sent unexpected result: expected:\n%v\ngot:\n%v\n", want, got)
	}
}

func TestAddEdge(t *testing.T) {
	g := graph.NewGraph(3)
	g = g.AddEdge(graph.Edge{1, 2})

	if g.Adj.At(1, 2) != 1 || g.Adj.At(2, 1) != 1 {
		t.Errorf("Edge not added correctly")
	}
}

func TestAddEdges(t *testing.T) {
	g := graph.NewGraph(4)

	edges := []graph.Edge{
		{0, 1},
		{2, 1},
		{0, 3},
	}

	g = g.AddEdges(edges)

	if g.Adj.At(0, 1) != 1 ||
		g.Adj.At(1, 0) != 1 ||
		g.Adj.At(2, 1) != 1 ||
		g.Adj.At(1, 2) != 1 ||
		g.Adj.At(0, 3) != 1 ||
		g.Adj.At(3, 0) != 1 {
		t.Errorf("Edges not added correctly")
	}
}
