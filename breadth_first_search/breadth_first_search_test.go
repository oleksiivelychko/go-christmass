package breadth_first_search

import "testing"

/*
V(you) 				-> V(rabbit)
V(you) 				-> V(piglet)
V(you) 				-> V(winnie-the-pooh)
V(piglet)			-> V(gopher)
V(piglet) 			-> V(eeyore)
V(rabbit) 			-> V(tiger)
V(winnie-the-pooh)	-> V(kanga)
V(winnie-the-pooh)	-> V(roo)
V(tigger)			-> V(owl)
V(tigger)			-> V(gopher)
*/
var unweightedGraph = map[string][]string{
	"you":             {"rabbit", "piglet", "winnie-the-pooh"},
	"piglet":          {"gopher", "eeyore"},
	"rabbit":          {"tigger"},
	"winnie-the-pooh": {"kanga", "roo"},
	"eeyore":          {},
	"tigger":          {"owl", "gopher"},
	"kanga":           {},
	"roo":             {},
	"owl":             {},
	"gopher":          {},
}

func TestBreadthFirstSearch(t *testing.T) {
	node := BreadthFirstSearch(unweightedGraph, "gopher", "you")
	if node == "" {
		t.Error("unable to find node")
	}
}

func BenchmarkBreadthFirstSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BreadthFirstSearch(unweightedGraph, "gopher", "you")
	}
}
