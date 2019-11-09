package graph_test

import (
	"graphtheory/graph"
	"testing"
)

func BenchmarkPermutationIth(b *testing.B) {
	for n := uint8(2); n < 10; n++ {
		for i := uint64(0); i < 8*7*6*5*4*3*2; i++ {
			graph.PermutationIth(8, i)
		}
	}
}

func TestEdgeSet(t *testing.T) {
	n := uint8(3)

	// Expected set out of order on purpose, to prevent order specific implementation
	expected := [3]graph.Edge{{1, 2}, {0, 2}, {1, 0}}
	edges := graph.EdgeSet(n)

	// Check that all the expected edges are in the returned edges
	for _, v1 := range expected {
		inSet := false
		for _, v2 := range edges {
			if v1.Equal(v2) {
				inSet = true
				break
			}
		}

		if inSet == false {
			t.Errorf("EdgeSet for size %d missing an element: expected an edge set that contains %v, got edge set %v", n, v1, edges)
		}
	}

	// Check that all the returned edges were expected
	for _, v1 := range edges {
		inSet := false
		for _, v2 := range expected {
			if v1.Equal(v2) {
				inSet = true
				break
			}
		}

		if inSet == false {
			t.Errorf("EdgeSet for size %d has an unexpected element: got the edge set %v, which contains the unexpected edge %v", n, edges, v1)
		}
	}
}

// Compare two slices irrespective of order, i.e., as sets
func setEquality(a, b []uint8) bool {
	if len(a) != len(b) {
		return false
	}

	// Check that everything in a is in b
	for _, v1 := range a {
		inSet := false
		for _, v2 := range b {
			if v1 == v2 {
				inSet = true
				break
			}
		}

		if inSet == false {
			return false
		}
	}

	// Check that everything in b is in a
	for _, v1 := range b {
		inSet := false
		for _, v2 := range a {
			if v1 == v2 {
				inSet = true
				break
			}
		}

		if inSet == false {
			return false
		}
	}

	return true
}

func compareSetofSets(t *testing.T, n uint8, expected, received [][]uint8, functionName, objectName string) {
	// Check that all the expected sets are in the returned set of sets
	for _, v1 := range expected {
		inSet := false
		for _, v2 := range received {
			if setEquality(v1, v2) {
				inSet = true
				break
			}
		}

		if inSet == false {
			t.Errorf("%s for size %d missing an element: expected %s containing %v, got %v", functionName, n, objectName, v1, received)
		}
	}

	// Check that all the returned sets were expected
	for _, v1 := range received {
		inSet := false
		for _, v2 := range expected {
			if setEquality(v1, v2) {
				inSet = true
				break
			}
		}

		if inSet == false {
			t.Errorf("%s for size %d has an unexpected element: got %s %v, which contains the unexpected set %v", functionName, n, objectName, received, v1)
		}
	}
}

func TestPowerset(t *testing.T) {
	n := uint8(3)

	// Expected set purposefully out of order, to prevent order specific implementation
	expected := [][]uint8{{}, {2, 0, 1}, {1}, {2}, {1, 0}, {0, 2}, {1, 2}, {0}}
	received := graph.Powerset(n)

	compareSetofSets(t, n, expected, received, "Powerset", "powerset")
}

func TestPowersetIth(t *testing.T) {
	n := uint8(3)

	// Expected set purposefully out of order, to prevent order specific implementation
	expected := [][]uint8{{}, {2, 0, 1}, {1}, {2}, {1, 0}, {0, 2}, {1, 2}, {0}}

	// Build slice of sets for all i from 0 to 2**n, then compare
	// This also prevents order specific implementation, but ensures coverage of all possible sets
	received := make([][]uint8, 1<<n)

	for i := range received {
		received[i] = graph.PowersetIth(n, uint64(i))
	}

	compareSetofSets(t, n, expected, received, "PowersetIth", "powerset")
}

func TestFactorial(t *testing.T) {
	expected := []uint64{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880, 3628800, 39916800, 479001600}

	for i := range expected {
		v := graph.Factorial(uint8(i))
		if v != expected[i] {
			t.Errorf("Factorial: expected %d, got %d", expected[i], v)
		}
	}
}

func TestPermutationIth(t *testing.T) {
	n := uint8(3)

	// Expected set purposefully out of order (as compared to implementation)
	// This prevents order specific implementation
	expected := [][]uint8{{1, 0, 2}, {0, 1, 2}, {0, 2, 1}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}}

	// Build slice of permutations for all i from 0 to n!, then compare
	// This also prevents order specific implementation, but ensures coverage of all possible permutations
	received := make([][]uint8, graph.Factorial(n))

	for i := range received {
		received[i] = graph.PermutationIth(n, uint64(i))
	}

	compareSetofSets(t, n, expected, received, "PermutationIth", "permutation")
}

func TestNumNGraphs(t *testing.T) {
	expected := []uint64{1, 1, 2, 8, 64, 1024, 32768}

	for i := range expected {
		v := graph.NumNGraphs(uint8(i))
		if v != expected[i] {
			t.Errorf("NumNGraphs: expected %d, got %d", expected[i], v)
		}
	}
}

func compareSetofGraphs(t *testing.T, n uint8, expected, received []graph.Graph, functionName string) {
	// Check that all the expected graphs are in the returned set of graphs
	for _, v1 := range expected {
		inSet := false
		for _, v2 := range received {
			if v1.InvariantMatch(v2) {
				inSet = true
				break
			}
		}

		if inSet == false {
			t.Errorf("%s for size %d missing an element: expected set containing %v, got %v", functionName, n, v1, received)
		}
	}

	// Check that all the returned graphs were expected
	for _, v1 := range received {
		inSet := false
		for _, v2 := range expected {
			if v1.InvariantMatch(v2) {
				inSet = true
				break
			}
		}

		if inSet == false {
			t.Errorf("%s for size %d has an unexpected element: got %v, which contains the unexpected graph %v", functionName, n, received, v1)
		}
	}
}

func TestAllNGraphs(t *testing.T) {
	n := uint8(3)

	// Expected edges out of order to prevent order specific implementation
	expectedEdges := [][]graph.Edge{
		{},
		{{0, 1}},
		{{0, 2}},
		{{1, 2}},
		{{0, 1}, {0, 2}},
		{{0, 1}, {1, 2}},
		{{0, 1}, {0, 2}, {1, 2}},
		{{0, 2}, {1, 2}},
	}

	expected := make([]graph.Graph, 0, graph.NumNGraphs(n))

	for i := range expectedEdges {
		expected = append(expected, graph.NewGraph(n).AddEdges(expectedEdges[i]))
	}

	received := graph.AllNGraphs(n)

	compareSetofGraphs(t, n, expected, received, "AllNGraphs")
}

func TestAllNGraphsIth(t *testing.T) {
	n := uint8(3)

	// Expected edges out of order to prevent order specific implementation
	expectedEdges := [][]graph.Edge{
		{},
		{{0, 1}},
		{{0, 2}},
		{{1, 2}},
		{{0, 1}, {0, 2}},
		{{0, 1}, {1, 2}},
		{{0, 1}, {0, 2}, {1, 2}},
		{{0, 2}, {1, 2}},
	}

	expected := make([]graph.Graph, 0, graph.NumNGraphs(n))

	for i := range expectedEdges {
		expected = append(expected, graph.NewGraph(n).AddEdges(expectedEdges[i]))
	}

	// Build slice of n-graphs for all i from 0 to NumNGraphs, then compare
	// This also prevents order specific implementation, but ensures coverage of all possible n-graphs
	received := make([]graph.Graph, 0, graph.NumNGraphs(n))

	for i := range expected {
		received = append(received, graph.AllNGraphsIth(n, uint64(i)))
	}

	compareSetofGraphs(t, n, expected, received, "AllNGraphsIth")
}
