package graph

import (
	"fmt"
	"sync/atomic"
)

type empty = struct{}
type semaphore = chan empty

func findAllConnectedNGraphs() {
	connected := uint64(0)
	nonconnected := uint64(0)
	n := uint8(7)
	numNGraphs := NumNGraphs(n)

	f := func(g Graph) {
		if g.IsConnected() {
			atomic.AddUint64(&connected, 1)
		} else {
			atomic.AddUint64(&nonconnected, 1)
		}
		if (connected+nonconnected)%50000 == 0 {
			fmt.Printf("%d/%d =\t%v\n", connected+nonconnected, numNGraphs, 100*float64(connected+nonconnected)/float64(numNGraphs))
		}
	}

	ConcurrentNGraphs(n, f)

	fmt.Println("Connected graphs:    ", connected)
	fmt.Println("Nonconnected graphs: ", nonconnected)
	fmt.Println("Total graphs:        ", connected+nonconnected)
	fmt.Println("Total number n-graphs", numNGraphs)
}

func ConcurrentNGraphs(n uint8, f func(g Graph)) {
	numNGraphs := NumNGraphs(n)

	sem := make(semaphore, 2000)

	for i := uint64(0); i < numNGraphs; i++ {
		sem <- empty{}
		go func(i uint64) {
			g := AllNGraphsIth(n, i)

			f(g)
			<-sem
		}(i)
	}

	for len(sem) != 0 {
	}
}
