package book

import "fmt"

// interface
type Printable interface {
	PrintInfo()
}

func Print(p Printable) {
	p.PrintInfo()
}

type Book struct {
	Title  string
	Author string
	pages  int //si inicia con minuscula la primera letra se considera privado
}

func NewBook(title string, author string, pages int) *Book {
	return &Book{
		Title:  title,
		Author: author,
		pages:  pages,
	}
}

func (book *Book) SetPages(pages int) {
	book.pages = pages
}

func (book *Book) GetPages() int {
	return book.pages
}

func (book *Book) PrintInfo() {
	fmt.Printf("Title: %s\n Author: %s\n Pages: %d\n", book.Title, book.Author, book.pages)
}

type TextBook struct {
	Book
	editorial string
	level     string
}

func NewTextBook(title, author string, pages int, editorial, level string) *TextBook {
	return &TextBook{
		Book:      Book{title, author, pages},
		editorial: editorial,
		level:     level,
	}
}

func (textBook *TextBook) PrintInfo() {
	fmt.Printf("Title: %s\n Author: %s\n Pages: %d\n Editoral: %s\n Level: %s \n",
		textBook.Title, textBook.Author, textBook.pages, textBook.editorial, textBook.level)
}
