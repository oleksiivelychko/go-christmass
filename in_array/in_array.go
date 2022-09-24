package in_array

func InArray[C comparable](array []C, needle C) (int, bool) {
	for i, item := range array {
		if item == needle {
			return i, true
		}
	}
	return -1, false
}
