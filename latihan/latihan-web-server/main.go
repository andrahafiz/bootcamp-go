package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Index called...")
	view := "Hello World!"
	w.Write([]byte(view))
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Println("About called...")
	view := "Ini adalah halaman about!"
	w.Write([]byte(view))
}

func handleRoute() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
}

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
