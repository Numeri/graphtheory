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

func TestPath(t *testing.T) {
	edges := []graph.Edge{
		{0, 1},
		{0, 2},
		{0, 3},
		{1, 3},
		{1, 2},
		{2, 4},
	}

	g := graph.NewGraph(6)
	g = g.AddEdges(edges)

	a := graph.Vertex(0)
	b := graph.Vertex(4)

	expected := []graph.Vertex{0, 2, 4}
	received := g.Path(a, b)

	for i := range received {
		if received[i] != expected[i] {
			t.Errorf("Path: for the vertices %v and %v, expected path %v, received %v", a, b, expected, received)
			break
		}
	}

	expected = []graph.Vertex{}
	a = 0
	b = 5
	received = g.Path(a, b)

	for i := range received {
		if received[i] != expected[i] {
			t.Errorf("Path: for the vertices %v and %v, expected path %v, received %v", a, b, expected, received)
			break
		}
	}
}

func TestConnectivityWeight(t *testing.T) {
	edges := []graph.Edge{
		{0, 2},
		{1, 2},
		{2, 3},
		{2, 4},
		{4, 5},
	}

	g := graph.NewGraph(6)
	g = g.AddEdges(edges)

	expected := []uint8{1, 1, 4, 1, 2, 1}
	received := make([]uint8, g.Size)

	for i := range received {
		received[i] = g.ConnectivityWeight(graph.Vertex(i))
	}

	for i := range received {
		if received[i] != expected[i] {
			t.Errorf("ConnectivityWeight: for the vertices %v, expected %v, received %v", i, expected[i], received[i])
		}
	}
}

func TestConnectivitySequence(t *testing.T) {
	edges := []graph.Edge{
		{0, 2},
		{1, 2},
		{2, 3},
		{2, 4},
		{4, 5},
	}

	g := graph.NewGraph(7)
	g = g.AddEdges(edges)

	expected := []uint8{4, 2, 1, 1, 1, 1, 0}
	received := g.ConnectivitySequence()

	for i := range received {
		if received[i] != expected[i] {
			t.Errorf("ConnectivitySequence: expected %v, received %v", expected, received)
		}
	}
}
