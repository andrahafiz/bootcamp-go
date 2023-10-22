package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Indec called...")
	hello := "Hello World!"
	w.Write([]byte(hello)) // Konversi dari string ke slice of byte
}

// Untuk melakukan routing
func handleRoute() {
	http.HandleFunc("/", index)
}

// Untuk membuat server
func startServer(port string) {
	handleRoute()
	fmt.Println("Server running at port", port)
	http.ListenAndServe(port, nil)
}

func main() {
	/* Port bebas, yang penting belum digunakan, gunakan port yang mudah  */
	port := ":9999"
	startServer(port)
}
