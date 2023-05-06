/*
BenchmarkSubsetSum
BenchmarkSubsetSum-8   	 4042200	       293.7 ns/op
*/
package subset_sum

import (
	"math/rand"
	"testing"
	"time"
)

func TestSubsetSum(t *testing.T) {
	set := []int{3, 4, 5, 2}

	if !SubsetSum(set, len(set)-1, 9) {
		t.Error("calculation is wrong")
	}
	if !SubsetSum(set, len(set)-1, 10) {
		t.Error("calculation is wrong")
	}
	if !SubsetSum(set, len(set)-1, 11) {
		t.Error("calculation is wrong")
	}
	if !SubsetSum(set, len(set)-1, 12) {
		t.Error("calculation is wrong")
	}
	if SubsetSum(set, len(set)-1, 13) {
		t.Error("calculation is wrong")
	}
}

func BenchmarkSubsetSum(b *testing.B) {
	rand.NewSource(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		set := rand.Perm(5)
		SubsetSum(set, len(set)-1, rand.Int())
	}
}
