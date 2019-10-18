package graph

import "testing"

func BenchmarkFilterNGraphIsos(b *testing.B) {
	n := uint8(6)
	FilterNGraphIsos(n)
}
