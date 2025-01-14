package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type book struct {
	title       string
	isbn        string
	authors     []string
	description string
}

type books []book

func newbook(line string) book {
	v := strings.Split(line, ";")
	if len(v) != 4 {
		return book{}
	}

	authorList := strings.Split(strings.TrimSpace(v[2]), ",")

	return book{
		title:       strings.TrimSpace(v[0]),
		isbn:        strings.TrimSpace(v[1]),
		authors:     authorList,
		description: strings.TrimSpace(v[3]),
	}
}

func loadBooks(filename string) books {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var a books
	for _, line := range (strings.Split(string(bs), "\n"))[1:] {
		if strings.Contains(line, ";") {
			a = append(a, newbook(line))
		}
	}
	return a
}

func (b books) findByISBN(isbn string) book {
	for _, eachBook := range b {
		if eachBook.isbn == isbn {
			return eachBook
		}
	}
	return book{}
}

func (b books) findByEmail(email string) book {
	for _, eachBook := range b {
		for _, eachEmail := range eachBook.authors {
			if eachEmail == email {
				return eachBook
			}
		}
	}
	return book{}
}

func (b book) toString() string {
	if b.title == "" && b.isbn == "" && len(b.authors) == 0 && b.description == "" {
		return ""
	}
	authors := strings.Join(b.authors, ",")
	return fmt.Sprintf("%s[%s]\n%s\n%s", b.title, b.isbn, authors, b.description)
}

func (b books) toString() string {
	var bookList []string
	for _, book := range b {
		bookList = append(bookList, book.toString())
	}
	return strings.Join(bookList, "\n")
}

func (b books) printBooks() {
	for _, book := range b {
		fmt.Println(book.toString())
	}
}
