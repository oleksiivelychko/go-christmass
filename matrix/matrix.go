package matrix

type IMatrix interface {
	int | int8 | int16 | int32 | int64
}

func IntMatrix[M IMatrix](rows M, cols M) [][]M {
	cells := make([][]M, rows)
	for i := range cells {
		cells[i] = make([]M, cols)
	}

	return cells
}
