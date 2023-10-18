package main

import "fmt"

func main() {
	letter := "b"

	switch letter {
	case "a", "i", "u", "e", "o":
		fmt.Println("Huruf " + letter + " adalah huruf vokal")
	default:
		fmt.Println("Huruf " + letter + " adalah huruf konsonan")
	}
}
