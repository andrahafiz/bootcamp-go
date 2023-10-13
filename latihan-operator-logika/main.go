package main

import "fmt"

func main() {
	var age int8 = 18
	var gender string = "male"

	if gender == "female" && age >= 21 {
		fmt.Println("Pelamar dapat dipekerjakan.")
	} else {
		fmt.Println("Pelamar tidak dapat dipekerjakan.")
	}

}
