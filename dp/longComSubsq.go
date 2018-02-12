package dp

import (
	"../genutils"
)

func LongestCommonSubsequence(word1, word2 string) int {
	n, m := len(word1), len(word2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = genutils.MaxOfTwo(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	genutils.Print2D(dp)
	return dp[n][m]
}

