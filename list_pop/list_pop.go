package list_pop

type IList interface {
	int | int8 | int16 | int32 | int64 | string
}

func ListPop[L IList](array []L, index int) []L {
	if index < 0 {
		// pop off last item
		return array[:len(array)-1]
	} else {
		// ...or by index
		return append(array[:index], array[index+1:]...)
	}
}
