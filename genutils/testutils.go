package genutils

func CheckIntSliceEqual(a, b []int) bool {
	if a == nil && b == nil {
		return true;
	}
	if a == nil || b == nil {
		return false;
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func CheckStringSliceEqual(a, b []string) bool {
	if a == nil && b == nil {
		return true;
	}
	if a == nil || b == nil {
		return false;
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func CheckBoolSliceEqual(a, b []bool) bool {
	if a == nil && b == nil {
		return true;
	}
	if a == nil || b == nil {
		return false;
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func CheckFloat64SliceEqual(a, b []float64) bool {
	if a == nil && b == nil {
		return true;
	}
	if a == nil || b == nil {
		return false;
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Check2DIntSliceEqual(a, b [][]int) bool {
	if a == nil && b == nil {
		return true;
	}
	if a == nil || b == nil {
		return false;
	}
	rowA, colA, rowB, colB := len(a), len(a[0]), len(b), len(b[0])
	if rowA != rowB || colA != colB {
		return false
	}
	soln := true
	for i:= 0; i < rowA; i++ {
		soln = soln && CheckIntSliceEqual(a[i], b[i])
	}
	return soln
}

func Check2DFloat64SliceEqual(a, b [][]float64) bool {
	if a == nil && b == nil {
		return true;
	}
	if a == nil || b == nil {
		return false;
	}
	rowA, colA, rowB, colB := len(a), len(a[0]), len(b), len(b[0])
	if rowA != rowB || colA != colB {
		return false
	}
	isEqual := true
	for i:= 0; i < rowA; i++ {
		isEqual = isEqual && CheckFloat64SliceEqual(a[i], b[i])
	}
	return isEqual
}