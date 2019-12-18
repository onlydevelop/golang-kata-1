package main

import (
	"testing"
)

func assert(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Errorf("Expected \"%v\", but got \"%v\"", expected, actual)
	}
}
