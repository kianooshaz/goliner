package goliner

import "math"

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

// LUDecompositionInPlace ...
func (m Matrix) LUDecompositionInPlace(pivot PivotFunc) Matrix {
	rowLen := len(m)
	columnLen := len(m[0])
	n := int(math.Max(float64(rowLen), float64(columnLen)))
	matrix := m

	for i := 0; i < n; i++ {
		matrix = pivot(matrix, i)
		matrix[i][i] = 1 / matrix[i][i]

		for j := i + 1; j < n; j++ {
			matrix[i][j] *= matrix[i][i]
		}

		for j := i + 1; j < n; j++ {
			matrix[j][i] *= -1
			for k := i + 1; k < n; k++ {
				matrix[j][k] += matrix[j][i] * matrix[i][k]
			}
		}
	}

	return matrix
}
