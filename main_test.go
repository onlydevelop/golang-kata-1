package main

import (
	"testing"
)

func TestWelcomeMessage(t *testing.T) {
	expected := "Hello world!"
	actual := welcomeMessage()
	assert(t, expected, actual)
}
