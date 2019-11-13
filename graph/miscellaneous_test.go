package graph_test

import (
	"graphtheory/graph"
	"testing"
)

func TestCountAllConnectedNGraphs(t *testing.T) {
	n := uint8(6)
	expected := uint64(26704)
	expected_non := graph.NumNGraphs(n) - expected

	connected, nonconnected := graph.CountAllConnectedNGraphs(n)

	if connected != expected {
		t.Errorf("CountAllConnectedNGraphs: got %d connected graphs of size %d, expected %d", connected, n, expected)
	}

	if nonconnected != expected_non {
		t.Errorf("CountAllConnectedNGraphs: got %d non-connected graphs of size %d, expected %d", nonconnected, n, expected_non)
	}
}
