package fibonacci

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	f := Fibonacci()
	for i := 0; i < 10; i++ {
		sequenceNum := f()
		if i == 0 && sequenceNum != 0 {
			t.Errorf("[func Fibonacci() func() int] -> %d != 0", sequenceNum)
		}
		if i == 1 && sequenceNum != 1 {
			t.Errorf("[func Fibonacci() func() int] -> %d != 1", sequenceNum)
		}
		if i == 2 && sequenceNum != 1 {
			t.Errorf("[func Fibonacci() func() int] -> %d != 1", sequenceNum)
		}
		if i == 3 && sequenceNum != 2 {
			t.Errorf("[func Fibonacci() func() int] -> %d != 2", sequenceNum)
		}
		if i == 4 && sequenceNum != 3 {
			t.Errorf("[func Fibonacci() func() int] -> %d != 3", sequenceNum)
		}
		if i == 5 && sequenceNum != 5 {
			t.Errorf("[func Fibonacci() func() int] -> %d != 5", sequenceNum)
		}
		if i == 6 && sequenceNum != 8 {
			t.Errorf("[func Fibonacci() func() int] -> %d != 8", sequenceNum)
		}
		if i == 7 && sequenceNum != 13 {
			t.Errorf("[func Fibonacci() func() int] -> %d != 13", sequenceNum)
		}
		if i == 8 && sequenceNum != 21 {
			t.Errorf("[func Fibonacci() func() int] -> %d != 21", sequenceNum)
		}
		if i == 9 && sequenceNum != 34 {
			t.Errorf("[func Fibonacci() func() int] -> %d != 34", sequenceNum)
		}
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci()
	}
}
