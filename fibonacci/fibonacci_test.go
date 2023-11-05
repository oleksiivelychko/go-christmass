package fibonacci

import "testing"

func TestFibonacci(t *testing.T) {
	var (
		tErr = func(e, g int) { t.Errorf("expected %d, got %d", e, g) }
		f    = fibonacci()
	)

	for i := 0; i < 10; i++ {
		num := f()

		if i == 0 && num != 0 {
			tErr(0, num)
		}
		if i == 1 && num != 1 {
			tErr(1, num)
		}
		if i == 2 && num != 1 {
			tErr(1, num)
		}
		if i == 3 && num != 2 {
			tErr(2, num)
		}
		if i == 4 && num != 3 {
			tErr(3, num)
		}
		if i == 5 && num != 5 {
			tErr(5, num)
		}
		if i == 6 && num != 8 {
			tErr(8, num)
		}
		if i == 7 && num != 13 {
			tErr(13, num)
		}
		if i == 8 && num != 21 {
			tErr(21, num)
		}
		if i == 9 && num != 34 {
			tErr(34, num)
		}
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci()
	}
}
