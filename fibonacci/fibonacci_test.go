package fibonacci

import "testing"

func TestFibonacci(t *testing.T) {
	f := fibonacci()

	tErr := func(num, expect int) { t.Errorf("[func fibonacci() func() int] -> %d != %d", num, expect) }

	for i := 0; i < 10; i++ {
		num := f()

		if i == 0 && num != 0 {
			tErr(num, 0)
		}
		if i == 1 && num != 1 {
			tErr(num, 1)
		}
		if i == 2 && num != 1 {
			tErr(num, 1)
		}
		if i == 3 && num != 2 {
			tErr(num, 2)
		}
		if i == 4 && num != 3 {
			tErr(num, 3)
		}
		if i == 5 && num != 5 {
			tErr(num, 5)
		}
		if i == 6 && num != 8 {
			tErr(num, 8)
		}
		if i == 7 && num != 13 {
			tErr(num, 13)
		}
		if i == 8 && num != 21 {
			tErr(num, 21)
		}
		if i == 9 && num != 34 {
			tErr(num, 34)
		}
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci()
	}
}
