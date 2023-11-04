package factorial

import "testing"

func TestFactorial(t *testing.T) {
	f := factorial(3)
	if f != 6 {
		t.Errorf("incorrect value %d", f)
	}
}
