package graph

import (
	"gonum.org/v1/gonum/mat"
	"sort"
)

func (g Graph) NumEdges() uint {
	result := uint(0)

	for i := 0; i < int(g.Size); i++ {
		for j := i + 1; j < int(g.Size); j++ {
			result += uint(g.Adj.At(i, j))
		}
	}

	return result
}

func (g Graph) Degree(v Vertex) uint {
	result := uint(0)

	for i := 0; i < int(g.Size); i++ {
		result += uint(g.Adj.At(i, int(v)))
	}

	return result
}

func (g Graph) DegreeSequence() []uint {
	result := make([]uint, g.Size)

	for i := 0; i < int(g.Size); i++ {
		result[i] = g.Degree(Vertex(i))
	}

	sort.Slice(result, func(i, j int) bool { return result[i] > result[j] })
	return result
}

func (g Graph) AdjecentVertices(v Vertex) []Vertex {
	result := make([]Vertex, 0)

	for i := 0; i < int(g.Size); i++ {
		if g.Adj.At(i, int(v)) == 1 {
			result = append(result, uint8(i))
		}
	}

	return result
}

func (g Graph) VerticesConnected(u, v Vertex) bool {
	if u == v {
		return false
	}

	visited := make([]bool, g.Size)
	ch := make(chan bool)
	go g.verticesConnected_(ch, u, v, visited)
	return <-ch
}

func (g Graph) verticesConnected_(ch chan bool, u, v Vertex, visited []bool) {
	visited[int(u)] = true

	if g.Adj.At(int(u), int(v)) == 1 {
		ch <- true
		return
	}

	neighbours := g.AdjecentVertices(u)

	numGoroutines := 0
	ch2 := make(chan bool)

	for _, n := range neighbours {
		if !visited[int(n)] {
			numGoroutines++
			go g.verticesConnected_(ch2, n, v, visited)
		}
	}

	if numGoroutines == 0 {
		ch <- false
		return
	}

	for i := 0; i < numGoroutines; i++ {
		if <-ch2 == true {
			ch <- true
			return
		}
	}

	ch <- false
}

func (g Graph) Path(u, v Vertex) []Vertex {
	if u == v {
		return []Vertex{}
	}

	visited := make([]bool, g.Size)
	path := make([]Vertex, 0, g.Size/2)
	ch := make(chan []Vertex)

	go g.path_(ch, u, v, visited, path)
	return <-ch
}

func (g Graph) path_(ch chan []Vertex, u, v Vertex, visited []bool, path []Vertex) {
	visited[int(u)] = true

	path = append(path, u)

	if g.Adj.At(int(u), int(v)) == 1 {
		ch <- append(path, v)
		return
	}

	neighbours := g.AdjecentVertices(u)

	numGoroutines := 0
	ch2 := make(chan []Vertex)

	for _, n := range neighbours {
		if !visited[int(n)] {
			numGoroutines++
			go g.path_(ch2, n, v, visited, path)
		}
	}

	if numGoroutines == 0 {
		ch <- []Vertex{}
		return
	}

	shortestPath := []Vertex{}

	for i := 0; i < numGoroutines; i++ {
		returnedPath := <-ch2
		if len(returnedPath) != 0 {
			if len(shortestPath) != 0 || len(returnedPath) < len(shortestPath) {
				shortestPath = returnedPath
			}
		}
	}

	ch <- shortestPath
}

func (g Graph) ConnectivityGraph() Graph {
	cg := mat.NewSymDense(int(g.Size), nil)
	cg.CopySym(g.Adj)

	for i := 0; i < int(g.Size); i++ {
		neighbours := make([]int, 0, g.Size)

		for j := 0; j < int(g.Size); j++ {
			if cg.At(i, j) == 1 {
				neighbours = append(neighbours, j)
			}
		}

		for n := range neighbours {
			for m := n + 1; m < len(neighbours); m++ {
				cg.SetSym(neighbours[n], neighbours[m], 1)
			}
		}
	}

	return Graph{g.Size, cg}
}

func (g Graph) IsConnected() bool {
	cg := g.ConnectivityGraph()

	for i := 1; i < int(g.Size); i++ {
		if cg.Adj.At(0, i) == 0 {
			return false
		}
	}

	return true
}

func (g Graph) Complement() Graph {
	comp_adj := mat.NewSymDense(int(g.Size), nil)

	for i := 0; i < int(g.Size); i++ {
		for j := i + 1; j < int(g.Size); j++ {
			if g.Adj.At(i, j) == 0 {
				comp_adj.SetSym(i, j, 1)
			}
		}
	}

	return Graph{g.Size, comp_adj}
}

//TODO: account for multiple shortest paths
func (g Graph) ConnectivityWeight(v Vertex) uint8 {
	weight := uint8(0)

	for i := 0; i < int(g.Size); i++ {
		for j := i + 1; j < int(g.Size); j++ {
			shortestPath := g.Path(Vertex(i), Vertex(j))

			for _, p := range shortestPath {
				if p == v {
					weight++
				}
			}
		}
	}

	return weight
}

func (g Graph) ConnectivitySequence() []uint8 {
	weights := make([]uint8, g.Size)

	for i := range weights {
		weights[i] = g.ConnectivityWeight(Vertex(i))
	}

	sort.Slice(weights, func(i, j int) bool { return weights[i] > weights[j] })

	return weights
}
