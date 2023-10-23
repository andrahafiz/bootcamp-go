package main

import (
	"fmt"
)

func divison(number1, number2 int) {
	fmt.Printf("Masukan angka pertama: %d\n", number1)
	fmt.Printf("Masukan angka kedua: %d\n", number2)
	defer catchError()
	if number2 == 0 {
		panic("Panic: Tidak dapat membagi oleh nol")
	} else {
		result := number1 / number2
		fmt.Printf("Hasil pembagian: %d\n", result)
	}
}

func catchError() {
	err := recover()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	divison(10, 5)
}
