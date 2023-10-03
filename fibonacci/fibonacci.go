package fibonacci

func fibonacci() func() int {
	x, y := 0, 1

	return func() int {
		defer func() {
			x = x + y
			x, y = y, x
		}()

		return x
	}
}
