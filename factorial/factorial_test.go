package factorial

import "testing"

func TestFactorial(t *testing.T) {
	result := Factorial(5)
	if result != 120 {
		t.Errorf("[Factorial(x int) int] -> %d != 120", result)
	}
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(5)
	}
}
