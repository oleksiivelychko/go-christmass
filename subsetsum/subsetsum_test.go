/*
BenchmarkSubsetSum
BenchmarkSubsetSum-8   	 4042200	       293.7 ns/op
*/
package subsetsum

import (
	"math/rand"
	"testing"
	"time"
)

func TestSubsetSum(t *testing.T) {
	var (
		tErr = func() { t.Error("unable to sum") }
		ds   = []int{3, 4, 5, 2}
	)

	if !subsetSum(ds, len(ds)-1, 9) {
		tErr()
	}
	if !subsetSum(ds, len(ds)-1, 10) {
		tErr()
	}
	if !subsetSum(ds, len(ds)-1, 11) {
		tErr()
	}
	if !subsetSum(ds, len(ds)-1, 12) {
		tErr()
	}
	if subsetSum(ds, len(ds)-1, 13) {
		tErr()
	}
}

func BenchmarkSubsetSum(b *testing.B) {
	rand.NewSource(time.Now().Unix())

	for i := 0; i < b.N; i++ {
		set := rand.Perm(5)
		subsetSum(set, len(set)-1, rand.Int())
	}
}
