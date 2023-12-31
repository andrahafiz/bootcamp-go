package main

import "fmt"

// Interface method
type Calculate interface {
	Divide() float32
	Multiply() float32
	Sum() float32
	Substract() float32
}

type Num1 struct {
	Num     float32
	Divider float32
}

func NewNum1(num float32, divider float32) Calculate {
	return &Num1{Num: num, Divider: divider}
}

func (n Num1) Divide() float32 {
	return n.Num / n.Divider
}

func (n Num1) Multiply() float32 {
	return n.Num * n.Num
}

func (n Num1) Sum() float32 {
	return n.Num + n.Num
}

func (n Num1) Substract() float32 {
	return n.Num - n.Num
}

func main() {
	// Interface method
	num := NewNum1(4.0, 2.0)

	fmt.Println(num.Divide())
	fmt.Println(num.Multiply())
	fmt.Println(num.Sum())
	fmt.Println(num.Substract())
}
