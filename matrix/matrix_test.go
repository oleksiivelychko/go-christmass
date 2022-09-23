package matrix

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestIntMatrix(t *testing.T) {
	matrix := IntMatrix(5, 5)

	lenMatrix := len(matrix)
	capMatrix := cap(matrix)

	if lenMatrix != 5 {
		t.Errorf("[func IntMatrix[M IMatrix](rows M, cols M) [][]M -> %d != '5'", lenMatrix)
	}
	if capMatrix != 5 {
		t.Errorf("[func IntMatrix[M IMatrix](rows M, cols M) [][]M -> %d != '5'", capMatrix)
	}

	fmt.Printf("%#v\n", matrix)
}

func BenchmarkIntMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMatrix(rand.Intn(100), rand.Intn(100))
	}
}
