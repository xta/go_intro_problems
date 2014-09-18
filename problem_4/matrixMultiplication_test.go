package main

import (
	"testing"
)

func TestMatrix(t *testing.T) {
	actual := matrix()
	expected := [3][3]int{[3]int{1, 2, 3}, [3]int{2, 2, 2}, [3]int{-1, 0, 1}}
	if expected != actual {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestValidate(t *testing.T) {
	actual := validate([3][3]int{[3]int{1, 2, 3}, [3]int{2, 2, 2}, [3]int{-1, 0, 1}})
	expected := true
	if expected != actual {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestTransposeMatrix(t *testing.T) {
	actual := transposeMatrix([3][3]int{[3]int{1, 2, 3}, [3]int{2, 2, 2}, [3]int{-1, 0, 1}})
	expected := [3][3]int{[3]int{1, 2, -1}, [3]int{2, 2, 0}, [3]int{3, 2, 1}}
	if expected != actual {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestMultiplyMatrices(t *testing.T) {
	m1 := [3][3]int{[3]int{1, 2, 3}, [3]int{2, 2, 2}, [3]int{-1, 0, 1}}
	m2 := [3][3]int{[3]int{1, 2, -1}, [3]int{2, 2, 0}, [3]int{3, 2, 1}}

	actual := multiplyMatrices(m1, m2)
	expected := [3][3]int{[3]int{14, 12, 2}, [3]int{12, 12, 0}, [3]int{2, 0, 2}}
	if expected != actual {
		t.Error("Expected", expected, "got", actual)
	}
}
