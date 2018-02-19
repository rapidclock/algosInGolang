package main

import (
	"./dp"
	"./str"
	"fmt"
)

func fullRunner() {
	fmt.Println(str.RabinKarp("Hello", "ll"))
	fmt.Println(str.RabinKarp("Hello", ""))
	fmt.Println(str.RabinKarp("", "1"))
	fmt.Println(str.RabinKarp("Hellollohlldj", "ll"), "\n")

	fmt.Println(str.KmpSubstringSearch("Hello", "ll"))
	fmt.Println(str.KmpSubstringSearch("Hello", ""))
	fmt.Println(str.KmpSubstringSearch("", "1"))
	fmt.Println(str.KmpSubstringSearch("Hellollohlldj", "ll"), "\n")

	fmt.Println(str.CreatePrefixArray("aacecaaa#aaacecaa"), "\n")

	fmt.Println(dp.LongestCommonSubsequence("gxtxayb", "aggtpanb"))
	fmt.Println(dp.LongestCommonSubsequence("gxt", "aggt"), "\n")

	fmt.Println(dp.EggDropping(1, 6))
	fmt.Println(dp.EggDropping(2, 6))
	fmt.Println(dp.EggDropping(4, 7))
	fmt.Println(dp.EggDropping(4, 100), "\n")
}

func main() {
	fullRunner()
}
