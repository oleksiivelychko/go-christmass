package factorial

import "testing"

func TestFactorial(t *testing.T) {
	if f := factorial(3); f != 6 {
		t.Errorf("expected %d, got %d", 6, f)
	}
}
