package main

import (
	"fmt"
	"github.com/kianooshaz/goliner"
)

func main() {
	fmt.Println("Epsilon of machine:", goliner.Epsilon)

	vector := goliner.Vector{1, 2, 4, 5}
	fmt.Println("Vector :", vector)
	vector.InsertSorted(3)
	fmt.Println("Insert ", 3)
	fmt.Println("InsertSorted Output: ", vector)

	vector1 := goliner.Vector{1.0e-6, 2.0e-8, 3.0e-2}
	vector2 := goliner.Vector{4.0e-9, 5.0e-3, 6.0e-5}
	fmt.Println("Vector 1:", vector1)
	fmt.Println("Vector 2:", vector2)
	fmt.Println("Dot Product Output: ", vector1.Dot(vector2))

	matrix := goliner.Matrix{{1, 2, 3}, {4, 5, 6}}
	vector3 := goliner.Vector{1, 2, 3}
	fmt.Println("Matrix:", matrix)
	fmt.Println("Vector:", vector3)
	fmt.Println("Dot Vector Row Major Output: ", matrix.DotVector(vector3, true))
	fmt.Println("Dot Vector Column Major Output: ", matrix.DotVector(vector3, false))

	matrix1 := goliner.Matrix{{1, 2}, {3, 4}}
	matrix2 := goliner.Matrix{{5, 6}, {7, 8}}
	fmt.Println("Matrix 1:", matrix1)
	fmt.Println("Matrix 2:", matrix2)
	fmt.Println("Dot Matrix Row Major Output: ", matrix1.Dot(matrix2, true))
	fmt.Println("Dot Matrix Column Major Output: ", matrix1.Dot(matrix2, false))

	fmt.Println("LU Decomposition in place of matrix:")
	m := goliner.Matrix{
		goliner.Vector{12, 6, 5},
		goliner.Vector{7, 4, 2},
		goliner.Vector{3, 2, 1},
	}

	for _, r := range m {
		fmt.Println(r)
	}

	fmt.Println("answer:")

	m = m.LUDecompositionInPlace(goliner.NonePivot)
	for _, r := range m {
		fmt.Println(r)
	}
}
