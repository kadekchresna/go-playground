package main

import (

	"testing"
)



func BenchmarkMatrixReversedCombination(b *testing.B) {
	matrixA := [][]uint8{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{4, 5, 6, 7},
		{4, 5, 6, 7},
	}
	matrixB := [][]uint8{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{4, 5, 6, 7},
		{4, 5, 6, 7},
	}
	for n := 0; n < b.N; n++ {
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				matrixA[i][j] = matrixA[i][j] + matrixB[j][i]
			}
		}
	}

}