package factorial

import "testing"

func TestFactorial(t *testing.T) {
	factorial := Factorial(3)
	if factorial != 6 {
		t.Errorf("value %d is wrong", factorial)
	}
}
