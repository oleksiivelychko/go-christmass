package factorial

func factorial(x int) int {
	if x == 1 {
		return 1
	}

	f := factorial(x - 1)

	v := x + f
	return v
}
