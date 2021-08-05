package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	press   string
}

func main() {
	var bookGo Books
	var bookPython Books

	bookGo.title = "Go编程"
	bookGo.author = "ahao"
	bookGo.subject = "Go语言开发"
	bookGo.press = "电子工业"

	bookPython.title = "python"
	bookPython.author = "aho"
	bookPython.subject = "Phthon开发"
	bookPython.press = "工商"

	fmt.Printf("bookGo title:%s\n", bookGo.title)

	fmt.Printf("bookPython title:%s\n", bookPython.title)

	fmt.Println("----------------")
	printBook(bookGo)
	printBook(bookPython)

	fmt.Println("----------------")
	//结构体指针
	printBooks(&bookGo)
	printBooks(&bookPython)
}

func printBook(books Books) {
	fmt.Printf("Book title:%s\n", books.title)
	fmt.Printf("Book author:%s\n", books.author)
	fmt.Printf("Book subject:%s\n", books.subject)
	fmt.Printf("Book press:%s\n", books.press)
}

//结构体指针
func printBooks(book *Books) {
	fmt.Printf("Book title:%s\n", book.title)
	fmt.Printf("Book author:%s\n", book.author)
	fmt.Printf("Book subject:%s\n", book.subject)
	fmt.Printf("Book press:%s\n", book.press)
}
