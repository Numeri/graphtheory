package graph_test

import (
	"graphtheory/graph"
	"testing"
)

func TestNumEdges(t *testing.T) {
	edges := []graph.Edge{
		{0, 1},
		{0, 2},
		{0, 3},
		{1, 3},
		{1, 2},
	}

	g := graph.NewGraph(4)
	g = g.AddEdges(edges)

	received := g.NumEdges()
	expected := uint(len(edges))

	if received != expected {
		t.Errorf("NumEdges: expected %d, received %d", expected, received)
	}
}

func TestDegreeSequence(t *testing.T) {
	edges := []graph.Edge{
		{0, 1},
		{0, 2},
		{0, 3},
		{1, 3},
		{1, 2},
	}

	g := graph.NewGraph(5)
	g = g.AddEdges(edges)

	received := g.DegreeSequence()
	expected := []uint{3, 3, 2, 2, 0}

	for i := range received {
		if received[i] != expected[i] {
			t.Errorf("DegreeSequence: expected %v, received %v", expected, received)
		}
	}
}
