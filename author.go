package main

import (
	"fmt"
	"strings"
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
	var authorEmails []string
	for _, author := range a {
		authorEmails = append(authorEmails, author.email)
	}
	return strings.Join(authorEmails, ",")
}

func (a authors) print() {

}
