package graph_test

import (
	"graphtheory/graph"
	"testing"
)

func BenchmarkFilterNGraphIsos(b *testing.B) {
	n := uint8(6)
	graph.FilterNGraphIsos(n)
}
