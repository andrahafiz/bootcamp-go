package main

import (
	"fmt"
	"strconv"
)

func main() {
	sum(5, 5)
}

func sum(inputA int, inputB int) {
	defer cetak("Program selesai")
	input := "Masukan angka pertama : " + strconv.Itoa(inputA)
	cetak(input)
	input = "Masukan angka kedua : " + strconv.Itoa(inputB)
	cetak(input)
	result := inputA + inputB
	cetak("Hasil penjumlahan : " + strconv.Itoa(result))

}

func cetak(text string) {
	fmt.Println(text)
}
