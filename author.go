package main

import (
	"fmt"
	"io/ioutil"
	"log"
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
	if len(v) != 3 {
		return author{}
	}
	return author{
		email:     strings.TrimSpace(v[0]),
		firstname: strings.TrimSpace(v[1]),
		lastname:  strings.TrimSpace(v[2]),
	}
}

func loadAuthors(filename string) authors {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var a authors
	for _, line := range (strings.Split(string(bs), "\n"))[1:] {
		if strings.Contains(line, ";") {
			a = append(a, newAuthor(line))
		}
	}
	return a
}

func (a author) toString() string {
	if a.firstname == "" && a.lastname == "" && a.email == "" {
		return ""
	}
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
	for _, author := range a {
		fmt.Println(author.toString())
	}
}
