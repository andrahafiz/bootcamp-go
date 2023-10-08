package main

import "fmt"

func main() {
	var myAge int8 = 20
	var sisterAge int8 = 30

	fmt.Printf("Umur saya sekarang adalah : %d\n", myAge)
	fmt.Printf("Umur kakak sekarang adalah : %d\n", sisterAge)

	// myAge++
	// fmt.Printf("Umur saya tahun depan adalah : %d\n", myAge)
	// sisterAge++
	// fmt.Printf("Umur kakak tahun depan adalah : %d\n", sisterAge)

	myAge += 5
	fmt.Printf("Umur saya tahun depan adalah : %d\n", myAge)
	sisterAge +=5
	fmt.Printf("Umur kakak tahun depan adalah : %d\n", sisterAge)

}
