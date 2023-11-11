package fibonacci

import "testing"

func TestFibonacci(t *testing.T) {
	var f = fibonacci()

	for i := 0; i < 10; i++ {
		num := f()

		if i == 0 && num != 0 {
			t.Errorf("expected %d, got %d", 0, num)
		}

		if i == 1 && num != 1 {
			t.Errorf("expected %d, got %d", 1, num)
		}

		if i == 2 && num != 1 {
			t.Errorf("expected %d, got %d", 1, num)
		}

		if i == 3 && num != 2 {
			t.Errorf("expected %d, got %d", 2, num)
		}

		if i == 4 && num != 3 {
			t.Errorf("expected %d, got %d", 3, num)
		}

		if i == 5 && num != 5 {
			t.Errorf("expected %d, got %d", 5, num)
		}

		if i == 6 && num != 8 {
			t.Errorf("expected %d, got %d", 8, num)
		}

		if i == 7 && num != 13 {
			t.Errorf("expected %d, got %d", 13, num)
		}

		if i == 8 && num != 21 {
			t.Errorf("expected %d, got %d", 21, num)
		}

		if i == 9 && num != 34 {
			t.Errorf("expected %d, got %d", 34, num)
		}
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci()
	}
}
