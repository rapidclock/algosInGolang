package main

import (
	"./dp"
	"./str"
	ds "./dstructs"
	"fmt"
	"../algosInGolang/extras"
	"../algosInGolang/graphs"
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

	treetest()
	stableMatchingTest()
	maxFlowTest()
}

func treetest() {
	tree := ds.NewTreeNode(1)
	tree.Left = ds.NewTreeNode(2)
	tree.Right = ds.NewTreeNode(3)
	fmt.Println(tree.SerializeTree())
	treeA := ds.BuildTree("1,2,4,#,#,#,3")
	fmt.Println(treeA)
	fmt.Println(treeA.PreOrder())
	fmt.Println(treeA.InOrder())
	fmt.Println(treeA.PostOrder())
	fmt.Println(treeA.CountNodes())
	treeB := ds.BuildTree("1,2,3,4,5")
	fmt.Println(treeB)
}

func stableMatchingTest() {
	men := [][]int{
		{1, 3, 4, 0, 2},
		{2, 1, 3, 4, 0},
		{1, 0, 3, 2, 4},
		{4, 3, 0, 2, 1},
		{1, 4, 3, 0, 2},
	}
	women := [][]int{
		{4, 2, 0, 1, 3},
		{0, 1, 3, 4, 2},
		{1, 3, 0, 2, 4},
		{1, 4, 3, 0, 2},
		{1, 3, 0, 4, 2},
	}
	result := extras.StableMatching(men, women)
	fmt.Println(result)
}

func maxFlowTest() {
	g := graphs.NewWeightedDirectedGraph()
	g.AddWeightedDirectedEdge(0, 1, 10)
	g.AddWeightedDirectedEdge(0, 2, 10)
	g.AddWeightedDirectedEdge(2, 1, 1)
	g.AddWeightedDirectedEdge(2, 3, 10)
	g.AddWeightedDirectedEdge(1, 3, 10)
	fmt.Println(g)
	paths, maxFlow := graphs.FordFulkersonMaxFlow(g, 0, 3)
	fmt.Println(maxFlow)
	for _, v := range paths {
		fmt.Println(v)
	}
}

func main() {
	fullRunner()
}
