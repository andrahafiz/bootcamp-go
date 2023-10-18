package main

import "fmt"

// Membuat struct
type Book struct {
	Title  string
	Author string
}

func main() {
	book1 := Book{
		Title:  "Pemrograman Go",
		Author: "John Doe",
	}

	fmt.Println("Informasi tentang Book 1:")
	fmt.Println("Judul :", book1.Title)
	fmt.Println("Penulis :", book1.Author)

	book2 := Book{
		Title:  "Algoritma dan Struktur Data",
		Author: "Jane Smith",
	}

	fmt.Println("Informasi tentang Book 1:")
	fmt.Println("Judul :", book2.Title)
	fmt.Println("Penulis :", book2.Author)

	book3 := Book{"Machine Learning for Beginners", " David Johnson"}
	fmt.Println("Judul :", book3.Title)
	fmt.Println("Penulis :", book3.Author)
}
