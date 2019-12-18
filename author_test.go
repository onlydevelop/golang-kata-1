package main

import (
	"testing"
)

func getTestAuthors() authors {
	return authors{
		author{
			email:     "null-walter@echocat.org",
			firstname: "Paul",
			lastname:  "Walter",
		},
		author{
			email:     "null-mueller@echocat.org",
			firstname: "Max",
			lastname:  "Müller",
		},
	}

}

func TestAuthorToString(t *testing.T) {
	testAuthor := getTestAuthors()[0]
	expected := "Paul Walter <null-walter@echocat.org>"
	actual := testAuthor.toString()
	assert(t, expected, actual)
}

func TestAuthorsToString(t *testing.T) {
	testAuthor := getTestAuthors()
	expected := "null-walter@echocat.org,null-mueller@echocat.org"
	actual := testAuthor.toString()
	assert(t, expected, actual)
}

func TestNewAuthor(t *testing.T) {
	line := "null-mueller@echocat.org;Max;Müller"
	expected := getTestAuthors()[1]
	actual := newAuthor(line)
	assert(t, expected.toString(), actual.toString())
}
