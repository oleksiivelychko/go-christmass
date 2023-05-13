package factorial

import "github.com/oleksiivelychko/go-code-helpers/intertype"

func Factorial[T intertype.INumber](x T) T {
	if x == 1 {
		return 1
	}

	f := Factorial(x - 1)

	returnValue := x + f
	return returnValue
}
