package tests

import (
	"testing"

	"../genutils"
	"../str"
)

func TestAhoCorasick(t *testing.T) {
	testTable := []struct {
		sentence string
		dict     []string
		expected map[string][]int
	}{
		{"bcaab", []string{"a", "ab", "aab", "aac", "bc", "bd"}, map[string][]int{"bc":[]int{0}, "a":[]int{2,3}, "aab":[]int{2}, "ab":[]int{3}}},
		{"bcaab", []string{"", "a", "ab", "aab", "aac", "bc", "bd"}, map[string][]int{"bc":[]int{0}, "a":[]int{2,3}, "aab":[]int{2}, "ab":[]int{3}}},
		{"bcxaabp", []string{"a", "ab", "aab", "aac", "bc", "bd"}, map[string][]int{"bc":[]int{0}, "a":[]int{3,4}, "aab":[]int{3}, "ab":[]int{4}}},
		{"aaaaaaab", []string{"a", "aa", "aaaa", "ab", "b"}, map[string][]int{"a":[]int{0,1,2,3,4,5,6},"aa":[]int{0,1,2,3,4,5},"aaaa":[]int{0,1,2,3}, "ab":[]int{6}, "b":[]int{7}}},
		{"xcvvvt", []string{"a", "ab", "aab", "aac", "bc", "bd"}, map[string][]int{}},
	}
	for _, x := range testTable {
		actual := str.AhoCorasick(x.sentence, x.dict)
		if len(actual) != len(x.expected) {
			t.Errorf("For %s, len(expected) = %d, len(actual) = %d", x.sentence, len(x.expected), len(actual))
		}
		for k, actVal := range actual {
			if expVal, ok := x.expected[k]; !ok && !genutils.CheckIntSliceEqual(actVal, expVal){
				t.Errorf("For %s, Expected occ for %s is %v, but found %v", x.sentence, k, expVal, actVal)
			}
		}
	}
}