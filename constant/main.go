package main

import "fmt"

//Membuat constat di luar fungsi

const LENGTH int = 10

func main() {
	// Membuat constant di dalam fungsi
	const width = 5
	const PI = 3.14

	fmt.Println(LENGTH)
	fmt.Println(width)
	fmt.Println(PI)

	// Deklarasi multiple constant
	const (
		 X int 	= 10
		 y			= 3.14
	)
}
