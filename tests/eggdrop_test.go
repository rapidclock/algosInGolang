package tests

import (
	"testing"
	"../dp"
)

func TestMinEggDrops(t *testing.T) {
	tables := []struct {
		eggs int
		floors int
		expected int
	}{
		{1,6, 6},
		{2,6, 3},
		{4,7, 3},
		{2,100, 14},
		{4,100, 8},
	}
	for _, x := range tables {
		actual := dp.EggDropping(x.eggs, x.floors)
		if actual != x.expected {
			t.Errorf("Error: For input (%d, %d) expected %d, but got %d instead", x.eggs, x.floors, x.expected, actual)
		}
	}
}