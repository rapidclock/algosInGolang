package genutils

import "fmt"

func Print2D(dp [][]int) {
	for _, v := range dp {
		fmt.Println(v)
	}
}

func MaxOfTwo(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func MaxOfThree(a, b, c int) int {
	return MaxOfTwo(a, MaxOfTwo(b, c))
}

func MinOfTwo(a, b int) int {
	return -MaxOfTwo(-a,-b)
}

func MinOfThree(a, b, c int) int {
	return -MaxOfThree(-a, -b, -c)
}

func Abs(val int) int {
	if val < 0 {
		return -val
	} else {
		return val
	}
}