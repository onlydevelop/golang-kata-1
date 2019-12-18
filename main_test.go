package main

import (
	"testing"
)

func TestWelcomeMessage(t *testing.T) {
	expected := "Hello world1!"
	actual := welcomeMessage()
	assert(t, expected, actual)
}

func assert(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
