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

func CheckFloat64SliceEqual(a, b []bool) bool {
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