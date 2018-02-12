package str

import (
	"math"
)

func RabinKarp(bigString string, subString string) []int {
	n, m := len(bigString), len(subString)
	soln := make([]int,0)
	if n == 0 {
		return soln
	}
	if m == 0 {
		return append(soln, 0)
	}
	prime := 101
	effLenSub, effLenBig := len(subString)-1, len(bigString)-1
	subStrHash := calcHash(subString, 0, effLenSub, prime)
	bigStrHash := calcHash(bigString, 0, effLenSub, prime)
	for i, j := 0, effLenSub; j <= effLenBig; i, j = i+1, j+1 {
		if i > 0 {
			bigStrHash = calcRollingHash(bigStrHash, int(bigString[i-1]), int(bigString[j]), prime, effLenSub)
		}
		if subStrHash == bigStrHash && isSubStrBetweenIndex(bigString, subString, i, j) {
			soln = append(soln, i)
		}
	}
	return soln
}

func calcHash(str string, start, end, prime int) int {
	var hash int
	for i := 0; i < end-start+1; i++{
		charVal := int(str[start+i])
		hash += int(charVal * int(math.Pow(float64(prime), float64(i))))
	}
	return hash
}

func calcRollingHash(total, out, inc, prime, pow int) int {
	total = (total - out) / prime
	incVal := int(inc * int(math.Pow(float64(prime), float64(pow))))
	return total + incVal
}

func isSubStrBetweenIndex(str, subStr string, start, end int) bool {
	for i := 0; i < len(subStr); i++ {
		if str[start+i] != subStr[i] {
			return false
		}
	}
	return true
}