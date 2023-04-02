package goliner

type Matrix []Vector

// DotVector computes the dot product of a matrix and a vector.
func (m Matrix) DotVector(v Vector, isRowMajor bool) Vector {
	if len(m[0]) != len(v) {
		panic("Matrix and vector must have compatible dimensions")
	}

	result := make(Vector, len(m))

	if isRowMajor {
		for i, mv := range m {
			result[i] = mv.Dot(v)
		}
	} else {
		products := make(Matrix, len(m))
		for i := range m {
			products[i] = make(Vector, len(v))
		}

		for j, val := range v {
			for i := range m {
				products[i][j] = m[i][j] * val
			}
		}

		for i := range m {
			result[i] = products[i].Total()
		}
	}

	return result
}

// Dot computes the dot product of two matrices.
func (m Matrix) Dot(other Matrix, isRowMajor bool) Matrix {
	if len(m[0]) != len(other) {
		panic("matrices must have compatible dimensions")
	}

	result := make(Matrix, len(m))
	for i := range result {
		result[i] = make(Vector, len(other[0]))
	}

	if isRowMajor {
		for i, row := range m {
			newRow := make(Vector, len(other[0]))
			for j := range other[0] {
				products := make(Vector, len(row))
				for k, val := range row {
					products[k] = val * other[k][j]
				}
				newRow[j] = products.Total()
			}
			result[i] = newRow
		}
	} else {
		for j := range other[0] {
			column := make(Vector, len(other))
			for i := range other {
				column[i] = other[i][j]
			}
			product := m.DotVector(column, isRowMajor)
			for i, val := range product {
				result[i][j] = val
			}
		}
	}

	return result
}
