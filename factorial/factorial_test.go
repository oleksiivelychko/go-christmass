package factorial

import "testing"

func TestFactorial(t *testing.T) {
	result := Factorial(3)
	if result != 6 {
		t.Errorf("[Factorial(x int) int] -> %d != 6", result)
	}
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(3)
	}
}
