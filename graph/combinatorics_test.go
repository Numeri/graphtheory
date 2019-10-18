package graph

import "testing"

func BenchmarkPermutationIth(b *testing.B) {
	for n := uint8(2); n < 10; n++ {
		for i := uint64(0); i < 8*7*6*5*4*3*2; i++ {
			PermutationIth(8, i)
		}
	}
}
