# goliner

goliner is a Go library for linear algebra computations, including matrix and vector operations. It provides efficient implementations of basic linear algebra operations such as dot product, matrix multiplication, and matrix-vector multiplication.

## Installation

To use goliner, you can simply install it using go get command:

```bash
go get github.com/kianooshaz/goliner
```

## Usage
Here is an example of how to use goliner:
```go
package main

import (
	"fmt"

	"github.com/kianooshaz/goliner"
)

func main() {
	// Create a vector
	vector := goliner.Vector{1, 2, 4, 5}

	// Insert a new number into the vector in sorted order
	vector.InsertSorted(3)

	// Compute the dot product of two vectors
	vector1 := goliner.Vector{1.0e-6, 2.0e-8, 3.0e-2}
	vector2 := goliner.Vector{4.0e-9, 5.0e-3, 6.0e-5}
	dotProduct := vector1.Dot(vector2)

	// Compute the dot product of a matrix and a vector
	matrix := goliner.Matrix{{1, 2, 3}, {4, 5, 6}}
	vector3 := goliner.Vector{1, 2, 3}
	dotProductMatrixVector := matrix.DotVector(vector3, true)

	// Compute the dot product of two matrices
	matrix1 := goliner.Matrix{{1, 2, 3}, {4, 5, 6}}
	matrix2 := goliner.Matrix{{7, 8}, {9, 10}, {11, 12}}
	dotProductMatrix := matrix1.Dot(matrix2, false)

	fmt.Println("InsertSorted Output: ", vector)
	fmt.Println("Dot Product Output: ", dotProduct)
	fmt.Println("Dot Vector Row Major Output: ", dotProductMatrixVector)
	fmt.Println("Dot Matrix Output: ", dotProductMatrix)
}
```
## API Reference
### Vector
A vector is represented as a slice of `float64` values.

`func (v Vector) Dot(other Vector) float64`

Computes the dot product of two vectors.

* `v`: the first vector
* `other`: the second vector
Returns the dot product of the two vectors.

`func (v Vector) Total() float64`

Calculates the sum of elements in the vector by repeatedly combining the two smallest values. Returns the sum of elements in the vector.

`func (v *Vector) InsertSorted(number float64)`

Inserts a new number into the vector in sorted order.
* `v`: a pointer to the vector
* `number`: the new number to be inserted
* 
### Matrix
A matrix is represented as a slice of `Vector` values.

`func (m Matrix) DotVector(v Vector, isRowMajor bool) Vector`
Computes the dot product of a matrix and a vector.

* `m`: the matrix
* `v`: the vector
* `isRowMajor`: a boolean value indicating whether the matrix is row-major (true) or column-major (false)

Returns the dot product of the matrix and the vector.

`func (m Matrix) Dot(other Matrix, isRowMajor bool) Matrix`
Computes the dot product of two matrices.

* `m`: the first matrix
* `other`: the second matrix
* `isRowMajor`: a boolean




