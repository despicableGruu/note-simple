package test

import "testing"

func CheckError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func CheckErrorFatal(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func StringSliceEq(a, b []string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
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