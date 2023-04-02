package goliner

import (
	"sort"
)

// Define a type alias "vector" for a slice of float64s.
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
	sort.Float64s(v)

	for len(v) > 1 {
		sum := v[0] + v[1]
		if v[1] != 0 && sum != v[0] {
			v = v[1:]
			v[0] = sum
		} else {
			num := v[0]
			v = v[1:]
			v.InsertSorted(num)
		}
	}

	return v[0]
}

// InsertSorted inserts a new number into the vector in sorted order.
func (v *Vector) InsertSorted(number float64) {
	// If the number is less than or equal to the first element, insert it at the beginning.
	if number <= (*v)[0] {
		*v = append([]float64{number}, (*v)...)
		return
	}

	// Otherwise, find the position where the number should be inserted and insert it there.
	for i, nv := range *v {
		if number <= nv {
			(*v) = append((*v)[:i], append([]float64{number}, (*v)[i:]...)...)
			return
		}
	}

	// If the number is greater than all the elements, append it to the end.
	*v = append(*v, number)
}
