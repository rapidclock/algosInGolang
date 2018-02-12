package str

func KmpSubstringSearch(bigString string, subString string) []int {
	n, m := len(bigString), len(subString)
	soln := make([]int,0)
	if n == 0 {
		return soln
	}
	if m == 0 {
		return append(soln, 0)
	}
	dp := CreatePrefixArray(subString)
	for i, j := 0, 0; i < n; {
		if bigString[i] == subString[j] {
			i , j = i + 1, j + 1
			if j >= m {
				soln = append(soln, i - m)
				j = 0
			}
		} else {
			if j > 0 {
				j = dp[j-1]
			} else {
				i += 1
			}
		}
	}
	return soln
}

func CreatePrefixArray(inputStr string) []int {
	n := len(inputStr)
	dp := make([]int, n)
	for i, j := 1, 0; i < n; {
		if inputStr[i] == inputStr[j] {
			dp[i] = j + 1
			i , j = i + 1, j + 1
		} else {
			if j > 0 {
				j = dp[j-1]
			} else {
				dp[i] = 0
				i += 1
			}
		}
	}
	return dp
}