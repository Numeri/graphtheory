package graph

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

type Graph struct {
	Size uint8
	Adj mat.Symmetric
}

type Edge struct {
	U, V Vertex
}

type Vertex = uint8

func (g Graph) String() string {
	fa := mat.Formatted(g.Adj, mat.Prefix(""), mat.Squeeze())
	return fmt.Sprintf("\n%v\n", fa)
}

func NewGraph(size uint8) Graph {
		return Graph{size, mat.NewSymDense(int(size), nil)}
}

func (g Graph) AddEdge(e Edge) Graph {
	adj := mat.NewSymDense(int(g.Size), nil)
	adj.CopySym(g.Adj)

	adj.SetSym(int(e.U), int(e.V), 1)
	return Graph{g.Size, adj}
}

func (g Graph) AddEdges(edges []Edge) Graph {
	adj := mat.NewSymDense(int(g.Size), nil)
	adj.CopySym(g.Adj)

	for _, e := range edges {
		adj.SetSym(int(e.U), int(e.V), 1)
	}

	return Graph{g.Size, adj}
}

