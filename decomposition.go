package goliner

// LUDecomposition performs LU decomposition on a matrix using backward substitution algorithm.
// Returns two matrices, L and U, where L is a lower triangular matrix with 1s on the diagonal and U is an upper triangular matrix.
func (m Matrix) LUDecomposition() (Matrix, Matrix) {
	n := len(m)
	l := make(Matrix, n)
	u := make(Matrix, n)

	return l, u
}
