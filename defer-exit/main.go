package main

import (
	"fmt"
	"os"
)

func cetak(text string) {
	fmt.Println(text)
}

func main() {
	// defer cetak("Text 1")
	// defer cetak("Text 2")
	// cetak("Text 3")
	// defer cetak("Text 4")
	// cetak("Text 5")

	// Defer akan di jalankan di akhir blok function
	num1 := 5

	if num1%2 > 0 {
		cetak("Ini adalah bilangan ganjil")
		defer cetak("akan berakhir setelah proses diatas selesai")

		//Anynomous function
		// func() {
		// 	defer cetak("akan berakhir setelah proses diatas selesa")
		// }()
	}

	cetak("Oh tentu tidak, Defer di eksekusi setelah ini")

	//Exit
	names := []string{"NooB", "Bee", "Go", "NodeJS"}

	for _, name := range names {
		if name == "Go" {
			os.Exit(1)
		}
		cetak(name)
	}
}
