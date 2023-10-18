package main

import "fmt"

func main() {
	fruits := map[string]int{
		"Apple":  10,
		"Banana": 15,
		"Orange": 8,
		"Grape":  20,
	}

	fmt.Println("Sebelum ditambah/hapus:")
	fmt.Println(fruits)

	fmt.Println("Setelah menambahkan Mango dan Strawberry:")

	fruits["Manggo"] = 12
	fruits["Strawberry"] = 7
	fmt.Println(fruits)

	fmt.Println("Setelah menambahkan Mango dan Strawberry:")
	delete(fruits, "Orange")
	fmt.Println(fruits)

	fmt.Println("Jumlah apel yang tersedia adalah", fruits["Apple"])

	fmt.Println("Daftar buah-buahan beserta jumlahnya:")
	for index, fruit := range fruits {
		fmt.Println(index, ":", fruit)
	}
}
