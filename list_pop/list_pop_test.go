package list_pop

import "testing"

var listStr = []string{"one", "two", "three"}
var listInt = []int{1, 2, 3}

func TestListPop(t *testing.T) {
	list := ListPop(listStr, -1)
	if len(list) != 2 {
		t.Errorf("[func ListPop[L IList](array []L, index int) []L] -> %d != '2'", len(list))
	}
	if list[0] != "one" {
		t.Errorf("[func ListPop[L IList](array []L, index int) []L] -> %s != 'one'", list[0])
	}
	if list[1] != "two" {
		t.Errorf("[func ListPop[L IList](array []L, index int) []L] -> %s != 'two'", list[1])
	}
}

func TestListPopByIndex(t *testing.T) {
	list := ListPop(listStr, 1)
	if len(list) != 2 {
		t.Errorf("[func ListPop[L IList](array []L, index int) []L] -> %d != '2'", len(list))
	}
	if list[0] != "one" {
		t.Errorf("[func ListPop[L IList](array []L, index int) []L] -> %s != 'one'", list[0])
	}
	if list[1] != "three" {
		t.Errorf("[func ListPop[L IList](array []L, index int) []L] -> %s != 'three'", list[1])
	}
}

func BenchmarkStringPop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ListPop(listStr, 0)
	}
}

func TestIntPop(t *testing.T) {
	list := ListPop(listInt, -1)
	if len(list) != 2 {
		t.Errorf("[func ListPop[L IList](array []L, index int) []L] -> %d != '2'", len(list))
	}
	if list[0] != 1 {
		t.Errorf("[func ListPop[L IList](array []L, index int) []L] -> %d != 1", list[0])
	}
	if list[1] != 2 {
		t.Errorf("[func ListPop[L IList](array []L, index int) []L] -> %d != 2", list[1])
	}
}

func TestIntPopByIndex(t *testing.T) {
	list := ListPop(listInt, 1)
	if len(list) != 2 {
		t.Errorf("[func ListPop[L IList](array []L, index int) []L] -> %d != '2'", len(list))
	}
	if list[0] != 1 {
		t.Errorf("[func ListPop[L IList](array []L, index int) []L] -> %d != 1", list[0])
	}
	if list[1] != 3 {
		t.Errorf("[func ListPop[L IList](array []L, index int) []L] -> %d != 3", list[1])
	}
}

func BenchmarkIntPop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ListPop(listInt, 0)
	}
}
