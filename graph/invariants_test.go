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

func TestAdjecentVertices(t *testing.T) {
	edges := []graph.Edge{
		{0, 1},
		{0, 2},
		{0, 3},
		{1, 3},
		{1, 2},
	}

	g := graph.NewGraph(5)
	g = g.AddEdges(edges)

	received := g.AdjecentVertices(0)
	expected := []graph.Vertex{1, 2, 3}

	if len(received) != len(expected) {
		t.Errorf("DegreeSequence: expected %v, received %v", expected, received)
	}

	for i := range received {
		if received[i] != expected[i] {
			t.Errorf("DegreeSequence: expected %v, received %v", expected, received)
		}
	}
}

func TestVerticesConnected(t *testing.T) {
	edges := []graph.Edge{
		{0, 1},
		{0, 2},
		{0, 3},
		{1, 3},
		{1, 2},
	}

	g := graph.NewGraph(5)
	g = g.AddEdges(edges)

	input := [][]graph.Vertex{
		{0, 1},
		{0, 4},
		{2, 3},
		{4, 4},
	}
	expected := []bool{true, false, true, false}
	received := make([]bool, len(expected))

	for i := range received {
		received[i] = g.VerticesConnected(input[i][0], input[i][1])
	}

	for i := range received {
		if received[i] != expected[i] {
			t.Errorf("VerticesConnected: for the vertices %v and %v, expected %v, received %v", input[i][0], input[i][1], expected[i], received[i])
		}
	}
}
