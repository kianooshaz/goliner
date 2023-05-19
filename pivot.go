package goliner

type PivotFunc func(Matrix, int) Matrix

func NonePivot(matrix Matrix, i int) Matrix {
	return matrix
}

func PartialPivot(matrix Matrix, i int) Matrix {
	// TODO
	return matrix
}

func FullPivot(matrix Matrix, i int) Matrix {
	// TODO
	return matrix
}
