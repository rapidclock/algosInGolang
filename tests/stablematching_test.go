package tests

import (
	"testing"
	"../extras"
	"reflect"
)

func TestStableMatching_SingleCase(t *testing.T) {
	men := [][]int{{0}}
	women := [][]int{{0}}
	expected := map[int]int{0: 0}
	actual := extras.StableMatching(men, women)
	if equal := reflect.DeepEqual(expected, actual); !equal {
		t.Errorf("Single Case Failed. \nExpected: %v\n Actual: %v\n", expected, actual)
	}
}

func TestStableMatching_MultipleCase(t *testing.T) {
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
	expected := map[int]int{0: 1, 1: 2, 2: 0, 3: 4, 4: 3}
	actual := extras.StableMatching(men, women)
	if equal := reflect.DeepEqual(expected, actual); !equal {
		t.Errorf("Multiple Case Failed. \nExpected: %v\n Actual: %v\n", expected, actual)
	}
}

func TestStableMatching_AllSamePrefs(t *testing.T) {
	men := [][]int{
		{0, 1, 2, 3, 4},
		{0, 1, 2, 3, 4},
		{0, 1, 2, 3, 4},
		{0, 1, 2, 3, 4},
		{0, 1, 2, 3, 4},
	}
	women := [][]int{
		{0, 1, 2, 3, 4},
		{0, 1, 2, 3, 4},
		{0, 1, 2, 3, 4},
		{0, 1, 2, 3, 4},
		{0, 1, 2, 3, 4},
	}
	expected := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4}
	actual := extras.StableMatching(men, women)
	if equal := reflect.DeepEqual(expected, actual); !equal {
		t.Errorf("AllSamePrefs Case Failed. \nExpected: %v\n Actual: %v\n", expected, actual)
	}
}

func TestStableMatching_GroupsWithSamePref(t *testing.T) {
	men := [][]int{
		{0, 1, 2, 3, 4},
		{0, 1, 2, 3, 4},
		{0, 1, 2, 3, 4},
		{0, 1, 2, 3, 4},
		{0, 1, 2, 3, 4},
	}
	women := [][]int{
		{4, 3, 2, 1, 0},
		{4, 3, 2, 1, 0},
		{4, 3, 2, 1, 0},
		{4, 3, 2, 1, 0},
		{4, 3, 2, 1, 0},
	}
	expected := map[int]int{0: 4, 1: 3, 2: 2, 3: 1, 4: 0}
	actual := extras.StableMatching(men, women)
	if equal := reflect.DeepEqual(expected, actual); !equal {
		t.Errorf("GroupsWithSamePref Case Failed. \nExpected: %v\n Actual: %v\n", expected, actual)
	}
}
