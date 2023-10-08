package main

import "fmt"

func main() {
	var productA int = 10000
	var productB int = 15000
	productC := 7000

	var total int
	total = productA + productB + productC

	fmt.Printf("Total harga belanja : %d\n", total)
}
