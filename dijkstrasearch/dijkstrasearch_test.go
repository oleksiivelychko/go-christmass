package dijkstrasearch

import (
	"math"
	"testing"
)

/*
V(start)	-> V(a) 		= E(6)
V(start) 	-> V(b) 		= E(2)
V(a) 		-> V(finish)	= E(1)
V(b) 		-> V(a) 		= E(3)
V(b) 		-> V(finish) 	= E(5)
*/
var graph = map[string]map[string]float64{
	"start": {
		"a": 6,
		"b": 2,
	},
	"a": {
		"finish": 1,
	},
	"b": {
		"a":      3,
		"finish": 5,
	},
	"finish": {},
}

var costs = map[string]float64{
	"a":      6,
	"b":      2,
	"finish": math.Inf(1),
}

func TestDijkstraSearch(t *testing.T) {
	c := search(graph, costs)
	if c != 6 {
		t.Errorf("incorrect cost %f", c)
	}
}

func BenchmarkDijkstraSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		search(graph, costs)
	}
}
