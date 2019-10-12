package graph

func EdgeSet(n uint8) []Edge {
	edges := make([]Edge, n*(n-1)/2)

	for i, index := uint8(0), 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			edges[index] = Edge{i, j}
			index += 1
		}
	}

	return edges
}

func Powerset(n uint8) [][]uint8 {
	ps := make([][]uint8, 1 << n)

	for i := uint64(0); i < 1 << n; i++ {
		numElems := uint8(0)
		for j := uint8(0); j < n; j++ {
			// This gets the nth digit from the right
			numElems += uint8((i >> j) % 2)
		}

		ps[i] = make([]uint8, numElems)

		for j, k := uint8(0), 0; j < n; j++ {
			if (i >> j) % 2 == 1 {
				ps[i][k] = j
				k += 1
			}
		}
	}

	return ps
}

func AllNGraphs(n uint8) []Graph {
	es := EdgeSet(n)
	ps := Powerset(uint8(len(es)))
	graphs := make([]Graph, len(ps))
	avgSize := uint(len(es)/2)

	for i, set := range ps {
		edges := make([]Edge, 0, avgSize)
		for _, v := range set {
			edges = append(edges, es[v])
		}

		graphs[i] = NewGraph(n).AddEdges(edges)
	}

	return graphs
}

