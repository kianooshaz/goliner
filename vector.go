package goliner

import (
	"sort"
)

// Vector Define a type alias "vector" for a slice of float64.
type Vector []float64

// Dot computes the dot product of two vectors.
func (v Vector) Dot(other Vector) float64 {
	if len(v) == 0 || len(v) != len(other) {
		panic("vectors must have the same non-zero length")
	}

	products := make(Vector, len(v))
	for i, val := range v {
		products[i] = val * other[i]
	}

	return products.Total()
}

// Total calculates the sum of elements in the vector by repeatedly combining the two smallest values
func (v Vector) Total() float64 {
	vector := v

	sort.Float64s(vector)

	for len(vector) > 1 {
		sum := vector[0] + vector[1]
		if vector[1] != 0 && sum != vector[0] {
			vector = vector[1:]
			vector[0] = sum
		} else {
			num := vector[0]
			vector = vector[1:].InsertSorted(num)
		}
	}

	return vector[0]
}

// InsertSorted inserts a new number into the vector in sorted order.
func (v Vector) InsertSorted(number float64) Vector {

	// If the number is less than or equal to the first element, insert it at the beginning.
	if number <= v[0] {
		return append([]float64{number}, v...)
	}

	// Otherwise, find the position where the number should be inserted and insert it there.
	for i, nv := range v {
		if number <= nv {
			return append(v[:i], append([]float64{number}, v[i:]...)...)
		}
	}

	// If the number is greater than all the elements, append it to the end.
	return append(v, number)
}
