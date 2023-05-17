package goliner

// ForwardSubstitution substitution forward
// Important matrix should be low triangular
func (m Matrix) ForwardSubstitution(v Vector) Vector {
	n := len(m[0])

	answer := make(Vector, n)
	for i := 0; i < n; i++ {
		sum := 0.0
		for j := 0; j < i; j++ {
			sum += m[i][j] * answer[j]
		}
		answer[i] = (v[i] - sum) / m[i][i]
	}

	return answer
}

// BackwardSubstitution substitution backward
// Important matrix should be up triangular
func (m Matrix) BackwardSubstitution(v Vector) Vector {
	n := len(m[0])

	answer := make(Vector, n)
	for i := n - 1; i >= 0; i-- {
		sum := 0.0
		for j := n - 1; j > i; j-- {
			sum += m[i][j] * answer[i]
		}
		answer[i] = (v[i] - sum) / m[i][i]
	}

	return answer
}
