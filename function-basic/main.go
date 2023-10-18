package main

import "fmt"

//Membuat function
func cetak() {
	fmt.Println("Fungsi cetak() di panggil")
	hello()
}

func hello() {
	fmt.Println("Hello World!")
}

// Parameter
func getName(name string) {
	fmt.Println(name)
}

func cetakNama(name string) {
	fmt.Println(name)
}

func cetakUmur(age int) {
	fmt.Println(age)
}

// Function berparemeter interface
func cetakNamaUmur(arg interface{}) {
	fmt.Println(arg)
}

// Multiple parameter
func getPeople(myName string, myAge int) {
	fmt.Println("Nama saya", myName, "Umur saya", myAge, "tahun")
}

// Function return value
func getSum(a, b int) int {
	return a + b
}

// Function combineString()
func combineString(str1 string, str2 string) string {
	return str1 + " " + str2
}

// Named return value
func addNumbers(x int, y int) (result int) {
	result = x + y
	return
}

func main() {
	// Memanggil function
	cetak()

	// Memanggil function berparameter
	getName("Andra")

	var names = []string{"Rey", "Jo", "NooBee"}
	var ages = []int{10, 13, 20}

	for _, name := range names {
		cetakNama(name)
	}

	for _, age := range ages {
		cetakUmur(age)
	}

	// Multiple parameter
	getPeople("Noobee Id", 26)

	fmt.Println(getSum(5, 1))

	// Named return value
	total := addNumbers(20, 30)
	fmt.Println(total)
}
