package matrix

import (
	"errors"
)

func Multiply2DInt(matA, matB [][]int) ([][]int, error) {
	n, k1, k2, m := len(matA), len(matA[0]), len(matB), len(matB[0])
	if k1 != k2 {
		return nil, errors.New("math error: The two matrices are not of the form A[n*k], B[k*m]")
	}
	result := make([][]int, n)
	chans := make(chan bool, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, m)
		go rowProduct(matA, matB, result, i, m, k1, chans)
	}
	for i := 0; i < n; i++ {
		<- chans
	}
	return result, nil
}

func rowProduct(matA, matB, matC [][]int, i, m, k int, c chan bool) {
	for j := 0; j < m; j++ {
		matC[i][j] = singleIntProduct(matA, matB, i, j, k)
	}
	c <- true
}

func singleIntProduct(matA, matB [][]int, n, m, k int) int {
	sum := 0
	for i := 0; i < k; i++ {
		sum += matA[n][i] * matB[i][m]
	}
	return sum
}
