package main

import (
	"os"
	"testing"
)

func TestHashFile(t *testing.T) {
	f, _ := os.Open("sample.txt")
	defer f.Close()

	expected := "edc715389af2498a623134608ba0a55b"
	actual := hashFile(f)

	if expected != actual {
		t.Error("Expected", expected, "got", actual)
	}
}
