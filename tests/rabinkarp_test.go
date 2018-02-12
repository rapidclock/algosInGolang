package tests

import (
	"testing"
	"../str"
	"../genutils"
)

func TestRabinKarp(t *testing.T) {
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
		actual := str.RabinKarp(x.bigStr, x.subStr)
		if !genutils.CheckIntSliceEqual(actual, x.expected) {
			t.Errorf("Error: For input (%s, %s) expected %d, but got %d instead", x.bigStr, x.subStr, x.expected, actual)
		}
	}
}