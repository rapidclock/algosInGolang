package tests

import (
	"testing"
	"../str"
	"../genutils"
)

func TestEmpty(t *testing.T) {
	tables := []struct {
		x string
		expected []int
	}{
		{"Hello", []int{2}},
		{"sfdsf", []int{2}},
	}
	for _, x := range tables {
		actual := x.expected
		if !genutils.CheckIntSliceEqual(actual, x.expected) {
			t.Errorf("Error: For input (%s) expected %d, but got %d instead", x.x, x.expected, actual)
		}
	}
}

func TestKMP(t *testing.T) {
	tables := []struct {
		bigStr   string
		subStr   string
		expected []int
	}{
		{"Hello", "ll", []int{2}},
		{"Hello", "", []int{0}},
		{"", "1", []int{}},
		{"Hellollohlldj", "ll", []int{2, 5, 9}},
	}

	for _, x := range tables {
		actual := str.KmpSubstringSearch(x.bigStr, x.subStr)
		if !genutils.CheckIntSliceEqual(actual, x.expected) {
			t.Errorf("Error: For input (%s, %s) expected %d, but got %d instead", x.bigStr, x.subStr, x.expected, actual)
		}
	}
}

func TestPrefixArray(t *testing.T) {
	tables := []struct {
		inpStr   string
		expected []int
	}{
		{"aba", []int{0, 0, 1}},
		{"abab", []int{0, 0, 1, 2}},
		{"ababc", []int{0, 0, 1, 2, 0}},
		{"ababa", []int{0, 0, 1, 2, 3}},
		{"abcaby", []int{0, 0, 0, 1, 2, 0}},
		{"abcabcab", []int{0, 0, 0, 1, 2, 3, 4, 5}},
		{"abcabcabc", []int{0, 0, 0, 1, 2, 3, 4, 5, 6}},
		{"acdacdaed", []int{0, 0, 0, 1, 2, 3, 4, 0, 0}},
		{"abaababaab", []int{0, 0, 1, 1, 2, 3, 2, 3, 4, 5}},
		{"abxefabxefabxef", []int{0, 0, 0, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"acacabacacabacacac", []int{0, 0, 1, 2, 3, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 4}},
		{"aacecaaa#aaacecaa", []int{0, 1, 0, 0, 0, 1, 2, 2, 0, 1, 2, 2, 3, 4, 5, 6, 7}},
	}

	for _, x := range tables {
		actual := str.CreatePrefixArray(x.inpStr)
		if !genutils.CheckIntSliceEqual(actual, x.expected) {
			t.Errorf("Error: For input %s expected %d, but got %d instead", x.inpStr, x.expected, actual)
		}
	}
}
