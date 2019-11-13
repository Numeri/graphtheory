package graph

import (
	"sync"
	"sync/atomic"
)

func CountAllConnectedNGraphs(n uint8) (connected, nonconnected uint64) {
	connected = 0
	nonconnected = 0

	f := func(g Graph) {
		if g.IsConnected() {
			atomic.AddUint64(&connected, 1)
		} else {
			atomic.AddUint64(&nonconnected, 1)
		}
	}

	ConcurrentNGraphMap(n, f)

	return connected, nonconnected
}

func ConcurrentNGraphMap(n uint8, f func(g Graph)) {
	numNGraphs := NumNGraphs(n)

	sem := make(semaphore, 100)
	var wg sync.WaitGroup

	for i := uint64(0); i < numNGraphs; i++ {
		sem <- empty{}
		wg.Add(1)
		go func(i uint64) {
			defer wg.Done()
			g := AllNGraphsIth(n, i)
			f(g)
			<-sem
		}(i)
	}

	wg.Wait()
}
