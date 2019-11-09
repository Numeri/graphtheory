package graph

func (g Graph) invariantMatch_(h Graph, isComplement bool) bool {
	var equal bool = g.Size == h.Size && g.NumEdges() == h.NumEdges()

	gds := g.DegreeSequence()
	hds := h.DegreeSequence()

	for i := range gds {
		if gds[i] != hds[i] {
			return false
		}
	}

	if isComplement == false {
		gc := g.Complement()
		hc := g.Complement()

		equal = equal && gc.invariantMatch_(hc, true)
	}

	return equal
}

func (g Graph) InvariantMatch(h Graph) bool {
	return g.invariantMatch_(h, false)
}

func (g Graph) BruteForceIsoCheck(h Graph) bool {
	if h.Size != g.Size {
		return false
	}

	for perm_i := uint64(0); perm_i < Factorial(g.Size); perm_i++ {
		perm := PermutationIth(g.Size, perm_i)

		match := true

		for i := uint8(0); i < g.Size && match; i++ {
			for j := i + 1; j < g.Size && match; j++ {
				if g.Adj.At(int(i), int(j)) != h.Adj.At(int(perm[i]), int(perm[j])) {
					match = false
				}
			}
		}

		if match {
			return true
		}
	}

	return false
}

func FilterIsosFunc(min, max uint64, f func(i uint64) Graph) []Graph {
	isos := make([]Graph, 0, (max-min)/3)

	for graph_i := min; graph_i < max; graph_i++ {
		isoFound := false
		g := f(graph_i)

		for j := 0; j < len(isos) && isoFound == false; j++ {
			isoFound = isos[j].InvariantMatch(g)
			if isoFound {
				isoFound = isos[j].BruteForceIsoCheck(g)
			}
		}

		if isoFound == false {
			isos = append(isos, g)
		}
	}

	return isos
}

func FilterIsosSlice(graphs []Graph) []Graph {
	min := uint64(0)
	max := uint64(len(graphs))
	closure := func(i uint64) Graph { return graphs[i] }

	return FilterIsosFunc(min, max, closure)
}

func FilterNGraphIsos(n uint8) []Graph {
	min := uint64(0)
	max := NumNGraphs(n)
	closure := func(i uint64) Graph { return AllNGraphsIth(n, i) }

	return FilterIsosFunc(min, max, closure)
}

func FilterIsosConcurr(min, max uint64, f func(i uint64) Graph) []Graph {
	isos := make([]Graph, 0, (max-min)/3)
	isos_chan := make(chan Graph)
	worker_chan := make(semaphore, 2)

	for graph_i := min; graph_i < max; graph_i++ {
		isos_len := uint64(len(isos))
		groupFunc := func(isos_len uint64) func(i uint64) Graph {
			return func(i uint64) Graph {
				if i == isos_len {
					return f(graph_i)
				} else {
					return isos[i]
				}
			}
		}(isos_len)

		go filterIsosConcurr_(min, isos_len+1, groupFunc, isos_chan, worker_chan)
	}

	for i := min; i < max; {
		select {
		case g := <-isos_chan:
			isos = append(isos, g)
		case <-worker_chan:
			i++
		}
	}

	return isos
}

func filterIsosConcurr_(min, max uint64, f func(i uint64) Graph, isos_chan chan Graph, worker_chan semaphore) {
	g := f(max - 1)
	isoFound := false

	for i := min; i < max-1 && isoFound == false; i++ {
		h := f(i)
		isoFound = h.InvariantMatch(g)
		if isoFound {
			isoFound = h.BruteForceIsoCheck(g)
		}
	}

	if isoFound == false {
		isos_chan <- g
	}

	worker_chan <- empty{}
}
