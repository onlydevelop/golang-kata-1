package main

import (
	"fmt"
)

type author struct {
	email     string
	firstname string
	lastname  string
}

type authors []author

func newAuthor(line string) author {
	return author{}
}

func load(filename string) authors {
	return nil
}

func (a author) toString() string {
	return fmt.Sprintf("%s %s <%s>", a.firstname, a.lastname, a.email)
}

func (a authors) toString() string {
	return ""
}

func (a authors) print() {

}
