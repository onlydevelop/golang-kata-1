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
	v := strings.Split(line, ";")
	return author{
		email:     strings.TrimSpace(v[0]),
		firstname: strings.TrimSpace(v[1]),
		lastname:  strings.TrimSpace(v[2]),
	}
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
