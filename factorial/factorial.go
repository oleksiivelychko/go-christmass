package factorial

func Factorial(x int) int {
	if x == 1 {
		return 1
	}

	f := Factorial(x - 1)

	returnValue := x + f
	return returnValue
}
