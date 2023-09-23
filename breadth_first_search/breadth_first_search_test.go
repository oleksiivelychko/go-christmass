package breadth_first_search

import "testing"

/*
V(robin) 			-> V(rabbit)
V(robin) 			-> V(piglet)
V(robin) 			-> V(winnie-the-pooh)
V(piglet)			-> V(eeyore)
V(winnie-the-pooh)	-> V(kanga)
V(winnie-the-pooh)	-> V(roo)
V(rabbit) 			-> V(tiger)
V(rabbit) 			-> V(winnie-the-pooh)
V(tigger)			-> V(owl)
*/
var unweightedGraph = map[string][]string{
	"robin":           {"rabbit", "piglet", "winnie-the-pooh"},
	"piglet":          {"eeyore"},
	"rabbit":          {"tigger", "winnie-the-pooh"},
	"winnie-the-pooh": {"kanga", "roo"},
	"eeyore":          {},
	"tigger":          {"owl"},
	"kanga":           {},
	"roo":             {},
	"owl":             {},
	"gopher":          {},
}

func TestBreadthFirstSearch(t *testing.T) {
	node := BreadthFirstSearch(unweightedGraph, "owl", "robin")
	if node == "" {
		t.Error("unable to find node")
	}
}

func BenchmarkBreadthFirstSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BreadthFirstSearch(unweightedGraph, "owl", "robin")
	}
}
