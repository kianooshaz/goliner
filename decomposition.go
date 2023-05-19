package goliner

import "math"

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

func (m Matrix) CholeskyDecomposition() Matrix {
	matrix := make(Matrix, len(m))
	for i := range matrix {
		matrix[i] = make(Vector, len(m[0]))
	}

	for i := 0; i < len(m); i++ {
		for j := 0; j <= i; j++ {
			if i == j {
				sum := 0.0
				for k := 0; k < j; k++ {
					sum += matrix[j][k] * matrix[j][k]
				}
				matrix[i][j] = math.Sqrt(m[i][j] - sum)
			} else {
				sum := 0.0
				for k := 0; k < j; k++ {
					sum += matrix[i][k] * matrix[j][k]
				}
				matrix[i][j] = (m[i][j] - sum) / matrix[j][j]
			}
		}
	}

	return matrix
}
