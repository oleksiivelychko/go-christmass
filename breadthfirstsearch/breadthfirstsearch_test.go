package breadthfirstsearch

import "testing"

var ds = map[string][]string{
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
	if search(ds, "owl", "robin") == "" {
		t.Error("unable to find")
	}
}

func BenchmarkBreadthFirstSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		search(ds, "owl", "robin")
	}
}
