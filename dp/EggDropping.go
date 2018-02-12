package dp

import (
	"../genutils"
)

func EggDropping(eggs, floors int) int {
	dp := make([][]int, eggs + 1)
	for i := range dp {
		dp[i] = make([]int, floors + 1)
	}
	for i := range dp[0] {
		dp[1][i] = i
	}
	for i := 2; i <= eggs; i++{
		for j := 1; j <= floors; j++{
			for ;j < i; j++ {
				dp[i][j] = dp[i-1][j]
			}
			dp[i][j] = j
			for k := 1; k < j; k++{
				cur := 1 + genutils.MaxOfTwo(dp[i-1][k-1], dp[i][j-k])
				dp[i][j] = genutils.MinOfTwo(cur, dp[i][j])
			}
		}
	}
	return dp[eggs][floors]
}
