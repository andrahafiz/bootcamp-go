package main

import "fmt"

var x int = 1

// y := 1  --> tidak bisa diluar fungsi

func main() {
	// VARIABLE EKSPLISIT
	var myName string = "NooBee Id"

	var yourName string
	yourName = "Ubit"

	var nickName = "Noobee"

	// VARIABLE IMPLISIT
	myAge := 22

	fmt.Println(myName)
	fmt.Println(yourName)
	fmt.Println(nickName)
	fmt.Println(myAge)
	fmt.Println(x)

	// Deklarasi variabel tanpa nilai awal
	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Deklarasi multiple variable
	var d, e, f, g = 1, 2, 3, "Hello"
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	// Deklarasi multiple variable dalam blok
	var (
		firstName string = "andra"
		age       int    = 28
	)

	fmt.Println(firstName)
	fmt.Println(age)

}
