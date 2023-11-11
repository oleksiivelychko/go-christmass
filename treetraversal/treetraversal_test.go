package treetraversal

import "testing"

func tree() *node {
	var t = &node{value: 0, left: nil, right: nil}
	t.left = &node{value: 1, left: nil, right: nil}
	t.left.left = &node{value: 2, left: nil, right: nil}
	t.left.left.left = &node{value: 3, left: nil, right: nil}
	t.left.left.right = &node{value: 4, left: nil, right: nil}
	t.left.right = &node{value: 5, left: nil, right: nil}
	t.right = &node{value: 6, left: nil, right: nil}
	return t
}

func TestTreeTraversal(t *testing.T) {
	if sum := traverse(tree()); sum != 21 {
		t.Errorf("expected %d, got %d", 21, sum)
	}
}

func BenchmarkTreeTraversal(b *testing.B) {
	t := tree()
	for i := 0; i < b.N; i++ {
		traverse(t)
	}
}
