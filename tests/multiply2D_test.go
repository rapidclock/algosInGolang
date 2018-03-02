package tests

import (
	"testing"
	"math/rand"
	mat "../matrix"
	util "../genutils"
)

func getRandomMatrix(m, n int) [][]int {
	X := make([][]int, m)
	for i := 0; i < m; i++ {
		X[i] = make([]int, n)
		for j := 0; j < n; j++ {
			X[i][j] = rand.Int()
		}
	}
	return X
}

func BenchmarkMultiplyRowParallel(b *testing.B) {
	A := getRandomMatrix(10, 999)
	B := getRandomMatrix(999, 10)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		mat.Multiply2DInt(A, B)
	}
}

func TestMultiply2DIntMatrix(t *testing.T) {
	A := [][]int{{1,2,3,4,},{5,6,7,8},{4,5,8,2}}
	B := [][]int{{1,2,3,4},{5,6,7,8},{4,5,8,2},{3,10,5,6}}
	expected := [][]int{{35,69,61,50}, {87,161,153,130}, {67,98,121,84}}
	actual, _ := mat.Multiply2DInt(A, B)
	if !util.Check2DIntSliceEqual(actual, expected) {
		t.Errorf("Multiply did Not work! Expected :\n%v\n, Actual:\n%v\n", expected, actual)
	}
}