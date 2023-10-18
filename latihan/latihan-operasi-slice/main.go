package main

import "fmt"

func main() {

	mySlice := [...]int{10, 20, 30, 40, 50}

	// mySlice := []string{"Cat", "Dog", "Cow"}
	slice1 := mySlice[0:3]
	slice2 := mySlice[3:]

	fmt.Println("====[mySlice]====")
	fmt.Println("Data \t :", mySlice)
	fmt.Println("Len \t :", len(mySlice))
	fmt.Println("Capacity :", cap(mySlice))
	fmt.Println("")

	fmt.Println("====[slice1]====")
	fmt.Println("Data \t :", slice1)
	fmt.Println("Len \t :", len(slice1))
	fmt.Println("Capacity :", cap(slice1))
	fmt.Println("")

	fmt.Println("====[slice2]====")
	fmt.Println("Data \t :", slice2)
	fmt.Println("Len \t :", len(slice2))
	fmt.Println("Capacity :", cap(slice2))
	fmt.Println("")

	slice1after := append(slice1, 60)
	fmt.Println("====[Setelah Append ke Slice1]====")
	fmt.Println("Data Sclice1 \t :", slice1after)
	fmt.Println("Len  Slice1 \t :", len(slice1after))
	fmt.Println("Capacity  Slice1 \t :", cap(slice1after))
	fmt.Println("Data Sclice2 \t :", slice2)
	fmt.Println("Len  Slice2 \t :", len(slice2))
	fmt.Println("Capacity  Slice2 \t :", cap(slice2))

	slice2after := append(slice2, 70)
	fmt.Println("====[Setelah Append ke Slice2]====")
	fmt.Println("Data Sclice1 \t :", slice1after)
	fmt.Println("Len  Slice1 \t :", len(slice1after))
	fmt.Println("Capacity  Slice \t :", cap(slice1after))
	fmt.Println("Data Sclice2 \t :", slice2)
	fmt.Println("Len  Slice2 \t :", len(slice2after))
	fmt.Println("Capacity  Slice2 \t :", cap(slice2))

}
