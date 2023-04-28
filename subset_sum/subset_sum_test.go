package subset_sum

import (
	"math/rand"
	"testing"
	"time"
)

func TestSubsetSum(t *testing.T) {
	set := []int{3, 4, 5, 2}

	isSubsetSum := SubsetSum(set, len(set)-1, 9)
	if isSubsetSum != true {
		t.Errorf("[func SubsetSum(set []int, n, sum int) bool] -> %t != true", isSubsetSum)
	}

	isSubsetSum = SubsetSum(set, len(set)-1, 10)
	if isSubsetSum != true {
		t.Errorf("[func SubsetSum(set []int, n, sum int) bool] -> %t != true", isSubsetSum)
	}

	isSubsetSum = SubsetSum(set, len(set)-1, 11)
	if isSubsetSum != true {
		t.Errorf("[func SubsetSum(set []int, n, sum int) bool] -> %t != true", isSubsetSum)
	}

	isSubsetSum = SubsetSum(set, len(set)-1, 12)
	if isSubsetSum != true {
		t.Errorf("[func SubsetSum(set []int, n, sum int) bool] -> %t != true", isSubsetSum)
	}

	isSubsetSum = SubsetSum(set, len(set)-1, 13)
	if isSubsetSum != false {
		t.Errorf("[func SubsetSum(set []int, n, sum int) bool] -> %t != false", isSubsetSum)
	}
}

func BenchmarkSubsetSum(b *testing.B) {
	rand.NewSource(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		set := rand.Perm(5)
		SubsetSum(set, len(set)-1, rand.Int())
	}
}
