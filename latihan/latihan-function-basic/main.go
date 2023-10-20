package main

import "fmt"

func main() {
	var cars = map[string]string{
		"name":  "BMW",
		"color": "Black",
	}

	message := mergeString(cars["name"], cars["color"])
	cetak(message)
}

func mergeString(name string, color string) string {
	return "Mobil " + name + " berwarna " + color
}

func cetak(message string) {
	fmt.Println(message)
}
