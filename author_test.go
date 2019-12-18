package main

import (
	"testing"
)

func getTestAuthors() authors {
	return authors{
		author{
			email:     "null-mueller@echocat.org",
			firstname: "Paul",
			lastname:  "Walter",
		},
		author{
			email:     "null-walter@echocat.org",
			firstname: "Max",
			lastname:  "Müller",
		},
	}

}

func TestAuthorToString(t *testing.T) {
	testAuthor := getTestAuthors()[0]
	expected := "Paul Walter <null-mueller@echocat.org>"
	actual := testAuthor.toString()
	assert(t, expected, actual)
}

func assert(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
