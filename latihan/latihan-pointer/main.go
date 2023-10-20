package main

import "fmt"

type Student struct {
	Name  string
	Class string
}

func (s *Student) SetMyName(name string) {
	s.Name = name
}

func (s Student) CallMyName() {
	fmt.Printf("Hello My Name is %s", s.Name)
}

func main() {
	Student := Student{
		Name:  "Noobee",
		Class: "A1",
	}

	fmt.Println("Nama awal student:", Student.Name)
	Student.SetMyName("Ubit")
	Student.CallMyName() // Output yang diharapkan: "Nama pemain: John, Skor: 15"
}
