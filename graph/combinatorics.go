package graph

func EdgeSet(n uint8) []Edge {
	edges := make([]Edge, n*(n-1)/2)

	for i, index := uint8(0), 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			edges[index] = Edge{i, j}
			index += 1
		}
	}

	return edges
}

func Powerset(n uint8) [][]uint8 {
	ps := make([][]uint8, 1<<n)

	for i := uint64(0); i < 1<<n; i++ {
		numElems := uint8(0)
		for j := uint8(0); j < n; j++ {
			// This gets the nth digit from the right
			numElems += uint8((i >> j) % 2)
		}

		ps[i] = make([]uint8, numElems)

		for j, k := uint8(0), 0; j < n; j++ {
			if (i>>j)%2 == 1 {
				ps[i][k] = j
				k += 1
			}
		}
	}

	return ps
}

func PowersetIth(n uint8, i uint64) []uint8 {
	numElems := uint8(0)

	for j := uint8(0); j < n; j++ {
		// This gets the nth digit from the right
		numElems += uint8((i >> j) % 2)
	}

	ps := make([]uint8, numElems)

	for j, k := uint8(0), 0; j < n; j++ {
		if (i>>j)%2 == 1 {
			ps[k] = j
			k += 1
		}
	}

	return ps
}

func Factorial(n uint8) uint64 {
	res := uint64(1)

	//if n == 0 {
	//	return 1
	//}

	switch n {
	case 0:
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 6
	case 4:
		return 24
	case 5:
		return 120
	case 6:
		return 720
	case 7:
		return 5040
	case 8:
		return 40320
	case 9:
		return 362880
	case 10:
		return 3628800
	}

	for i := uint64(1); i <= uint64(n); i++ {
		res = res * i
	}

	return res
}

func PermutationIth(n uint8, i uint64) []uint8 {
	j := i
	m := n - 1
	sym := make([]uint8, n)
	for k := uint8(0); k < n; k++ {
		sym[k] = k
	}

	perm := make([]uint8, n)

	//Nasty, but unsigned integers force us to it
	for m+1 != 0 {
		f := Factorial(m)
		div := j / f
		perm[n-m-1] = sym[div]
		sym = append(sym[:div], sym[div+1:]...)
		j = j % f
		m--
	}

	return perm
}

func NumNGraphs(n uint8) uint64 {
	return 1 << (n * (n - 1) / 2)
}

func AllNGraphs(n uint8) []Graph {
	es := EdgeSet(n)
	ps := Powerset(uint8(len(es)))
	graphs := make([]Graph, len(ps))
	avgSize := uint(len(es) / 2)

	for i, set := range ps {
		edges := make([]Edge, 0, avgSize)
		for _, v := range set {
			edges = append(edges, es[v])
		}

		graphs[i] = NewGraph(n).AddEdges(edges)
	}

	return graphs
}

func AllNGraphsIth(n uint8, i uint64) Graph {
	es := EdgeSet(n)
	ps := PowersetIth(uint8(len(es)), i)

	edges := make([]Edge, 0, len(es)/2)

	for _, v := range ps {
		edges = append(edges, es[v])
	}

	return NewGraph(n).AddEdges(edges)
}
